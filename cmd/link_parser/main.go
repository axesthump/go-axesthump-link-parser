package main

import (
	"fmt"
	"go-axesthump-link-parser/internal/config"
	"golang.org/x/net/html"
	"os"
	"strings"
)

type link struct {
	Link string
	Body string
}

func main() {
	data, err := config.NewAppData()
	if err != nil {
		exit(err.Error())
	}

	var f func(*html.Node)
	links := make([]link, 0)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			hrefBody := ""
			newLink := link{
				Link: getHref(n.Attr),
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				hrefBody += parseHrefBody(c)
			}
			newLink.Body = hrefBody
			links = append(links, newLink)
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(data.Doc)

	fmt.Println(links)
}

func parseHrefBody(n *html.Node) string {
	if n.Type != html.ElementNode {
		return strings.TrimSpace(n.Data) + " "
	}
	data := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		data += parseHrefBody(c) + " "
	}
	return data[:len(data)-1]
}

func getHref(attr []html.Attribute) string {
	for _, el := range attr {
		if el.Key == "href" {
			return el.Val
		}
	}
	return ""
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
