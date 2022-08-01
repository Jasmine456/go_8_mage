package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main(){
	router:=httprouter.New()
	router.ServeFiles("/file/*filepath",http.Dir("../../../static/site_a"))
	err:=http.ListenAndServe(":5666",router)
	if err!=nil{
		fmt.Println(err)
	}
}
