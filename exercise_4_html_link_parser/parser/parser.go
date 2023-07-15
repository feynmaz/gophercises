package parser

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseLinks(input string) []Link {
	links := make([]Link, 0)

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			if n.PrevSibling != nil && n.PrevSibling.Data == "" {
				return
			}

			link := Link{}
			for _, a := range n.Attr {
				if a.Key == "href" {
					link.Href = a.Val
					break
				}
			}

			text := ""
			var getAllText func(n *html.Node)
			getAllText = func(n *html.Node) {
				if n.Type == html.TextNode {
					text += n.Data
				}
				if n.FirstChild != nil {
					getAllText(n.FirstChild)
				}
				if n.NextSibling != nil {
					getAllText(n.NextSibling)
				}
			}
			if n.FirstChild == nil {
				text += n.Data
			} else {
				getAllText(n)
			}

			link.Text = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(text, "\n", " "), "\t", ""))
			links = append(links, link)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return links
}
