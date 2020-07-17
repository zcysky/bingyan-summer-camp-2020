package main

import (
	_ "bufio"
	"fmt"
	"github.com/jackdanger/collectlinks"
	_ "github.com/jackdanger/collectlinks"
	_ "io"
	"net/http"
	"os"
	_ "os"
	"strings"
)

func main() {
	url := "https://www.bilibili.com/anime/index/"
	download(url)
}

func tracefile(str_content string)  {
	fd,_:=os.OpenFile("Crawler.md",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	fd_content:=strings.Join([]string{str_content,"\n"},"")
	buf:=[]byte(fd_content)
	fd.Write(buf)
	fd.Close()
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

	links := collectlinks.All(resp.Body)
	for _, link := range links {
		content := "parse url"+link
		tracefile(fmt.Sprintf("%s",content))
		//fmt.Println("parse url", link)
	}

	//fmt.Println(string(body))
}