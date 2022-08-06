package core

import "fmt"

func NewGUI(gui string) GUI {
	switch gui {
	case "mac":
		return newMacGUI()
	case "win":
		return newWindowGUI()
	}

	return nil
}

// ----------------
type button struct {
	shape string
}

func newButton(shape string) *button {
	return &button{shape: shape}
}

func (b *button) click() {
	fmt.Println("button clicked with shape ", b.shape)
}

// -------------------
type mac struct {
	button *button
}

func newMacGUI() *mac {
	m := &mac{}
	m.button = newButton("oval")
	return m
}

func (m *mac) ClickButton() {
	m.button.click()
}

// -------------------
type win struct {
	button *button
}

func newWindowGUI() *win {
	w := &win{}
	w.button = newButton("rectangle")
	return w
}

func (w *win) ClickButton() {
	w.button.click()
}

// -------------------
