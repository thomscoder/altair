package view

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var pagesArr = []*tview.Pages{}

type Article struct {
	Title  string
	Author string
	Body   string
	Path   string
}

var Articles = make([]Article, 0)
var Root = tview.NewTreeNode("Get Started 😂")

// Tview
var pages = tview.NewPages()
var articleText = tview.NewTextView()
var placeholder = tview.NewTextView()
var app = tview.NewApplication()
var sitePages = tview.NewList().ShowSecondaryText(false)
var tree = tview.NewTreeView().SetGraphics(false).SetTopLevel(0).SetPrefixes([]string{
	"[red]* ",
	"[white]- ",
	"[darkmagenta]- ",
}).SetRoot(Root).SetCurrentNode(Root)

var flex = tview.NewFlex()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorOrangeRed).
	SetText("(t) to change theme | (q) to quit")

// File content
func setConcatText(article *Article) {
	articleText.Clear()
	text := "Content: \n%s"
	placeholder.SetText(article.Body)
	placeholder.SetBackgroundColor(tcell.ColorNavajoWhite)

	articleText.SetText(fmt.Sprintf(text, placeholder.GetText(false))).ScrollToBeginning()
}

// Create the view
func Create() {
	sitePages.SetSelectedFunc(func(index int, name string, second_name string, shortcut rune) {
		setConcatText(&Articles[index])
	})

	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(tree, 0, 1, true).
			AddItem(articleText, 0, 4, false), 0, 6, false).
		AddItem(tview.NewFlex().
			AddItem(text, 0, 4, false).
			AddItem(playButton, 0, 1, false), 1, 1, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case rune('q'):
			app.Stop()
			return event
		case rune('t'):
			pages.SwitchToPage("Change theme")
			return event
		default:
			return event
		}
	})

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Change theme", createThemePicker(), true, false)

	setDarkTheme()

	if err := app.SetRoot(pages, true).SetFocus(tree).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
