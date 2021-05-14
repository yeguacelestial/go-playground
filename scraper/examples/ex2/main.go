package main

import (
	"fmt"
	"strings"

	HtmlParser "github.com/yeguacelestial/go-playground/parse-html-from-url"
	link "github.com/yeguacelestial/go-playground/scraper"
)

var exampleHtml = HtmlParser.ParseHtmlFromUrl("https://raw.githubusercontent.com/gophercises/link/master/ex1.html")

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
