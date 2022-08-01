package main

import (
	"encoding/json"
	"fmt"
	"go_8_mage/week11/socket/common"
	"net"
	"strconv"
	"time"
)

func request(){
	ip:="127.0.0.1"
	port:=5656
	conn,err:=net.DialTimeout("tcp4",ip+":"+strconv.Itoa(port),30*time.Minute)
	common.CheckError(err)
	fmt.Printf("establish connection to server %s\n",conn.RemoteAddr().String())
	request:=common.Request{A:7,B:4}
	requestBytes,_:=json.Marshal(request)
	_,err=conn.Write(requestBytes)
	common.CheckError(err)
	fmt.Printf("write request %s\n",string(requestBytes))
	responseBytes:=make([]byte,256) //初始化后byte数组每个元素都是0
	read_len,err := conn.Read(responseBytes)
	common.CheckError(err)
	//fmt.Println(responseBytes)
	var response common.Response
	err=json.Unmarshal(responseBytes[:read_len],&response) //json反序列化时会把0都考虑在内，所以需要指定只读前read_len个字节
	common.CheckError(err)
	fmt.Printf("receive response ：%d\n",response.Sum)
	conn.Close()
}
// 序列化请求和响应的结构体
func main(){
	go request()
	go request()
	go request()
	go request()

	time.Sleep(2*time.Second)
}

