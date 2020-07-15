/*Day2 A Simple WebCrawler*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type LiveInfo struct {
	ID       []int
	Interval int
}

type ResInfo struct {
	ID []int
	Title []string
	UserName []string
}

func Live(c *gin.Context) {
	// Get liveID from json file:
	var info LiveInfo
	err := ParseJson("config/config.json", &info)
	if err != nil {
		fmt.Println(err)
	}

	// Request http:
	var resInfo ResInfo
	for _, id := range info.ID {
		url := "https://live.bilibili.com/" + strconv.Itoa(id)
		fmt.Println(url)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("User-Agent", "PostmanRuntime/7.26.1")
		req.Header.Add("Accept", "*/*")
		res, err := http.DefaultClient.Do(req)
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		doc.Find("#link-app-title").Each(func(i int, s *goquery.Selection) {
			rawData := s.Text()
			data := strings.Split(rawData,"-")
			fmt.Println(rawData)
			fmt.Println(data)
			if len(data)>1 {
				resInfo.Title = append(resInfo.Title, data[0])
				resInfo.UserName = append(resInfo.UserName, data[1])
			} else {
				resInfo.Title = append(resInfo.Title, "Unknown")
				resInfo.UserName = append(resInfo.UserName, "Unknown")
			}
			resInfo.ID = append(resInfo.ID, id)
		})
		//ioutil.WriteFile("data/"+strconv.Itoa(order)+".txt", []byte(buf.String()), 0666)
	}
	// Write the res into res.json
	jsonRes, err := json.Marshal(resInfo)
	fmt.Println(resInfo)
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("config/res.json", jsonRes, 0666)
	c.JSON(http.StatusOK, resInfo)
}

func ParseJson(path string, info *LiveInfo) error {
	/* Read json file: */
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	var contents []byte
	contents, err = ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	/* Parse json file:*/
	err = json.Unmarshal(contents, info)
	if err != nil {
		return err
	}
	//fmt.Print(info)
	return err
}
