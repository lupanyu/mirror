package main

import (
	"reflect"
	"fmt"
)

type art struct {
	A string
	B []byte
	C bool
	D map[string]string
	E art
}

type X interface {

}

func main(){
	b := 1
	a := art{}

	fmt.Println( reflect.TypeOf(a))
}