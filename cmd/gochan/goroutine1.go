package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan string)
	for i := 0 ; i < 10 ; i ++ {
		go helloworld(i,ch)
	}
	time.Sleep(time.Second)
	for {
		msg := <- ch
		fmt.Print(msg)
	}
}

func helloworld(i int ,ch chan string){
	ch <- fmt.Sprintf("hello world ! %d\n",i)
}
