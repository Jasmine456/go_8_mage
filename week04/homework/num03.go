package main

import "fmt"

/*
定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口

 */
type Fisher interface {
	swim()
}
type Crawler interface {
	crawl()
}


type Wuzhi interface { //空接口

}
//struct 理解为类
type Frog struct {
	Height float64//成员变量
	Local string
}

func (frog *Frog)swim(){
	fmt.Println("I will swim")
	fmt.Println(frog.Local)
	frog.Height+=10
}

func (frog Frog)crawl(){
	fmt.Println("I will crawl")
}


func main(){
	frog:=Frog{}
	var fs Fisher
	fs = &frog

	var cl Crawler
	cl=frog
	cl=&frog

	fs.swim()
	cl.crawl()

	frog.swim()

	var wz interface{}
	wz=frog
	var a int
	var b string
	wz=a
	wz=b
	fmt.Println(wz)


}