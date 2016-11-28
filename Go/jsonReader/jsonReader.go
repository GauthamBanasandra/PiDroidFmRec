package jsonReader

import (
	"os"
	"fmt"
	"encoding/json"
)

// Defining the JSON configuration structure
type Config struct {
	Action    []ActionElement `json:"action"` // The array of actions is executed sequentially.
	Ip        string `json:"ip"`              // IP address of the Android device.
	CmdDelay  int `json:"cmdDelay"`           // Delay (in seconds) between executing each action.
	DevicePwd string `json:"devicePwd"`       // Device password. Needed for unlocking the device.
}

type ActionElement struct {
	Cmd   string `json:"cmd"`      // Command to execute.
	Input InputType `json:"input"` // Input for the command to execute.
}

type InputType struct {
	X1   int `json:"x1"`      // For tap event x1 co-ordinate.
	Y1   int `json:"y1"`      // For tap event y1 co-ordinate.
	X2   int `json:"x2"`      // For tap event x2 co-ordinate.
	Y2   int `json:"y2"`      // For tap event y2 co-ordinate.
	Text string `json:"text"` // For text event input string.
}

// Function to read the config.json file
func ReadJson(filePath string) (config Config) {
	configFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening the config file", err.Error())
		os.Exit(1)
	}

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&config); err != nil {
		fmt.Println("error in parsing", err.Error())
		os.Exit(1)
	}

	configFile.Close()
	return
}
