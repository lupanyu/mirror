package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"bytes"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"regexp"
	"fmt"
)

const homeUrl  =  "http://www.biquge.com.tw/12_12489/"

type zhang struct {
	ID ,Title,SubUrl string  //章节，章节名，章节url
}

func (z *zhang)Body(){
	HttpGetHandler(z.SubUrl)
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func HttpGetHandler(url string)([]byte,error){
	request ,_ := http.NewRequest(http.MethodGet,url,nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	resp,err := http.DefaultClient.Do(request)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode  == http.StatusOK {
		bytes, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			log.Fatal(e)
		}
		utf8, err := GbkToUtf8(bytes)
		if err != nil {
			log.Fatal(err)
		}
		return utf8,nil
	}
	return nil,nil
}

func main() {
	str, e := HttpGetHandler(homeUrl)
	if e != nil {
		log.Fatal("get home err :",e)
	}
	//分组引用  链接 为1 章节名为2
	bodyRegStr := `<dd><a href="(.*?)">(.*?)</a></dd>`
	bodyReg := regexp.MustCompile(bodyRegStr)
	fmt.Println(string(str))
	all := bodyReg.FindAllSubmatch(str,-1)
	f:=all[0][1]   // 每一个符合 reg的第一个（）中的数据 uri ,同理第2个为[2]
	fmt.Println(string(f))



}
