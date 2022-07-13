package main

import (
	"os"

	c "github.com/zachblizz/hot-mic/client"
)

const (
	MIC_NO_VIDEO = "yellow"
	MIC_AND_VIDEO = "red"
)

func main() {
	client := c.NewHueClient()

	if os.Args[1] == "vid" {
		client.TurnLightOn("Office", MIC_AND_VIDEO)
	} else if os.Args[1] == "mic" {
		client.TurnLightOn("Office", MIC_NO_VIDEO)
	} else {
		client.TurnLightOff("Office")
	}
}
