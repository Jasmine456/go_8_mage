package main

import (
	"encoding/json"
	"fmt"
	"go_8_mage/week11/socket/common"
	"net"
	"strconv"
	"time"
)

//长连接
func main(){
	ip:="127.0.0.1"
	port:=5656
	udpAddr,err:=net.ResolveUDPAddr("udp",ip+":"+strconv.Itoa(port))
	common.CheckError(err)
	conn,err:=net.ListenUDP("udp",udpAddr) //UDP不需要创建连接，所以不需要向TCP那样通过Accept() 创建连接，这里的conn 是个假链接
	common.CheckError(err)
	conn.SetReadDeadline(time.Now().Add(30*time.Second))
	defer conn.Close()
	for {
		requestBytes:=make([]byte,256) //初始化后的数组每个元素都是0
		read_len,remoteAddr,err:=conn.ReadFromUDP(requestBytes) //一个conn可以对应多个client，ReadFrom可以返回是哪个
		if err!=nil{
			fmt.Printf("read from socket error:%s\n",err.Error())
			break //到达deadline后，退出for循环，关闭连接。client再用这个连接读写会发生错误
		}
		fmt.Printf("receive request %s from %s\n",string(requestBytes),remoteAddr.String())//[]byte转string时，0后面的会自动被借调

		var request common.Request
		json.Unmarshal(requestBytes[:read_len],&request) //json序列化时会把0都考虑在内，所以需要指定只读前read_len个字节
		response:=common.Response{Sum: request.A+request.B}

		responseBytes,_:=json.Marshal(response)
		_,err=conn.WriteToUDP(responseBytes,remoteAddr)// 由于UDP conn支持多对多通信，所以通信对方可能有多个EndPoint，通过WriteTo指定要写给那个EndPoint
		common.CheckError(err)
		fmt.Printf("write response %s to %s\n",string(responseBytes),remoteAddr.String())

	}



}
