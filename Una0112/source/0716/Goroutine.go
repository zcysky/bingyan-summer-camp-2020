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

	queue := make(chan string)
	go func() {
		queue <- url
	}()
	for uri := range queue {
		download(uri,queue)
	}
}

func tracefile(str_content string)  {
	fd,_:=os.OpenFile("Crawler.md",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	fd_content:=strings.Join([]string{str_content,"\n"},"")
	buf:=[]byte(fd_content)
	fd.Write(buf)
	fd.Close()
}

func download(url string,queue chan string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)
	for _, link := range links {
		content := "parse url"+link
		tracefile(fmt.Sprintf("%s",content))
		//fmt.Println("parse url", link)
		go func() {
			//link="https:"+link
			queue <- link
		}()
	}
}