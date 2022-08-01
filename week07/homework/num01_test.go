package main

import (
	"fmt"
	"testing"
)

func TestBsEqual2(t *testing.T) {
	arr:=make([]byte,0)
	brr:=make([]byte,0)
	fmt.Printf("%p \n%p\n",&arr,&brr)
	fmt.Println(BsEqual2(arr,brr))
}
