package main

import (
	"mirror/lib"
	"log"
	"fmt"
)

func main() {
	root,auth, auto,host,port := load_conf()
	conf := lib.Conf{root,auth, auto,host,port }
    iplist := lib.Ip()
	if remote(iplist,host) {
		//如果本机ip列表中和 配置中匹配 开启服务端
		buff ,err := lib.Server(conf.Host+string(conf.Port))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(buff)
	}else{
		//
	}
    }
