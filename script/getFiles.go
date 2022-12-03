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

	filepath.WalkDir(contentDir, func(path string, file fs.DirEntry, e error) error {
		if file.Name() != ".git" {
			if e != nil {
				return e
			}
			body, _ := os.ReadFile(path)
			stringified := fmt.Sprintf("{Title: `%s`, Author: `%s`, Body: `%s`, Path: `%s`, IsDir: %t}, ", file.Name(), "Thomas", string(body), path, file.IsDir())
			out.Write([]byte(stringified))
		}

		return nil
	})

}