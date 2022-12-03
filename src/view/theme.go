package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var modal = tview.NewModal().
	SetText("Choose theme").
	AddButtons([]string{"Dark", "Light"})

func createThemePicker() *tview.Pages {

	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		switch buttonLabel {
		case "Light":
			setLightTheme()
			pages.SwitchToPage("Menu")

			break

		default:
			setDarkTheme()
			pages.SwitchToPage("Menu")
			break

		}
	})

	themePage := tview.NewPages().
		AddPage("modal", modal, true, true)

	return themePage
}

func setLightTheme() {
	tree.SetBackgroundColor(tcell.ColorGhostWhite)
	tree.SetBorderPadding(1, 1, 1, 1)
	tree.SetBorder(true)
	tree.SetBorderColor(tcell.ColorNavajoWhite)

	text.SetBackgroundColor(tcell.ColorNavajoWhite)

	articleText.SetBorderPadding(1, 2, 1, 2)
	articleText.SetBackgroundColor(tcell.ColorFloralWhite)

	modal.SetBackgroundColor(tcell.ColorNavajoWhite)

	return
}

func setDarkTheme() {
	tree.SetBackgroundColor(tcell.ColorBlack)
	tree.SetBorderPadding(1, 1, 1, 1)
	tree.SetBorder(true)
	tree.SetBorderColor(tcell.ColorNavajoWhite)

	articleText.SetBackgroundColor(tcell.ColorBlack)
	articleText.SetBorderPadding(1, 2, 1, 2)

	text.SetBackgroundColor(tcell.ColorBlack)

	modal.SetBackgroundColor(tcell.ColorNavajoWhite)

	return
}
