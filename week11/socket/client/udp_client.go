package main

import (
	"encoding/json"
	"fmt"
	"go_8_mage/week11/socket/common"
	"net"
	"strconv"
	"sync"
	"time"
)

//长连接
func main(){
	ip:="127.0.0.1"
	port:=5656
	//跟tcp_client的唯一区别就是这行代码
	conn,err:=net.DialTimeout("udp",ip+":"+strconv.Itoa(port),30*time.Minute)//一个conn绑定一个本地端口
	common.CheckError(err)
	defer conn.Close()
	const P=2
	wg:=sync.WaitGroup{}
	wg.Add(P)
	for i:=0;i<P;i++{
		request:=common.Request{A:12,B: 16}
		requestBytes,_:=json.Marshal(request)
		go func(){
			defer wg.Done()
			for { //长连接，即建立连接后进行多轮的读写交互
				_,err=conn.Write(requestBytes)
				common.CheckError(err)
				fmt.Printf("write request %s\n",string(requestBytes))
				responseBytes:=make([]byte,256) //初始化后byte数组每个元素都是0
				read_len,err:=conn.Read(responseBytes)
				common.CheckError(err)

				var response common.Response
				//fmt.Println(responseBytes)
				json.Unmarshal(responseBytes[:read_len],&response)
				fmt.Printf("receive response: %d\n",response.Sum)
				time.Sleep(1*time.Second)
			}
		}()
	}
	wg.Wait()

}
