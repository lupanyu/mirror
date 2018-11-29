package main

import (
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer,"<h1>hello world</h1>")
	})
	http.ListenAndServe("0.0.0.0:8888",nil)
}
