package enums

type CarState int16

const (
	CarStateUp CarState = iota + 1
	CarStateDown
	CarStateTemporaryStop
	CarStateStopped
)
