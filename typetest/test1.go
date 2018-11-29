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
	fmt.Println( reflect.TypeOf(1))

}