package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

var fileChan = make(chan string,10000)

var writeFinish = make(chan struct{})
var wg2 sync.WaitGroup

func readFile2(fileName string){
	defer func(){
		wg2.Done()
	}()
	// 打开文件
	fin,err := os.Open(fileName)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer fin.Close()
//	构建FileReader
	reader :=bufio.NewReader(fin)
	for {
		line,err := reader.ReadString('\n')
		if err != nil{
			if err == io.EOF{
				if len(line) >0 { // 文件最后一行不是空行
					line += "\n"
					fileChan <- line
				}
				break
			} else{
				fmt.Println(err)
				break
			}
		}else {
			fileChan <- line
		}
	}
}

func writeFile2(fileName string){
	defer close(writeFinish)
	fout,err := os.OpenFile(fileName,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,os.ModePerm)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	defer fout.Close()
	writer:=bufio.NewWriter(fout)
	for {
		line,ok := <-fileChan
		if len(line) == 0{ //如果管道被关闭，Line就是空字符串
			if ok{ // Line是空字符串，是因为fileChan里面就存在一个空字符串
				writer.WriteString(line)
			} else {
				break
			}
		}else{
			writer.WriteString(line)
		}
	}
	writer.Flush()
}

func main(){
	wg2.Add(3)
	for i:=1;i<=3;i++{
		fileName :="data/"+strconv.Itoa(i)+".txt"
		go readFile2(fileName)
	}
	go writeFile2("data/merge.txt")
	wg2.Wait()
	close(fileChan)
	<-writeFinish
}