package main

import "fmt"

func main(){
	var m map[string]int
	m = make(map[string]int)
	m=make(map[string]int,200)
	m=map[string]int{"语文":0,"数学":39}

//	add delete
	m["英语"]=59
	m["英语"]=70
	delete(m,"数学")

	if value,exists := m["语文"];exists{
		fmt.Println(value)
	}else{
		fmt.Println("map里不存在[语文]这个key")
	}

//	range
	for key,value:=range m{
		fmt.Printf("%s:%d",key,value)
	}

}
