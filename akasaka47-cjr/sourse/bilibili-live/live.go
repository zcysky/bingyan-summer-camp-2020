package main

import (
	_ "context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	_ "log"
	"net/http"
	"strconv"
)

type Respon struct {
	Id     int    `json:"id"`
	Head   string `json:"head"`
	UID    string `json:"uid"`
	Status bool   `json:"status"`
	Infom  string `json:"info"`
}

func Getinfo(url string, r Rooms, db *mongo.Collection) {
	fmt.Println("get-" + url)
	//fmt.Println(r.Id)
	var Liveinfo Respon
	for i, id := range r.Id {
		urls := url + strconv.Itoa(id)
		i++
		fmt.Println(i, urls)
		Liveinfo.Id = id
		req, err := http.NewRequest(http.MethodGet, urls, nil)
		if err != nil {
			fmt.Println("error", err)
		}
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
		req.Header.Add("Accept", "text/html")
		//fmt.Println(req)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("error", err)
		}
		//fmt.Println(res)
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			fmt.Println("error", err)
		}
		doc.Find("#link-app-title").Each(func(c int, selection *goquery.Selection) {
			//fmt.Println(selection.Text())
			str := selection.Text()
			//var stringBuilder bytes.Buffer
			temp := -1
			sta := 0
			for _, r := range str {
				//fmt.Printf("%d\t%q\t%d\n", i, r, r)
				//fmt.Printf("%q", r)
				if r == 45 {
					sta = 1
					break
				}
				temp++
			}
			//fmt.Println(temp)
			if sta == 1 {
				newtitle := str[:temp*3]
				Liveinfo.Head = newtitle
				fmt.Println("直播间名称：" + newtitle)
				rest := str[temp*3+3:]
				Liveinfo.Status = true
				temp = -1
				for _, r := range rest {
					if r == 45 {
						break
					}
					temp++
				}
				//fmt.Println(rest)
				tuber := rest[:3*temp]
				fmt.Println("主播：" + tuber)
				Liveinfo.UID = tuber
				fmt.Println("状态：正在直播")
			} else {
				Liveinfo.Head = "none"
				Liveinfo.UID = "none"
				Liveinfo.Status = false
				fmt.Println("状态：未开播")
			}
			//fmt.Println(Liveinfo)
		})
		//doc.Find("div[class = script-requirement]")
	}
}
