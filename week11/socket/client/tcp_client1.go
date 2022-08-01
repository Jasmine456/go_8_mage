package main

import (
	"fmt"
	"go_8_mage/week11/socket/common"
	"net"
	"strconv"
)

//收发简单地字符串信息
func main(){
	ip:="127.0.0.1"
	port:=5656
	tcpAddr,err:=net.ResolveTCPAddr("tcp4",ip+":"+strconv.Itoa(port))
	common.CheckError(err)
	fmt.Printf("get_ip %s port %d\n",tcpAddr.IP.String(),tcpAddr.Port)
	conn,err:=net.DialTCP("tcp",nil,tcpAddr)// Dial 拨号
	common.CheckError(err)
	fmt.Printf("establish connection to server %s\n",conn.RemoteAddr().String())
	n,err:=conn.Write([]byte("jasmine"))
	common.CheckError(err)
	fmt.Printf("write request %d bytes\n",n)
	response:=make([]byte,256)
	_,err=conn.Read(response)
	common.CheckError(err)
	fmt.Printf("receive response:%s\n",string(response))
	//conn.Close()
	//fmt.Println("close connection")
}
