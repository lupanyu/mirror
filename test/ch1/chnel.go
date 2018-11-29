package main

import "fmt"

func chanDemo(){
	c := make(chan int)
	c <- 1
	c <- 2
	n := <- c
	fmt.Println(n)
}
