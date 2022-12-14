package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

//var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50&start=9999"

var baseURL string = "https://www.naver.com"

func main() {
	getPages()
}

func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".search_area").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with status: ", res.StatusCode)
	}
}
