package service

//通过这个接口约束客户端的调用和服务端的实现
//只要该接口公开，是不是对于client就完全知道该如何使用该RPC
//client.Greet("alice",&resp)
type  HelloService interface{
	Greet(string,*string) error
}
