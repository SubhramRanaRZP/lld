package core

import "fmt"

type Proxy struct {
	app *app
}

func NewProxy() *Proxy {
	return &Proxy{app: &app{}}
}

func (p *Proxy) Do() {
	fmt.Println("some checks and rate limiting and etc...")
	fmt.Println("proxy checks completed.")
	p.app.Do()
}
