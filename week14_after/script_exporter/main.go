package main

import (
	"fmt"
	"go_8_mage/week14_after/script_exporter/excutor/script"
	"net/http"
	"github.com/"
)

func ScriptHandler(w http.ResponseWriter,r *http.Request){
	sc:=script.NewExcutor("modules")
	qs:=r.URL.Query()
	err:=sc.Exec(qs.Get("module"),qs.Get("params"),w)
	if err!=nil{
		fmt.Fprintf(w,fmt.Sprintf("exec error,%s",err))
	}
}

func main(){
	zap.DevelopmentSetup()
	http.HandleFunc("/metrics",ScriptHandler)
	http.ListenAndServe(":8050",nil)
}