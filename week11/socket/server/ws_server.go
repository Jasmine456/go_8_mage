package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go_8_mage/week11/socket/common"
	"net"
	"net/http"
	"time"
)


func boy(w http.ResponseWriter,r *http.Request){
	upgrade := &websocket.Upgrader{
		HandshakeTimeout: 5*time.Second,//握手超时时间
		ReadBufferSize: 2048, //读缓冲大小
		WriteBufferSize: 1024,
	}
	conn,err:=upgrade.Upgrade(w,r,nil) //将http协议升级到websocket协议
	if err !=nil{
		fmt.Printf("把http升级为websocket时失败: %v\n",err)
		return
	}
	defer conn.Close()

	for{
		conn.SetReadDeadline(time.Now().Add(20*time.Second))
		var request common.Request
		if err:=conn.ReadJSON(&request);err!=nil{
			//判断是不是超时
			if netError,ok:=err.(net.Error);ok{ //如果ok==true，说明类型断言成功
				if netError.Timeout(){
					fmt.Printf("read message timeout,remote %s\n",conn.RemoteAddr())
					return
				}
			}
			//	忽略websocket.CloseGoingAway/websocket.CloseNormalClosure这2中closeErr，如果是其他closeErr就打一条错误日志
			if websocket.IsUnexpectedCloseError(err,websocket.CloseGoingAway,websocket.CloseNormalClosure){
				fmt.Printf("read message from %s error %v\n",conn.RemoteAddr().String())
			}
			return //只要ReadMessage发生错误，就关闭这条连接
		} else {
			response := common.Response{Sum: request.A+request.B}
			if err = conn.WriteJSON(&response);err!=nil{
				fmt.Printf("write response failed: %v\n",err)
			} else{
				fmt.Printf("write response %d\n",response.Sum)
			}
		}
	}
}


func main(){
	http.HandleFunc("/",boy)
	if err:= http.ListenAndServe("127.0.0.1:5657",nil); err !=nil{
		fmt.Println(err)
	}
}