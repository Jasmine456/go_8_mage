package main

import (
	"fmt"
	"go_8_mage/week14_after/micro/rpc/protobuf/sample/pb"
	"google.golang.org/protobuf/types/known/anypb"
)

type Blog struct {
	Status pb.STATUS
}

func main() {
	var a pb.STATUS
	a = pb.STATUS_DRAFT
	a = pb.STATUS_PUBLISHED
	fmt.Println(a)

	//e:=pb.Event{
	//	Type: pb.TYPE_RESOURCE_ALERT,
	//	Data: &pb.Event_ResourceChange{},
	//}
	//
	//switch e.Type {
	//case pb.TYPE_RESOURCE_ALERT:
	//	e.GetResourceAlert()
	//case pb.TYPE_RESOURCE_CHANGE:
	//	e.GetResourceChange()
	//}


	// any的传参
	any,err:=anypb.New(&pb.Blog{
		Title: "test",
		Author: "jasmine",
	})
	if err != nil {
		panic(err)
	}
	device := pb.Device{
		DeviceType: "a1",
		//不能直接把对象传递进去
		Data: any,
	}
	fmt.Println(device)

	//如何解析对象
	blogA1:=&pb.Blog{}
	err=device.Data.UnmarshalTo(blogA1)
	if err != nil {
		panic(err)
	}
	fmt.Println(blogA1)
}
