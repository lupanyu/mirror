package main

import (
	"fmt"
)

type worker struct {
	in chan int
	done chan bool
}

func doWWorker(id int,c chan int,done chan bool ){
	for n:= range c{
		fmt.Printf("worker %d received %c\n",id,n)
		done <- true
	}
}

func createWorker(id int)worker {
	w := worker{
		in: make(chan int),
		done: make(chan bool),
	}
	go doWWorker(id,w.in,w.done)
	return w
}

func chanDemo(){
	var workers [10]worker
	for i:=0 ; i< 10 ; i ++ {
		workers[i] = createWorker(i) //
	}
}
func main(){
	chanDemo()
}

