package main

import (
	"bufio"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*
3. 把一个目录下的所有.txt文件合一个大的.txt文件，再对这个大文件进行压缩
 */

func traver_dir(file_path string)([]string,error){
	file_slice := []string{}
	if fileInfos,err := ioutil.ReadDir(file_path);err != nil{
		fmt.Printf("%s 目录不存在",file_path)
		return file_slice,err
	}else{
		if len(fileInfos)==0{
			fmt.Printf("%s 目录为空",file_path)
			return file_slice,nil
		}
		for _,fileInfo := range fileInfos{
			//fmt.Println(fileInfo)
			//排除文件不是目录
			if fileInfo.IsDir()==false{
				//取目录名字是以.txt结尾的
				if strings.HasSuffix(fileInfo.Name(),".txt") {
					//将匹配的目录名字塞到 一个切片
					file_slice = append(file_slice, fileInfo.Name())
				}
			}
		}
	}
	return file_slice,nil
}

func read_file(file_name string)(file string,err error){
	if fin,err := os.Open(file_name);err !=nil{
		fmt.Printf("open file failed %v\n",err)
		return "",err
	}else {
		defer fin.Close()
		reader := bufio.NewReader(fin)
		for { //无限循环
			if line, err := reader.ReadString('\n'); err != nil { //指定分隔符
				if err == io.EOF {
					if len(line) > 0 { //如果最后一行没有换行符，则此时最后一行就存在line里
						file+=line
					}
					break //已读到文件末尾
				} else {
					fmt.Printf("read file failed: %v\n", err)
				}
			} else {
				//line = strings.TrimRight(line, "\n") //line里面是包含换行符的，需要去掉
				//fmt.Println(line)
				file+=line
			}
		}
	}
	fmt.Println(file)
	return file,nil
}
func write_file(file string,new_filename string) {
	//OpenFile()比Open()有更多的参数选项。os.O_WRONLY以只写的方式打开文件，os.O_TRUNC把文件之前的内容先清空掉，os.O_CREATE如果文件不存在则先创建，0666新建文件的权限设置
	if fout, err := os.OpenFile(new_filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err != nil {
		fmt.Printf("open file faied: %v\n", err)
	} else {
		defer fout.Close() //别忘了关闭文件句柄
		//写文本文件建议使用
		writer := bufio.NewWriter(fout)
		writer.WriteString(file)
		writer.Flush() //buffer中的数据量积累到一定程度后才会真正写入磁盘。调用Flush强行把缓冲中的所有内容写入磁盘
	}
}

func compress(new_filename,new_compress  string){
	fin,err := os.Open(new_filename)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fout,err := os.OpenFile(new_compress,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)
	if err != nil{
		fmt.Println(err)
		return
	}
	bs := make([]byte,1024)
	writer := zlib.NewWriter(fout)
	for {
		n,err := fin.Read(bs)
		if err != nil{
			if err == io.EOF{
				break
			}else{
				fmt.Println(err)
			}
		}else{
			writer.Write(bs[:n])
		}
	}
	writer.Close()
	fout.Close()
	fin.Close()
}

func main(){
	new_filename:="data/new.txt.a"
	new_compress := "data/new.zlib"
 	//1. 首先遍历这个目录下所有文件，把.txt的文件存到1个slice中
	file_slice,_:=traver_dir("data")

	//2. 批量读取，将读取的所有内容 写入一个新的文件
	var file string
	for num,ele:=range file_slice{
		fmt.Println(ele,num)
		file,_ =read_file("data/"+ele)
		//fmt.Println(file)
		write_file(file,new_filename)
	}
	//3. 压缩这个新的文件
	compress(new_filename,new_compress)
}