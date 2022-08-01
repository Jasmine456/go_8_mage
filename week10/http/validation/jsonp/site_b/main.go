package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main(){
	router := httprouter.New()
	router.ServeFiles("/file/*filepath",http.Dir("../../../static/site_b"))
	router.GET("/",func(w http.ResponseWriter,r *http.Request,p httprouter.Params){
		w.Write([]byte("yysd({'message':'Chinese'})")) //返回一段json回调函数
	})
	http.ListenAndServe(":5666",router)
}
