package main

import (
    "github.com/PuerkitoBio/goquery"
    "fmt"
)

func GetPage(url string) {
     doc, _ := goquery.NewDocument(url)
     doc.Find("img").Each(func(_ int, s *goquery.Selection) {
          url, _ := s.Attr("src")
          fmt.Println(url)
     })
}

func main() {
	url := "https://twitter.com/search?q=%23%E3%83%AD%E3%82%A2%E3%83%BC%E3%83%88&src=typeahead_click"
	GetPage(url)
	url = "https://twitter.com/search?q=from%3ARyu1__1uyR&src=typed_query&f=image"
    GetPage(url)
}
