package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func normalizeAddress(a string) string {
	if a[:4] != "http" {
		a = fmt.Sprint("https://" + a)
	}
	return a
}

func title(a string) string {

	resp, err := http.Get(a)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	title, ok := traverse(doc)

	if !ok {
		log.Fatalln("Shit happened and Could not retrieve title")
	}
	return title

}

func traverse(n *html.Node) (string, bool) {
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}
	return "", false

}
