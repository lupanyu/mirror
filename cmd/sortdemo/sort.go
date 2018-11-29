package main

import (
	"sort"
	"fmt"
)

func main(){
	a := []int{1,3,4,5,43,5,43,6,547,65}
	sort.Ints(a)
	for i,v := range a {
		fmt.Println(i,v)
	}
}
