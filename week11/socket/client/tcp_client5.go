package main

import (
	"fmt"
	"go_8_mage/week11/socket/common"
	"net"
	"time"
)

//定义分隔符
var delemeter = [5]byte{12,56,123,56,97}

func main(){
	tcpAddr,err:=net.ResolveTCPAddr("tcp4",":5656")
	common.CheckError(err)
	conn,err:=net.DialTCP("tcp",nil,tcpAddr)
	common.CheckError(err)
	defer conn.Close() //连接关闭后，再调用conn.Write()和conn.Read()会返回fatal error
	fmt.Printf("connect to server %s \n",conn.RemoteAddr().String())

	time.Sleep(2 * time.Second)

	// 连续写两次
	n,err := conn.Write([]byte("hello| 0"))
	common.CheckError(err)
	fmt.Printf("向服务端发送了%d个字节\n",n)
	//conn.Write([]byte{0})
	conn.Write(delemeter[:]) //[:]将数组转为切片

	n,err = conn.Write([]byte("world| 0"))
	common.CheckError(err)
	fmt.Printf("向服务端发送了%d个字节\n",n)
	//conn.Write([]byte{0})
	conn.Write(delemeter[:])

}