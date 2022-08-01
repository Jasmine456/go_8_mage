package main

import (
	"encoding/json"
	"fmt"
	"go_8_mage/week11/socket/common"
	"net"
	"strconv"
	"time"
)

func main() {

	ip := "127.0.0.1"
	port := 5656
	conn, err := net.DialTimeout("tcp4", ip+":"+strconv.Itoa(port), 30*time.Minute)
	common.CheckError(err)
	fmt.Printf("establish connection to server %s\n", conn.RemoteAddr().String())
	defer conn.Close()
	for {//長连接，即连接建立后进行多轮的读写交互
		request:=common.Request{A: 15,B: 12}
		requestBytes,_:=json.Marshal(request)
		_,err=conn.Write(requestBytes)
		common.CheckError(err)
		fmt.Printf("write request %s\n",string(requestBytes))
		responseBytes:=make([]byte,256)
		read_len,err:=conn.Read(responseBytes)
		common.CheckError(err)
		var response common.Response
		json.Unmarshal(responseBytes[:read_len],response)
		fmt.Printf("receive response: %d\n",response.Sum)
		time.Sleep(1*time.Second)
	}


}
