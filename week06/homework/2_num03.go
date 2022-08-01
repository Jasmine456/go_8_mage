package main

import (
	"bufio"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 把文件filename里的内容 追加到writer 所指的文件内

func appendFile(mergedFile string, fileName string) {
	fmt.Println(fileName)
	fout,err := os.OpenFile(mergedFile,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer fout.Close()
	writer := bufio.NewWriter(fout)

	fin,err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fin.Close()
	reader := bufio.NewReader(fin)

	for {
		line,_,err:=reader.ReadLine()
		//fmt.Println(string(line))
		if err == nil{
			writer.Write(line)
			writer.WriteByte('\n')
			//fmt.Println(string(line))
		}else {
			if err == io.EOF{
				writer.Write(line)
				writer.WriteByte('\n')
			}
			break
		}
	}
	writer.Flush()
}

// 把dir目录下的所有.txt文件合并到mergedFile文件里面去
func mergeFile(dir string,mergedFile string){
	fileInfos,err:=ioutil.ReadDir(dir) //获取dir的所有子目录/文件
	if err !=nil{
		fmt.Printf("遍历目录%s时出错:%s\n",dir,err.Error())
	}
	fmt.Println(fileInfos)
	for _,fileInfo := range fileInfos{ //遍历所有子目录/文件
		if fileInfo.IsDir(){// 如果是子目录
			subDir := filepath.Join(dir,fileInfo.Name())
			mergeFile(subDir,mergedFile) //把dir的子目录传给递归函数
		}else { //如果是子文件
			//fmt.Println(fileInfo.Name())
			if strings.HasSuffix(fileInfo.Name(),".txt"){ //确保文件以txt结尾
				appendFile(mergedFile,filepath.Join(dir,fileInfo.Name())) // 把子文件里的内容追加到mergedFile文件里面去
			}
		}
	}
}

func copress(infile,outfile string){
	//压缩
	fin,err := os.Open(infile)
	defer fin.Close()
	if err!=nil{
		fmt.Println(err)
	}
	fout,err:=os.OpenFile(outfile,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)
	defer fout.Close()
	if err!=nil{
		fmt.Println(err)
	}
	writer := zlib.NewWriter(fout)
	defer writer.Close()

	content := make([]byte,2048)
	for {
		n,err := fin.Read(content)
		if err ==nil{
			writer.Write(content[:n])
		} else{
			if err == io.EOF{
				if n>0{
					writer.Write(content[:n])
				}
			}else{
				fmt.Println(err)
			}
			break
		}
	}

//	解压
	fin,err = os.Open(outfile)
	defer fin.Close()
	if err !=nil{
		fmt.Println(err)
	}
	reader,err := zlib.NewReader(fin)
	if err !=nil{
		fmt.Println(err)
	}
	io.Copy(os.Stdout,reader)


}

func main(){
	mergeFile("data","data/new.md")
	copress("data/new.md","data/new.md.zlib")
}