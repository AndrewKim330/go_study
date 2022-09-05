package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

// var baseURL string = "https://kr.indeed.com/jobs?q=python&l=&ts=1662298375496&rq=1&rsIdx=0&fromage=last&newcount=1541&vjk=1015284880e2ff62"
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

	doc.find(".pagination").Each()

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
