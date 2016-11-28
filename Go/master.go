package main

import (
	"./jsonReader"
	"./deviceController"
)

func main() {
	config := jsonReader.ReadJson("config.json")
	deviceController.UnlockDevice(config.Ip, config.DevicePwd)
}
