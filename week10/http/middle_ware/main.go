package main

import (
	"log"
	"net/http"
	"time"
)

var limitCh = make(chan struct{}, 100)

func getBoy(w http.ResponseWriter, r *http.Request) {
	time.Sleep(150 * time.Millisecond)
	w.Write([]byte("hi boy"))
}

func getGirl(w http.ResponseWriter, r *http.Request) {
	time.Sleep(150 * time.Millisecond)
	w.Write([]byte("hi girl"))
}

func getMan(w http.ResponseWriter, r *http.Request) {
	time.Sleep(150 * time.Millisecond)
	w.Write([]byte("hi man"))
}

//中间件逻辑
//超时中间件
func timeMiddleWare(net http.Handler) http.Handler {
	// 通过HandlerFunc把一个 func(rw http.ResponseWriter,r *http.Request) 函数转为Handler
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		//time.Sleep(2*time.Second)
		net.ServeHTTP(rw, r)
		timeElapsed := time.Since(begin)
		log.Printf("request %s use %d ms\n", r.URL.Path, timeElapsed.Milliseconds())
	})
}

//限流中间件
func limitMiddleWare(net http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		limitCh <- struct{}{}
		log.Printf("concurrence %d\n", len(limitCh))
		net.ServeHTTP(rw, r)
		<-limitCh
	})
}

/*
更优雅的使用中间件组织形式
*/
type middleware func(handler http.Handler) http.Handler

type Router struct {
	middlewareChain []middleware
	mux             map[string]http.Handler //mux通常表示路由策略
}

//Router的构造函数
func NewRoute() *Router {
	return &Router{
		middlewareChain: make([]middleware, 0, 10),
		mux:             make(map[string]http.Handler, 10),
	}
}

func (router *Router) Use(m middleware) {
	router.middlewareChain = append(router.middlewareChain, m)
}

func (router *Router) Add(path string, handler http.Handler) {
	var mergedHandler = handler
	//var mergedHandler http.Handler
	for i := (len(router.middlewareChain) - 1); i >= 0; i-- {
		mergedHandler = router.middlewareChain[i](mergedHandler) //中间件层层嵌套
	}
	router.mux[path] = mergedHandler
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.RequestURI()
	//	在Handler内部实现路由
	if handler, exists := router.mux[requestPath]; !exists {
		http.NotFoundHandler().ServeHTTP(w, r)
	} else {
		handler.ServeHTTP(w, r)
	}
}

func main() {
	//第一种，不带中间件的路由
	//http.HandleFunc("/getBoy", getBoy) // 不带中间件的路由
	//http.HandleFunc("/getGirl", getGirl)
	//http.HandleFunc("/getMan", getMan)

	//第二种，带了中间件，但是代码冗余
	//http.Handle("/getBoy", limitMiddleWare(timeMiddleWare(http.HandlerFunc(getBoy))))  //中间层层嵌套，代码重复过多
	//http.Handle("/getBoy", limitMiddleWare(timeMiddleWare(http.HandlerFunc(getGirl)))) //带中间件的路由
	//http.Handle("/getBoy", limitMiddleWare(timeMiddleWare(http.HandlerFunc(getMan))))
	//
	//http.ListenAndServe(":5656", nil)

	//	第三种，带了中间件，用了面向对象 面向接口的思想，减少代码冗余
	router := NewRoute()
	router.Use(limitMiddleWare)
	router.Use(timeMiddleWare)
	//	以下演示了2个路径（还可以更多），每个路径都使用相同的middlewareChain
	router.Add("/boy", http.HandlerFunc(getBoy))
	router.Add("/girl", http.HandlerFunc(getGirl))
	router.Add("/man", http.HandlerFunc(getMan))
	http.ListenAndServe(":5656", router)

}

//02:15:20
