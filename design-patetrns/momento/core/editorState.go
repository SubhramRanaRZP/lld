package core

type editorState struct {
	content string
}

func newEditorState(c string) *editorState {
	return &editorState{content: c}
}
