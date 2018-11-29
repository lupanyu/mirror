package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	resp,err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		return
	}
	all ,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s ",all)
	}
}
