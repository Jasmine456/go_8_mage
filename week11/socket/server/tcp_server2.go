package main

import (
	"encoding/json"
	"fmt"
	"go_8_mage/week11/socket/common"
	"net"
	"strconv"
)

//序列化请求和响应的结构体
func main() {
	ip := "127.0.0.1"
	port := 5656
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ip+":"+strconv.Itoa(port))
	common.CheckError(err)
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	common.CheckError(err)
	fmt.Println("waiting for client connection ......")
	conn, err := listener.Accept()
	common.CheckError(err)
	fmt.Printf("establish connection to client %s \n", conn.RemoteAddr().String()) //操作系统会随机给客户端分配一个 49152~65535上的端口号
	requestBytes := make([]byte, 256)                                                   //设定一个最大程度，防止flood attack (DDos攻击)
	read_len,err:=conn.Read(requestBytes)
	common.CheckError(err)
	fmt.Printf("receive request %s\n",string(requestBytes))//[]byte转string时，0后面的会自动被截掉

	var request common.Request
	json.Unmarshal(requestBytes[:read_len],&request)// json反序列化时会把0都考虑在内，所以需要只读前read_len个字节
	response:=common.Response{Sum: request.A+request.B}

	responseBytes,_:=json.Marshal(response)
	_,err=conn.Write(responseBytes)
	common.CheckError(err)
	fmt.Printf("write response %s\n",string(responseBytes))
	conn.Close()



}
