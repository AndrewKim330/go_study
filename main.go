package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

// var baseURL string = "https://kr.indeed.com/jobs?q=python&l=&ts=1662298375496&rq=1&rsIdx=0&fromage=last&newcount=1541&vjk=1015284880e2ff62"
var baseURL = "https://kr.indeed.com/?from=gnav-jobsearch--indeedmobile"

//var baseURL string = "https://www.naver.com"

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

	fmt.Println(doc)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {

		//html, err := s.Html()
		//if err != nil {
		//	log.Fatalln(err)
		//}
		//fmt.Println(html)
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
