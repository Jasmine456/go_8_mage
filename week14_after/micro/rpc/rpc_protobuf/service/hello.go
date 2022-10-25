package service

// 通过这个接口约束客户端的调用和服务端的实现
// 只要该接口公开, 是不是对于client 就完全知道该如何使用该RPC
// client.Greet("alice", &resp)
// 发送的数据结构就是实现了 Protobuf 编解码的对象
type HelloService interface {
	Greet(*Request, *Response) error
}
