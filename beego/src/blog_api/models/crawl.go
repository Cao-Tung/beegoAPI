package models

import (

	"golang.org/x/net/html"
	"net/http"
	"io"
	"strings"

	"github.com/astaxie/beego/logs"
)



func Crawl(input string) []string{

	resp,err := http.Get(input)
	links := []string{}
	if err != nil{
		//fmt.Println(err)
		return links
	}
        logs.Info("crawl")
	links = crawlHref(resp.Body)
	return links
}


func crawlHref(b io.Reader) []string{
	links := []string{}
	z:=html.NewTokenizer(b)
	for ;; {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return links

		case tt == html.StartTagToken:
			t := z.Token()
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			ok, url := getHref(t)
			if !ok {
				continue
			}

			hasProto := strings.Index(url, "/") == 0

			if hasProto {
				links = append(links, url)
			}
		}
	}
	return links
}

func getHref(t html.Token) (ok bool, href string){
	for _,a := range t.Attr{
		if a.Key == "href"{
			href = a.Val
			ok = true
		}
	}
	return
}




