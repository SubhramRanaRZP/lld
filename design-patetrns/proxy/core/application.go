package core

import "fmt"

type app struct{}

func (app *app) Do() {
	fmt.Println("app.done")
}
