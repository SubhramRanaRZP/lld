package main

import (
	"lld/design-patetrns/command/core"
)

func main() {
	tv := core.NewTv()
	onCommand := core.NewOnCommand(tv)

	tvPowerButton := core.NewTvPowerButton(onCommand)
	remotePowerButton := core.NewTvPowerButton(onCommand)

	tvPowerButton.Click()
	remotePowerButton.Click()
}
