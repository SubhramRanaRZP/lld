package main

import (
	"lld/design-patetrns/proxy/core"
)

func main() {
	server := core.NewProxy()
	server.Do()
}
