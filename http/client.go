package main

import (
	"net/http"
	"fmt"
)
func header(){
	request ,_ := http.NewRequest(http.MethodGet,"http://www.imooc.com",nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	resp,err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Header.Get("User-Agent"))

	}

func main() {
	header()
	}
