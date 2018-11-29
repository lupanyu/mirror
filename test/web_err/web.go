package main

import (
	"net/http"
	"mirror/test/web_err/filelist"
	"os"
	"github.com/gpmgo/gopm/modules/log"
	"fmt"
)

type userError interface {
	error
	Message() string
}

type appHandle func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandle) func(w http.ResponseWriter,r *http.Request){
	return func(writer http.ResponseWriter, request *http.Request){

		defer func() {
			if r := recover(); r != nil {
				log.Print(1, "Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		fmt.Println(request.URL.Path)
		err := handler(writer,request)
		if err != nil {
			fmt.Println("31 --> Error occurred "+
				"handling request: ",
				err.Error())
			code := http.StatusOK
			if userError,ok := err.(userError); ok {
				fmt.Println("35")
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return

				switch {
				case os.IsNotExist(err):
					code = http.StatusNotFound
				case os.IsPermission(err):
					code = http.StatusForbidden
				default:
					code = http.StatusInternalServerError

				}
			}
			http.Error(writer,http.StatusText(code),code)
		}
	}
}

func main(){
	http.HandleFunc("/list/", errWrapper(filelist.HandleFilelist))
	err := http.ListenAndServe(":8888",nil)
	if err!= nil {
		panic(err)
	}
}
