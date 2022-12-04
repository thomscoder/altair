package view

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Add nodes to tree
func add(target *tview.TreeNode, path string, articles []Article) {

	// Build the tree from the article path
	// checks if there's a article whose Path ends with the current path
	for _, article := range articles {
		if strings.HasSuffix(article.Path, fmt.Sprintf("%s/%s", path, article.Title)) {

			node := tview.NewTreeNode(article.Title).
				SetReference(filepath.Join(path, article.Title)).
				SetSelectable(true)

			if article.IsDir {
				node.SetColor(tcell.Color120)
			} else {
				node.SetColor(tcell.ColorLightBlue)
			}

			target.AddChild(node)
		}

	}
}

// Build the tree
func TreeMaker(root *tview.TreeNode, rootDir string, articles []Article) {
	// Add the current directory to the root node.
	add(root, rootDir, articles)

	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		children := node.GetChildren()
		path := reference.(string)
		if len(children) == 0 {
			for _, article := range Articles {
				// Read the path from the article Path
				if article.Path == path {
					if article.IsDir {
						add(node, path, articles)
					} else {
						setConcatText(&article)
						run(path, &article)
					}
				}
			}

		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())

		}
	})

}
