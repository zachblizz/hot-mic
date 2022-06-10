package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	c "github.com/zachblizz/hot-mic/config"
	"github.com/zachblizz/hot-mic/models"
)

const FULL_BRIGHTNESS = 254

type HueClient struct {
	client *http.Client
	lightColor map[string]int
	baseUrl string
}

// GetLights - gets all the lights
func (h *HueClient) getLights(url string) *http.Response {
	resp, err := h.client.Get(fmt.Sprintf("%s%s", h.baseUrl, url))

	if err != nil {
		panic(err)
	}

	return resp
}

// GetBytesFromResp - gets the bytes from the response
func (h *HueClient) getBytesFromResp(resp *http.Response) []byte {
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return b
}

func (h *HueClient) getLightByName(name string) models.Light {
	lights := make(map[string]models.Light)

	resp := h.getLights("/lights")
	lightsResp := h.getBytesFromResp(resp)
	json.Unmarshal(lightsResp, &lights)

	for k, v := range lights {
		if v.Name == name {
			v.ID = k
			return v
		}
	}

	return models.Light{}
}

// TurnLightOn - turns the light on
func (h *HueClient) TurnLightOn(lightName, color string) {
	light := h.getLightByName(lightName)
	h.changeLightState(light, true, h.lightColor[color])
}

// TurnLightOff - turns the light off
func (h *HueClient) TurnLightOff(lightName string) {
	light := h.getLightByName(lightName)
	h.changeLightState(light, false, 0)
}

func createState(on bool, hue int) models.State {
	return models.State{
		On: on,
		Brightness: FULL_BRIGHTNESS,
		Sat: 254,
		Hue: hue,
	}
}

func createJsonState(on bool, hue int) []byte {
	state := createState(on, hue)

	json, err := json.Marshal(state)
	if err != nil {
		panic(err)
	}

	return json
}

func (h *HueClient) changeLightState(light models.Light, on bool, hue int) {
	jsonState := createJsonState(on, hue)
	jsonBytes := bytes.NewBuffer(jsonState)

	url := fmt.Sprintf("%s/lights/%s/state", h.baseUrl, light.ID)

	req, err := http.NewRequest("PUT", url, jsonBytes)
	if err != nil {
		panic(err)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)
}

// NewHueClient - creates a new hue client
func NewHueClient() *HueClient {
	config := c.GetConfig()

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{ InsecureSkipVerify: true },
	}
	client := &http.Client{Transport: transCfg}

	return &HueClient{ baseUrl: config.Hue.BaseUrl, client: client, lightColor: config.Hue.Colors }
}
