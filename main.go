package main

import (
	"fmt"

	c "github.com/zachblizz/hot-mic/client"
)

const (
	MIC_NO_VIDEO = "green"
	MIC_AND_VIDEO = "red"
)

func main() {
	fmt.Println("Hello, World!")

	client := c.NewHueClient()
	// client.TurnLightOn("Office", MIC_AND_VIDEO)
	client.TurnLightOff("Office")
}
