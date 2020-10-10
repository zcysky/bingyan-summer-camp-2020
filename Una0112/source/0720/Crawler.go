package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
)

func GetMovie(url string) {
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//bodyString, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bodyString))
	if resp.StatusCode != 200 {
		fmt.Println("err")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	//

	doc.Find("#content h1").Each(func(i int, s *goquery.Selection) {
		// name
		fmt.Println("name:" + s.ChildrenFiltered(`[property="v:itemreviewed"]`).Text())
		// year
		fmt.Println("year:" + s.ChildrenFiltered(`.year`).Text())
	})

	// #info > span:nth-child(1) > span.attrs
	director := ""
	doc.Find("#info span:nth-child(1) span.attrs").Each(func(i int, s *goquery.Selection) {
		// 导演
		director += s.Text()
		//fmt.Println(s.Text())
	})
	fmt.Println("导演:" + director)
	//fmt.Println("\n")

	pl := ""
	doc.Find("#info span:nth-child(3) span.attrs").Each(func(i int, s *goquery.Selection) {
		pl += s.Text()
	})
	fmt.Println("编剧:" + pl)

	charactor := ""
	doc.Find("#info span.actor span.attrs").Each(func(i int, s *goquery.Selection) {
		charactor += s.Text()
	})
	fmt.Println("主演:" + charactor)

	typeStr := ""
	doc.Find("#info > span:nth-child(8)").Each(func(i int, s *goquery.Selection) {
		typeStr += s.Text()
	})
	fmt.Println("类型:" + typeStr)
}

func GetToplist(url string) []string {
	var urls []string
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//bodyString, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bodyString))
	if resp.StatusCode != 200 {
		fmt.Println("err")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	doc.Find("#content div div.article ol li div div.info div.hd a").Each(func(i int, s *goquery.Selection) {
		// year
		fmt.Printf("%v", s)
		herf, _ := s.Attr("href")
		urls = append(urls, herf)
	})
	return urls
}

func main() {
	url := "https://movie.douban.com/top250?start="
	var urls []string
	var newUrl string
	fmt.Println("%v", urls)
	for i := 0; i < 10; i++ {
		start := i * 25
		newUrl = url + strconv.Itoa(start)//将数字转化为字符串
		urls = GetToplist(newUrl)

		for _, url := range urls {
			GetMovie(url)
		}
	}
}