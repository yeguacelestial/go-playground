package HtmlParser

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func parseHtmlFromUrl(url string) string {
	fmt.Printf("Parsing HTML from %s ...\n", url)

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Reads HTML as slice of bytes
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return string(html)
}

func main() {
	url := "https://raw.githubusercontent.com/gophercises/link/master/ex1.html"
	htmlString := parseHtmlFromUrl(url)

	print("HTML PARSED: \n%s", htmlString)
}
