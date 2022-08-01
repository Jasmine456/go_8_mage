package main

import (
	"encoding/json"
	"fmt"
	"go_8_mage/week11/socket/common"
	"net"
	"strconv"
)


func handleRequest(conn net.Conn){
	defer conn.Close()
	requestBytes:=make([]byte,256)
	read_len,err:=conn.Read(requestBytes)
	common.CheckError(err)
	fmt.Printf("receive request %s\n",string(requestBytes))

	var request common.Request
	json.Unmarshal(requestBytes[:read_len],&request)
	response := common.Response{Sum: request.A+request.B}

	responseBytes,_:=json.Marshal(response)
	_,err=conn.Write(responseBytes)
	common.CheckError(err)
	fmt.Printf("write response %s\n",string(responseBytes))

}



//接收多个客户端请求
func main(){
	ip := "127.0.0.1"
	port := 5656
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ip+":"+strconv.Itoa(port))
	common.CheckError(err)
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	common.CheckError(err)
	fmt.Println("waiting for client connection ......")
	for {
		conn,err:=listener.Accept()
		if err!=nil{
			continue
		}
		fmt.Printf("establish connection to client %s\n",conn.RemoteAddr().String())
		go handleRequest(conn)
	}

}
