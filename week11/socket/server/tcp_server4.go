package main

import (
	"encoding/json"
	"fmt"
	"go_8_mage/week11/socket/common"
	"net"
	"strconv"
	"time"
)

func handleRequest2(conn net.Conn){
	conn.SetReadDeadline(time.Now().Add(30*time.Second)) //30s后conn.Read会报出i/o timeout
	defer conn.Close()
	for {//长连接，即连接建立后进行多轮的读写交互
		requestBytes:=make([]byte,256)//初始化后byte数组每个元素都是0
		read_len,err:=conn.Read(requestBytes)
		if err!=nil{
			fmt.Printf("read from socket error:%s\n",err.Error())
			break //到达deadline后，退出for循环，关闭连接。client再用这个连接读写会发生错误
		}
		fmt.Printf("receive request %s\n",string(requestBytes))

		var request common.Request
		json.Unmarshal(requestBytes[:read_len],&request)
		response:=common.Response{Sum: request.A+request.B}

		responseBytes,_:=json.Marshal(response)
		_,err=conn.Write(responseBytes)
		common.CheckError(err)
		fmt.Printf("write response %s\n",string(responseBytes))

	}
}


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
		go handleRequest2(conn)
	}

}
