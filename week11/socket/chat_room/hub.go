package main

/*
功能解析：
1. Hub持有每一个Client的指针，broadcast管道里有数据时把它写入每一个Client的send管道中。
2. 注销Client时关闭Client的send管道
 */

type Hub struct {
	clients    map[*Client]bool //维护所有的Client
	broadcast  chan []byte      //需要广播的消息
	unregister chan *Client     //Client注销 销毁 请求通过管道来接收
	register   chan *Client     // Client注册请求通过管道来接收
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte), //同步管道，确保hub这里消息不会堆积。如果同时有多个client想给hub发送数据就阻塞
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case msg := <-hub.broadcast:
			for client := range hub.clients {
				client.send <- msg
			}
		case client := <- hub.register:
			hub.clients[client] = true //注册client
		case client := <- hub.unregister:
			if _, ok := hub.clients[client]; ok { //防止重复注销
				delete(hub.clients, client) //注销client
				close(client.send) //hub从此以后不需要再向该client广播消息了
			}
		}
	}
}
