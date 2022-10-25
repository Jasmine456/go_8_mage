package pb_test

import (
	"fmt"
	"go_8_mage/week14_after/micro/rpc/protobuf/hello/pb"
	"testing"
	"google.golang.org/protobuf/proto"
)

func TestMarshal(t *testing.T){
	b:=&pb.Blog{
		Title: "Go 语言Protobuf讲解",
		Summary:"Go 语言Protobuf讲解",
		Content: "xxx",
		Author: "jasmine",
		IsPublished: true,
	}

	// 序列化（编码）
	encodeB,err:=proto.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(encodeB))

//	解码
	b2:= &pb.Blog{}
	if err:= proto.Unmarshal(encodeB,b2);err!=nil{
		t.Fatal(err)
	}
	fmt.Println(b2)


}
