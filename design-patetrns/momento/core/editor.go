package core

type Editor struct {
	content      string
	stateHistory editorStateHistory
}

func NewEditor() *Editor {
	return &Editor{
		content:      "",
		stateHistory: editorStateHistory{},
	}
}

func (e *Editor) SetContent(content string) {
	state := newEditorState(content)
	e.stateHistory.addEditorState(state)
}

func (e *Editor) GetContent() string {
	return e.content
}

func (e *Editor) Undo() {
	e.content = e.stateHistory.restoreEditorState()
}

func (e *Editor) Redo() {
}
