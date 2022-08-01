package main

import (
	"fmt"
	"go_8_mage/week11/socket/common"
	"net"
	"time"
)



//收发简单地字符串消息
func main() {
	go common.DealSigPipe()
	ip := "127.0.0.1"
	port := "5656"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ip+":"+port)
	common.CheckError(err)
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	common.CheckError(err)
	fmt.Println("waiting for client connection ......")
	conn,err:=listener.Accept()
	common.CheckError(err)
	fmt.Printf("establish connection to client %s\n",conn.RemoteAddr().String()) //操作系统会随机给客户端分配一个 49152~65535上的端口号
	request := make([]byte,256) //设定一个最大程度，防止flood attack (DDos攻击)
	n,err:=conn.Read(request)
	common.CheckError(err)
	fmt.Printf("receive request %s \n",string(request[:n]))
	m,err:= conn.Write([]byte("hello "+string(request[:n])))
	common.CheckError(err)
	fmt.Printf("write response %d bytes\n",m)

	time.Sleep(3 * time.Second)
	m,err = conn.Read(request) //对方关闭连接后，再在conn上调用Read会报错 EOF
	if err !=nil{
		fmt.Println(err)
	} else{
		fmt.Printf("read request %d bytes\n",m)
	}
	m,err = conn.Read(request) //对方关闭连接后，再在conn上调用Read会报错 EOF
	if err !=nil{
		fmt.Println(err)
	} else{
		fmt.Printf("read request %d bytes\n",m)
	}

	m,err=conn.Write([]byte("oops")) //对方关闭连接后，再在conn上调用write可能会报错“broken pipe”，也可能不会，跟tcp缓冲区情况有关
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Printf("write response %d bytes\n",m)
	}

	m,err=conn.Write([]byte("oops")) //对方关闭连接后，再在conn上调用write可能会报错“broken pipe”，也可能不会，跟tcp缓冲区情况有关
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Printf("write response %d bytes\n",m)
	}

}
