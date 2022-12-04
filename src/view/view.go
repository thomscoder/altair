package view

import (
	"altair/src/graphics"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var Articles = make([]Article, 0)
var Root = tview.NewTreeNode(fmt.Sprintf("%s Get Started", graphics.Emoji["books"]))

// Tview
var pages = tview.NewPages()
var articleText = tview.NewTextView()
var placeholder = tview.NewTextView()
var app = tview.NewApplication()
var sitePages = tview.NewList().ShowSecondaryText(false)
var tree = tview.NewTreeView().SetGraphics(false).SetTopLevel(0).SetPrefixes([]string{
	"",
	fmt.Sprintf("%s ", graphics.Emoji["pointer"]),
	"[darkmagenta]- ",
}).SetRoot(Root).SetCurrentNode(Root)

var flex = tview.NewFlex()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorOrangeRed).
	SetText("(t) to change theme | (q) to quit")

// Create the view
func Create() {
	sitePages.SetSelectedFunc(func(index int, name string, second_name string, shortcut rune) {
		setConcatText(&Articles[index])
	})
	// Creates flexbox
	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(tree, 0, 1, true).
			AddItem(articleText, 0, 4, false), 0, 6, false).
		AddItem(tview.NewFlex().
			AddItem(text, 0, 4, false).
			AddItem(playButton, 0, 1, false), 1, 1, false)

	// Creates events
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

	// Set dark theme as default
	setDarkTheme()

	if err := app.SetRoot(pages, true).SetFocus(tree).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
