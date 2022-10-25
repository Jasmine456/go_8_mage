package service_test

import (
	"encoding/json"
	"go_8_mage/week14_after/micro/rpc/rpc_interface/service"
	"testing"
)


type Blog struct {
	Title string
	Author string
	Summary string
	IsPublished bool
}

func TestGob(t *testing.T){
	b:=&Blog{
		Title:"test",
		Author: "jasmine",
		Summary: "J",
		IsPublished: true,
	}
	encodeB,err:=service.GobEncode(b)
	if err != nil {
		t.Fatal(err)
	}

	jsonB,err:=json.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(encodeB),"\n",len(jsonB))
}
