package main

import (
	"fmt"
	"os"
)

/*
4. 自己实现一个BufferedFileWriter。
type BufferedFileWriter struct{}
func (BufferedFileWriter) Write(cont string){}  //大部分情况下只是把cont写到缓冲区(go 变量)，
当缓冲区写满之后才会触发真正的写磁盘操作

 */

// 有缓冲的写文件  BufferedFileWriter
type BufferedFileWriter struct {
	cache []byte
	position int
	fout *os.File
}

// BufferedFileWriter的构造函数，返回一个带指针的buffer
func NewBufferedFileWriter(fout *os.File,cacheSize int) *BufferedFileWriter{
	return &BufferedFileWriter{
		cache: make([]byte,cacheSize),
		fout: fout,
		position:0,
	}
}

func (writer *BufferedFileWriter) writeByte(content []byte){
	//	优先把content 追加到cache里面，如果cache填满了，则把cache里的内容写入磁盘，再清空cache
	for i:=0;i<len(content);i++{
		if writer.position <len(writer.cache){ //缓冲还没满
			writer.cache[writer.position] = content[i] //在缓冲里追加一个字节
			writer.position++ //position加1
		}else{ //缓冲刚刚满
			writer.fout.Write(writer.cache[0:writer.position])
			writer.position=0 //指针置回到起点
			writer.cache[writer.position] = content[i]
			writer.position++
		}
	}
}

func (writer *BufferedFileWriter) WriterString (content string){
	writer.writeByte([]byte(content))
}

func (writer *BufferedFileWriter) Flush(){
	if writer.position >0{
		fmt.Printf("cache 还剩余内容 %d\n",writer.position)
		writer.fout.Write(writer.cache[0:writer.position])
		writer.position=0
	}
}


func main(){
	fout,err := os.OpenFile("jasmine.txt",os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)
	defer fout.Close()
	if err!=nil{
		fmt.Println(err)
	}
	writer:=NewBufferedFileWriter(fout,2048)
	for i:=0;i<100;i++{
		writer.WriterString("111111111111111\n")
	}
	writer.Flush()
}

