package jsonReader

import (
	"os"
	"fmt"
	"encoding/json"
)

// Defining the JSON configuration structure
type Config struct {
	Action []ActionElement `json:"action"`	// The array of actions is executed sequentially.
	Ip     string `json:"ip"`	// IP address of the Android device.
}

type ActionElement struct {
	Cmd   string `json:"cmd"`	// Command to execute.
	Input InputType `json:"input"`	// Input for the command to execute.
}

type InputType struct {
	X int `json:"x"`	// For tap event x co-ordinate.
	Y int `json:"y"`	// For tap event y co-ordinate.
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
