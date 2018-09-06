package main

import (
	"net"
	"log"
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"

)

func load_conf() ( root  ,auth string, auto bool,host  string,port int64){
	config,err := yaml.ReadFile("mirror.yml")
	if err != nil {
		log.Fatal(err)
	}
	root,_  = config.Get("root")
	auth,_  = config.Get("auth")
	auto,_   = config.GetBool("auto")
	host, _  = config.Get("host")
	port, _  = config.GetInt("port")
	return root,auth,auto,host,port
}
type data struct{
	data_type string
	json      string
}

type conf struct {
	root string
	auth string
	auto bool
	Host string
	Port int64
}


type server interface {
	receive()  //接收数据
	auth( )    //认证
}
type client interface {
	conn(addr string)  //连接
	send(data) //发送数据
}

func x(a string){
	fmt.Println(a)
}

func remote( b  []string ,host string) ( bool){
	for i := 0; i < len(b); i ++ {
		if b[i] == host {
			return true
		}
	}
	return false
}

func ip() []string{
	//获取本机的所有ip地址
	local_ip , err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	var ip_slice []string
	for _, address := range local_ip {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//fmt.Println( ipnet.IP.String())
				ip_slice = append(ip_slice,ipnet.IP.String())
			}
		}
	}
	return  ip_slice
}
func main() {
	b := ip()
	fmt.Println(b)
	root,auth, auto,host,port := load_conf()
	fmt.Println(root,auth,auto,host,port)
	var L_conf conf
	L_conf = conf{root,auth,auto,host,port}
	if remote(b,host) {

		listener,err := net.Listen("tcp",L_conf.Host+string(L_conf.Port))
		if err != nil {
			log.Fatal(err)
		}
		listener.Accept()
	}

}

