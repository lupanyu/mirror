package main

import (
	"fmt"
	"errors"
)

type Man struct{
	Name string
	Age int
}

func (m Man)Hello(hi string)string{
	var data string
	data = "my name is " + m.Name + "my age is " + string(m.Age) + hi

	return data
}

func chifan() func(s string) string {
	defer fmt.Println("in chifan")
	result := "我有一只"
	return func(s string) string {
		result += s
		return result
	}

}

func testrecover(){

	defer func() {
		r := recover()
		if err,ok := r.(error); ok {
			fmt.Println()
		}
	}()
	panic(errors.New("test a panic"))
}
func main(){
	testrecover()
}