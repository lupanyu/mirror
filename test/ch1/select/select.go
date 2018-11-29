package main

import (
	"fmt"
	"time"
	"math/rand"
)
func worker(id int,c chan int){
	for n := range c{                           	//取c中的值如果没有则阻塞
		fmt.Printf("Worker %d received %d\n",id,n)   //有io输出
	}
}
func createWorker(id int) chan<- int{
	c := make(chan int) //创建个新的chan
	go worker(id,c)		//开启协程执行 并返回 c
	return c
}

func generator() chan int{
	out := make(chan int)
	go func(){
		i:=0
		for {
			time.Sleep(time.Duration(rand.Intn(1500))* time.Millisecond)
			out <- i
			i ++
		}
	}()
	return out
}
func main() {
	var c1, c2 = generator() ,generator()
	var worker = createWorker(0)
	//w := createWorker(0)
	n := 0
	hasValue := false
	for {
		var activeWorker chan int
		if hasValue{
			activeWorker = worker
		}
		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
			case activeWorker <- n:
				hasValue = false
time.Tick()		}
	}
}
