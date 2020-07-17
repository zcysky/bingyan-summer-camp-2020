package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://www.bilibili.com/anime/index/"
	download(url)
}

func download(url string) {
	//管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client
	client := &http.Client{}
	//提交请求
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header 设置方法：req.Header.Set("User-Agent","自定义的浏览器")
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	//执行
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	//函数结束后关闭相关链接
	defer resp.Body.Close()

	//读取文件
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	fmt.Println(string(body))
}