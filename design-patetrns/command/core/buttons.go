package core

type TvPowerButton struct {
	onCommand ICommand
}

func NewTvPowerButton(cmd ICommand) IButton {
	return &TvPowerButton{onCommand: cmd}
}

func (b *TvPowerButton) Click() {
	b.onCommand.Execute()
}

type RemotePowerButton struct {
	onCommand ICommand
}

func NewRemotePowerButton(cmd ICommand) IButton {
	return &RemotePowerButton{onCommand: cmd}
}

func (b *RemotePowerButton) Click() {
	b.onCommand.Execute()
}
