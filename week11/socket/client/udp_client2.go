package main

import (
	"fmt"
	"go_8_mage/week11/socket/common"
	"net"
	"time"
)

//UDP是无连接的。UDP是面向报文的
func main() {
	conn, err := net.DialTimeout("udp", ":5656", 30*time.Minute) //跟TCP不一样，可以先启动client。由于conn是个虚拟连接，所以该行代码不需要阻塞，会立即返回
	common.CheckError(err)
	fmt.Printf("connect to server %s\n", conn.RemoteAddr().String())

	time.Sleep(5 * time.Second)

	//	连续写2次
	_, err = conn.Write([]byte("hello|"))
	common.CheckError(err)

	_, err = conn.Write([]byte("world|"))
	common.CheckError(err)

	time.Sleep(1 * time.Second)
	fmt.Printf("close connect au %v\n", time.Now())
	conn.Close() //关闭连接

	//	关闭连接后，读写都会报错： use of closed network connection
	_,err=conn.Write([]byte("oops"))
	fmt.Println(err)
	content:=make([]byte,256)
	_,err=conn.Read(content)
	fmt.Println(err)


}
