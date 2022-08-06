package core

type House struct {
	HouseType  string `json:"house_type"`
	Floor      int    `json:"floor"`
	DoorType   string `json:"door_type"`
	WindowType string `json:"window_type"`
}
