package main

import (
	"fmt"
	"go_8_mage/week11/socket/common"
	"net"
	"time"
)

//UDP是无连接的，udp是面向报文的
func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":5656")
	common.CheckError(err)
	conn, err := net.ListenUDP("udp", udpAddr) //跟TCP不一样，及时没有client请求建立连接，该行代码也不会阻塞，它返回的conn是个虚拟连接
	common.CheckError(err)
	defer conn.Close()
	fmt.Println("build connect")

	time.Sleep(3 * time.Second)
	fmt.Println("awake")

	//content 内容
	//接收缓冲区有两条报文可读
	content := make([]byte, 256)
	n, remoteAddr, err := conn.ReadFromUDP(content) //UDP是面向报文的，一次Read只读一个报文
	common.CheckError(err)
	fmt.Println(string(content[:n]))

	content = make([]byte, 4)
	n, remoteAddr, err = conn.ReadFromUDP(content) //如果没有把一个报文读完，后面的内容会被丢弃掉，下次就读不到了

	//这里没读完 err会有一个错误，
	//A message sent on a datagram socket was larger than the internal message buffer or some other network limit, o
	//r the buffer used to receive a datagram into was smaller than the datagram itself
	//common.CheckError(err)
	fmt.Println(string(content[:n]))

	content = make([]byte, 4)
	n, remoteAddr, err = conn.ReadFromUDP(content) //如果没有把一个报文读完，后面的内容会被丢弃掉，下次就读不到了
	common.CheckError(err)
	fmt.Println(string(content[:n]))

	n, err = conn.WriteToUDP([]byte("oops"), remoteAddr) //不会报错
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("write response %d bytes\n", n)
	}

	n, err = conn.WriteToUDP([]byte("oops"), remoteAddr) //不会报错
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("write response %d bytes\n", n)
	}

	n, remoteAddr, err = conn.ReadFromUDP(content) // 接收缓冲区已无报文可读，该行代码会阻塞
	if err!=nil{
		fmt.Println(err)
	} else{
		fmt.Printf("read request %d bytes\n",n)
	}


}
