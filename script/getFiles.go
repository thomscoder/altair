package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// to include files at build time
// Url: https://go.dev/blog/generate

func main() {
	out, _ := os.Create("src/build/docs.go")
	out.Write([]byte("package build\n\nimport \"altair/src/view\"\n\n\nvar articles = []view.Article{\n"))
	getAltairFiles("docs", out)
	out.Write([]byte("}\n"))
}

func getAltairFiles(contentDir string, out *os.File) {

	files := map[string]fs.DirEntry{}

	filepath.WalkDir(contentDir, func(path string, file fs.DirEntry, e error) error {
		if file.Name() != ".git" {
			if e != nil {
				return e
			}
			files[path] = file
		}

		return nil
	})

	for path, file := range files {

		body, _ := os.ReadFile(path)
		stringified := fmt.Sprintf("{Title: `%s`, Author: `%s`, Body: `%s`, Path: `%s`,}, ", file.Name(), "Thomas", string(body), path)
		out.Write([]byte(stringified))
	}
}
