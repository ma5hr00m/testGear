/**********************************************
	拷贝漫画官网: https://copymanga.site/
	go原生爬虫的使用，colly框架的使用
	获取拷贝漫画官网数据，边学边练，最后直接投入使用了
***********************************************/
package main

/*************************************************
	"log"
		OfficialDoc: https://pkg.go.dev/log
		Function:    日志库

	"net/http"
		OfficialDoc: https://pkg.go.dev/net/http
		Function:    发送http请求

	"io/ioutil"
		OfficialDoc: https://pkg.go.dev/io/ioutil
		Function:    解析response
*************************************************/
import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 定义所需常量
	const (
		requestURL = "https://copymanga.site/"
	)

	// 向目标网站发送GET请求
	rsp, err := http.Get(requestURL)

	// 若出现错误则打印报错日志
	if err != nil {
		log.Println(err.Error())
		return
	}

	//
	body, err := ioutil.ReadAll(rsp.Body)

	// 若出现错误则打印报错日志
	if err != nil {
		log.Println(err.Error())
		return
	}

	// 将返回报文转化成字符串储存在 content 变量中
	content := string(body)

	// 关闭请求
	defer rsp.Body.Close()

	log.Println(content)
}
