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
		server ,err := lib.Open(conf.Host+string(conf.Port))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(server)
		lib.Server()
	}else{
		//开启客户端服务
		client,err := lib.Open(conf.Host+string(conf.Port))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(client)
		lib.Client()
	}
    }
