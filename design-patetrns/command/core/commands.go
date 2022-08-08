package core

// instead of invoker calling something like tv.On(), it will call ONCommand.Execute()
// which will inturn call tv.On() and which acts liek a generic OnCommand() and will be reusable
type OnCommand struct {
	device IDevice
}

func NewOnCommand(d IDevice) ICommand {
	return &OnCommand{device: d}
}

func (c *OnCommand) Execute() {
	c.device.On()
}
