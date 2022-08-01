package main

import (
	"bytes"
	"fmt"
	"go_8_mage/week11/socket/common"
	"net"
)

//定义分隔符
var delemeter = [5]byte{12,56,123,56,97}

func readDataGram(conn net.Conn) []string{
	dataGrams := []string{}
	content:=make([]byte,1024)
	n,err:=conn.Read(content)
	if err!=nil{
		return nil
	}
	begin:=0
	for i :=0;i<n;i++{
		if content[i] == '|'{
			dataGrams = append(dataGrams,string(content[begin:i]))
			begin=i+1
		}
	}
	return dataGrams
}


// TCP是面向字节流的
func main(){
	tcpAddr,err:=net.ResolveTCPAddr("tcp4","127.0.0.1:5656")
	common.CheckError(err)
	listener,err:=net.ListenTCP("tcp4",tcpAddr)
	common.CheckError(err)
	conn,err:=listener.Accept()
	common.CheckError(err)
	fmt.Printf("connect to client %s\n",conn.RemoteAddr().String())

	content :=make([]byte,2)
	buffer:=bytes.Buffer{}
	for {
		n,err:=conn.Read(content)
		common.CheckError(err)
		for i:=0;i<n;i++{
			c:=content[i]
			if c!=0{
				buffer.WriteByte(c)
			}else{
				fmt.Println(string(buffer.Bytes()))
				buffer.Reset()
			}
		}
	}


	//dataGrams:=readDataGram(conn)
	//fmt.Println(dataGrams)

	//content:=make([]byte,256)
	//n,err:=conn.Read(content) //同步IO模式，如果socket上没有数据可读，改行代码会阻塞。
	//common.CheckError(err)
	//fmt.Println(string(content[:n]))
	//
	//content=make([]byte,4)
	//n,err=conn.Read(content) //TCP是面向字节流的，一次Read到的数据可能包含了多个报文，也可能只包含了半个报文，一条报文在什么地方结束需要通信双方事先约定好
	//common.CheckError(err)
	//fmt.Println(string(content[:n]))
	//
	//n,err=conn.Read(content)
	//common.CheckError(err)
	//fmt.Println(string(content[:n]))
	//



}
