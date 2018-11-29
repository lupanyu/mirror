package  main

import (
	"mirror/test/jiqiren"
	"fmt"
	"mirror/test/zhenren"
)

type abc interface {
	Hello(string)string
}

func  hi( m abc)string {
	return m.Hello("")
}

func main(){
	var panyu jiqiren.Man
	panyu2 := zhenren.Man{"panyu" ,18 }
	fmt.Println(panyu2.Hello("lalalal"))
	fmt.Println(panyu.Hello("????"))
}
