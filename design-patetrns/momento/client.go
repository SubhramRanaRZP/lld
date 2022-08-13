package main

import (
	"lld/design-patetrns/momento/core"
)

func main() {
	editor := core.NewEditor()

	editor.SetContent("a")
	editor.GetContent()

	editor.SetContent("b")
	editor.SetContent("c")
	editor.GetContent()

	editor.Undo()
	editor.GetContent()

	editor.Redo()
	editor.GetContent()

	editor.Redo() // now nothing to redo, so the previous content will be displayed
	editor.GetContent()
}
