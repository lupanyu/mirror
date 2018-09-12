package lib

import (
	"net"
	"bufio"
	"encoding/gob"
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
	"io/ioutil"
	"os"
)

/**传输文件过程
1、先验证token         返回 True/Flase
2、client发送文件信息
3、server验证是否一致  返回 True/Flase    提取文件路径和名
4、客户端传文件开始    发送1   /如果返回Flase
5、客户端传内容        发送D
6、客户端传结束标志    发送0
7、服务端保存文件      发送True

文件扫描完毕，客户端 断开连接
**/
type ComplexData struct{
	A string              // 验证 等字符串
	M map[string]string  //文件名：md5
	D []byte 			  //文件 数据
	B bool 			      //文件是否一致 一致就不传
	I int 				  //标记 开始传输 1 结束传输 0
	C *ComplexData
}
type Conf struct {
	root string
	auth string
	auto bool
	Host string
	Port int64
}
func Load_conf() ( root  ,auth string, auto bool,host  string,port int64){
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
func Ip() []string{
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
/**
	net.Conn 实现了io.Reader  io.Writer  io.Closer接口
	Open 返回一个有超时的TCP链接缓冲readwrite
 */
func Open(addr string)(*bufio.ReadWriter,error){
	conn,err := net.Dial("tcp",addr)
	//开启监听
	if err != nil {
		return nil, err
	}
	return bufio.NewReadWriter(bufio.NewReader(conn),bufio.NewWriter(conn)), nil
	//返回一个io缓冲器
}
//解析交互数据
func HandleGob(rw *bufio.ReadWriter) ( ComplexData,error){
	var data ComplexData
	dec := gob.NewDecoder(rw)
	err := dec.Decode(&data)

	if err != nil {
		fmt.Println("无法解析的数据.")
		return data,err
	}
	return data,nil
}

func Write(filename string,bufdata []byte,pem os.FileMode ) int{
	err :=ioutil.WriteFile(filename,bufdata,pem)
	if err !=nil {
		return 0
	}else {
		return 1
	}
}
func Server(data ComplexData){

	switch HandleGob(data) {

	}

}

func Client(){

}