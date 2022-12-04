package view

import (
	"fmt"
	"strings"
)

var quit = make(chan bool)

// File content
func setConcatText(article *Article) {
	articleText.Clear()
	articleText.SetTitle(fmt.Sprintf("%s - author: %s", article.Title, article.Author))
	renderBody(article)
	articleText.ScrollToBeginning()
}

func renderBody(article *Article) {

	for _, word := range strings.Split(article.Body, " ") {
		fmt.Fprintf(articleText, "%s ", word)
	}
}
