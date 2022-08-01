package main

import (
	//"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go_8_mage/week11/socket/common"
	"io/ioutil"
	"net/http"
	"time"
)

func main(){
	dialer:=&websocket.Dialer{}
	header:=http.Header{
		"Cookie": []string{"name=jasmine"},
	}
	conn,resp,err:=dialer.Dial("ws://localhost:5657/add",header)
	//defer resp.Body.Close()
	if err!=nil{
		fmt.Printf("dial server error:%v\n",err)
		fmt.Println(resp.StatusCode)
		msg,_:=ioutil.ReadAll(resp.Body)
		fmt.Println(string(msg))
		return
	}
	fmt.Println("handshake response header")
	for key,values:=range resp.Header{
		fmt.Printf("%s:%s\n",key,values[0])
	}
	defer conn.Close()
	for i:=0;i<10;i++{
		request:=common.Request{A: 12,B: 16}
		//requestBytes,err:=json.Marshal(request)
		common.CheckError(err)
		err=conn.WriteJSON(request)//websocket.Conn直接提供发json序列化和反序列化方法
		common.CheckError(err)
		fmt.Printf("write request %s\n",request)
		var response common.Response
		err = conn.ReadJSON(&response)
		common.CheckError(err)
		fmt.Printf("receive response: %d\n",response.Sum)
		time.Sleep(1*time.Second)
	}
	time.Sleep(30*time.Second)
}
