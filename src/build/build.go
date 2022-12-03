package build

import (
	"altair/src/view"
)

func BuildAltair() {

	for _, article := range articles {
		view.Articles = append(view.Articles, article)
	}

	view.TreeMaker(view.Root, "docs", articles)
	view.Create()
}
