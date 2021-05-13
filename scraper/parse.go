package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Represents a link in an HTML document
type Link struct {
	Href string
	Text string
}

// Parse will take an HTML document and will return a slice of
// links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	msg := n.Data

	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}

	fmt.Println(padding, msg) // padding is for indenting children elements

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
