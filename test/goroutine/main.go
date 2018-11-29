package main

import (
	"fmt"
	"time"
)

func main(){
	for i := 1 ; i < 10000 ; i ++ {
		go func(i int){
			fmt.Printf("hello from gorutine %d \n",i)
		}(i)
	}
	time.Sleep(time.Minute)
}
