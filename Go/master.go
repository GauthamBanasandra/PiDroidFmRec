package main

import (
	"./jsonReader"
	"./deviceController"
)

func main() {
	config := jsonReader.ReadJson("config.json")
	deviceController.UnlockDevice(config.Ip, config.DevicePwd)
	/*for _, cmd := range deviceController.GetCommands(config.Action) {
		fmt.Println(cmd)
	}*/
}
