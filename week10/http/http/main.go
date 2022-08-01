package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func HelloHandler(w http.ResponseWriter,r *http.Request){
	/*
	具体看一下http协议
	 */
	fmt.Printf("request method: %s\n",r.Method)
	fmt.Printf("request host: %s\n",r.Host)
	fmt.Printf("request url: %s\n",r.URL)
	fmt.Printf("request proto: %s\n",r.Proto)
	fmt.Println("request header")
	for key,values := range r.Header{
		fmt.Printf("%s:%v\n",key,values)
	}
	fmt.Println()
	fmt.Println("request cookie")
	for _,cookie := range r.Cookies(){
		fmt.Printf("name=%s value=%s\n",cookie.Name,cookie.Value)
	}
	fmt.Println()
	fmt.Printf("request body:")
	io.Copy(os.Stdout,r.Body)
	fmt.Println()

	fmt.Fprint(w,"Hello Boy")
}

func main(){
	http.HandleFunc("/",HelloHandler)
	if err:=http.ListenAndServe(":5656",nil);err!=nil{
		panic(err)
	}
	/*
	有两种方案：
	1. 用http.Handle()实现路由，http.ListenAndServe()的Handler可以传nil
	2. 不用http.Handle()实现路由，http.ListenAndServe()时把Handler传进来，在该Handler内部实现路由
	 */
}
