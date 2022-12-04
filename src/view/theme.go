package view

import (
	"altair/src/graphics"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var lightThemeButton = tview.NewButton(graphics.Emoji["sun"]).
	SetSelectedFunc(func() {
		setLightTheme()
	}).
	SetBackgroundColorActivated(tcell.ColorBlack)

var darkThemeButton = tview.NewButton(graphics.Emoji["moon"]).
	SetSelectedFunc(func() {
		setDarkTheme()
	}).
	SetBackgroundColorActivated(tcell.ColorBlack)

func setLightTheme() {
	tree.SetBackgroundColor(tcell.ColorGhostWhite)
	tree.SetBorderPadding(1, 1, 1, 1)
	tree.SetBorder(true)
	tree.SetBorderColor(tcell.ColorNavajoWhite)

	text.SetBackgroundColor(tcell.ColorNavajoWhite)
	text.SetTextColor(tcell.ColorBlack)

	articleText.SetBorderPadding(1, 2, 1, 2)
	articleText.SetBackgroundColor(tcell.ColorPapayaWhip)
	articleText.SetTextColor(tcell.ColorBlack)
	darkThemeButton.SetBackgroundColor(tcell.Color101)
}

func setDarkTheme() {
	tree.SetBackgroundColor(tcell.ColorBlack)
	tree.SetBorderPadding(1, 1, 1, 1)
	tree.SetBorder(true)

	articleText.SetBackgroundColor(tcell.ColorBlack)
	articleText.SetBorderPadding(2, 2, 10, 10)
	articleText.SetBorder(true)

	text.SetBackgroundColor(tcell.ColorBlack)
}
