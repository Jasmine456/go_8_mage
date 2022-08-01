package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

/*
1. 前端（browser)请求建立websocket连接时，为这条websocket连接专门开启一个协程，创建一个client。
2. client把前端发过来的数据写入hub的broadcast管道
3， client把自身send管道里的数据写给前端
4. client跟前端的连接断开是请求从hub那儿注销自己


存活检测
当Hub发现client的send管道写不进数据时，把client注销掉

client给websocket连接设置一个读超时，并周期性的给前端发ping消息，
如果没有收到pong消息则下一次的conn.read()会报出超时错误，
此时client关闭websocket连接
*/

type Client struct {
	hub       *Hub
	conn      *websocket.Conn
	send      chan []byte
	frontName []byte //前端的名字，用于展示在消息面前
}

var (
	newLine = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//从websocket连接里读出数据，发给hub
func (client *Client) read() {
	defer func() { //收尾工作
		client.hub.unregister <- client //从hub那注销自身的client
		fmt.Printf("%s offline\n", client.frontName)
		fmt.Printf("close connection to %s\n", client.conn.RemoteAddr().String())
		client.conn.Close() //关闭websocket管道
	}()
	for {
		_, message, err := client.conn.ReadMessage() //如果前端主动断开连接，该行会报错，for循环会退出。注销client时，hub那儿会关闭client.send管道
		if err != nil {
			break //只要ReadMessage失败，就关闭websocket管道、注销client，退出
		} else {
			//换行符用空格替代，bytes.TrimSpace把首尾连续的空格去掉
			message = bytes.TrimSpace(bytes.Replace(message, newLine, space, -1))
			if len(client.frontName) == 0 {
				client.frontName = message //约定：从浏览器读到的第一条消息代表前端的身份标识，该信息不进行广播
				fmt.Printf("%s online\n", string(client.frontName))
			} else {
				//要广播的内容前面加上front的名字
				client.hub.broadcast <- bytes.Join([][]byte{client.frontName, message}, []byte(": "))
			}
		}
	}
}

//从hub的broadcast那儿读数据，写到websocket连接里面去
func (client *Client) write() {
	defer func() {
		fmt.Printf("close connection to %s\n", client.conn.RemoteAddr().String())
		client.conn.Close() //给前端写数据失败，就可以关闭连接了
	}()

	for {
		msg, ok := <-client.send
		if !ok {
			fmt.Println("管道已经被关闭")
			client.conn.WriteMessage(websocket.CloseMessage, []byte("bye bye"))
			return
		} else {
			err := client.conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				fmt.Printf("向浏览器发送数据失败：%v\n", err)
				return
			}
		}
	}
}

//启动一个websocket server
func ServWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) //http升级为websocket协议
	if err != nil {
		fmt.Printf("upgrade error: %v\n", err)
		return
	}
	fmt.Printf("connect to client %s\n", conn.RemoteAddr().String())
	//	每来一个前端请求，就会创建一个client
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	//	向hub注册client
	client.hub.register <- client

	//	启动子协程，运行ServeWs的协程退出后子协程也不会退出
	go client.read()
	go client.write()
}
