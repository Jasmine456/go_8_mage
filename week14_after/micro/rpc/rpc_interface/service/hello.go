package service

import (
	"bytes"
	"encoding/gob"
)

//通过这个接口约束客户端的调用和服务端的实现
//只要该接口公开，是不是对于client就完全知道该如何使用该RPC
//client.Greet("alice",&resp)
type  HelloService interface{
	Greet(string,*string) error
}

func GobEncode(val any) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	encoder := gob.NewEncoder(buf)
	if err := encoder.Encode(val); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func GobDecode(data []byte, value any) error {
	reader := bytes.NewReader(data)
	decoder := gob.NewDecoder(reader)
	return decoder.Decode(value)
}