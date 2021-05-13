/* STEPS
1. Find <a> nodes in the document
2. For each link node...
	2.a build a Link
3. Return the links
*/

package link

import (
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

	nodes := linkNodes(doc)

	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link

	// Iterate through all the attributes of the node
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}

	ret.Text = "TODO: Parse the text..."
	return ret
}

func linkNodes(n *html.Node) []*html.Node {
	// Verifiy that the node is a valid anchor tag
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	// If the node is not anchor tag, still verifiy the child elements
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
