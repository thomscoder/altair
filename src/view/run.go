package view

import (
	"fmt"

	"github.com/bytecodealliance/wasmtime-go/v3"
	"github.com/rivo/tview"
)

var playButton = tview.NewButton("Run")

func check(e error) error {
	return e
}

func run(path string, article *Article) {
	playButton.SetSelectedFunc(func() {
		engine := wasmtime.NewEngine()
		store := wasmtime.NewStore(engine)
		module, err := wasmtime.NewModuleFromFile(engine, path)
		check(err)
		instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})
		check(err)

		gcd := instance.GetExport(store, "gcd").Func()
		val, err := gcd.Call(store, 6, 27)
		check(err)
		s := fmt.Sprintf("gcd(6, 27) = %d\n", val.(int32))
		articleText.Clear()
		text := "article.Body %s"
		placeholder.SetText(s)
		articleText.SetText(fmt.Sprintf(text, placeholder.GetText(false))).ScrollToBeginning()
	})
}
