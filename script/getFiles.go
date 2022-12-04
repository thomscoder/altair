package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

			if file.IsDir() {
				stringified := fmt.Sprintf("{Title: `%s`, Author: `%s`, Body: `%s`, Path: `%s`, IsDir: %t}, ", file.Name(), "Thomas \u270C\uFE0F", string(body), path, file.IsDir())
				out.Write([]byte(stringified))
			}

			if filepath.Ext(path) == ".txt" {
				rxp := regexp.MustCompile("-([^-]+)\\S+")
				match := rxp.FindStringSubmatch(string(body))
				if len(match) > 0 {
					fileToGetContentFrom := strings.TrimSpace(match[len(match)-1])

					pathOfFileToGetContentFrom := filepath.Join(filepath.Dir(path), fileToGetContentFrom)
					newContent, _ := os.ReadFile(pathOfFileToGetContentFrom)
					newBody := rxp.ReplaceAllLiteral(body, newContent)

					stringified := fmt.Sprintf("{Title: `%s`, Author: `%s`, Body: `%s`, Path: `%s`, IsDir: %t, Snippet: `%s`}, ", file.Name(), "Thomas \u270C\uFE0F", string(newBody), path, file.IsDir(), string(newContent))
					out.Write([]byte(stringified))
				} else {
					stringified := fmt.Sprintf("{Title: `%s`, Author: `%s`, Body: `%s`, Path: `%s`, IsDir: %t,}, ", file.Name(), "Thomas \u270C\uFE0F", string(body), path, file.IsDir())
					out.Write([]byte(stringified))
				}
			}

		}

		return nil
	})

}
