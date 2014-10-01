package gowiki

import "github.com/PuerkitoBio/goquery"

const randomWikipediaArticle = "http://en.wikipedia.org/wiki/Special:Random"

type WikiNode struct {
	Url      string
	ImageUrl string
}

func Random() (WikiNode, error) {
	doc, err := goquery.NewDocument(randomWikipediaArticle)

	if err != nil {
		return WikiNode{}, err
	}

	selection := doc.Find("#mw-content-text .image img").First()

	src, exists := selection.Attr("src")

	node := WikiNode{}
	if exists {
		node.ImageUrl = src
	}

	return node, nil
}
