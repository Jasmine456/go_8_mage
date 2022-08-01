package main

import (
	"fmt"
	"sync"
)

/*
数组、slice、struct 允许并发修改（可能会脏写），并发修改map会发生panic，因为map的value是不可寻址的
如果需要并发修改map，请使用sync.Map
*/

type Student struct {
	Name string
	Age  int32
}

var arr = [10]int{}
var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	//	测试数组并发添加
	go func() {
		defer wg.Done()
		for i := 0; i < len(arr); i += 2 {
			arr[i] = 0
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < len(arr); i += 2 {
			arr[i] = 1
		}
	}()
	wg.Wait()
	fmt.Println(arr)
	fmt.Println("=====================")

	//	测试struct 结构体并发添加
	wg.Add(2)
	var stu Student
	go func() {
		defer wg.Done()
		stu.Name = "Fred"
	}()
	go func() {
		defer wg.Done()
		stu.Age = 20
	}()
	wg.Wait()
	fmt.Printf("%s %d\n", stu.Name, stu.Age)
	fmt.Println("=================")

	//	测试map并发添加
	wg.Add(2)
	go func(){
		defer wg.Done()
		m.Store("k1","v1")
	}()
	go func(){
		defer wg.Done()
		m.Store("k1","v2")
	}()
	wg.Wait()
	fmt.Println(m.Load("k1"))

}
