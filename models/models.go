package models

// State - the state of the light
type State struct {
	On bool `json:"on"`
	Brightness int `json:"bri"`
	Hue int `json:"hue"`
	Sat int `json:"sat"`
	Alert string `json:"alert"`
	Mode string `json:"mode"`
	Reachable bool `json:"reachable"`
}

// Light - the light struct
type Light struct {
	State State
	ID string
	Type string `json:"type"`
	Name string `json:"name"`
}
