package view

import (
	"fmt"
)

// File content
func setConcatText(article *Article) {
	articleText.Clear()
	text := article.Body

	articleText.SetTitle(fmt.Sprintf("%s - author: %s", article.Title, article.Author))
	articleText.SetText(text).ScrollToBeginning()
}

// func renderBody(article *Article) {
// 	textView := tview.NewTextView().
// 		SetDynamicColors(true).
// 		SetRegions(true).
// 		SetChangedFunc(func() {
// 			app.Draw()
// 		})
// 	numSelections := 0
// 	go func() {
// 		for _, word := range strings.Split(article.Body, " ") {
// 			if word == "the" {
// 				word = "[red]the[white]"
// 			}
// 			if word == "to" {
// 				word = fmt.Sprintf(`["%d"]to[""]`, numSelections)
// 				numSelections++
// 			}
// 			fmt.Fprintf(textView, "%s ", word)
// 			time.Sleep(200 * time.Millisecond)
// 		}
// 	}()
// 	textView.SetDoneFunc(func(key tcell.Key) {
// 		currentSelection := textView.GetHighlights()
// 		if key == tcell.KeyEnter {
// 			if len(currentSelection) > 0 {
// 				textView.Highlight()
// 			} else {
// 				textView.Highlight("0").ScrollToHighlight()
// 			}
// 		} else if len(currentSelection) > 0 {
// 			index, _ := strconv.Atoi(currentSelection[0])
// 			if key == tcell.KeyTab {
// 				index = (index + 1) % numSelections
// 			} else if key == tcell.KeyBacktab {
// 				index = (index - 1 + numSelections) % numSelections
// 			} else {
// 				return
// 			}
// 			textView.Highlight(strconv.Itoa(index)).ScrollToHighlight()
// 		}
// 	})
// }
