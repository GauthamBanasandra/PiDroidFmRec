package jsonReader

import (
	"os"
	"fmt"
	"encoding/json"
)

// Defining the JSON configuration structure
type Config struct {
	Action     []ActionElement `json:"action"` // The array of actions is executed sequentially.
	DeviceInfo DeviceInfo `json:"deviceInfo"`  // Contains device specific information.
	RecordInfo RecordInfo `json:"recordInfo"`  // Contains configuration information regarding recording.
}

type DeviceInfo struct {
	Ip        string `json:"ip"`        // IP address of the Android device.
	DevicePwd string `json:"devicePwd"` // Device password. Needed for unlocking the device.
}

type ActionElement struct {
	Cmd   string `json:"cmd"`      // Command to execute.
	Input InputType `json:"input"` // Input for the command to execute.
}

type InputType struct {
	X1          int `json:"x1"`             // For tap event x1 co-ordinate.
	Y1          int `json:"y1"`             // For tap event y1 co-ordinate.
	X2          int `json:"x2"`             // For tap event x2 co-ordinate.
	Y2          int `json:"y2"`             // For tap event y2 co-ordinate.
	Text        string `json:"text"`        // For text event input string.
	PackageName string `json:"packageName"` // Package name of the app to be launched.
	Key         int `json:"key"`            // Keyevents for the phone - simulating press of a button.
}

type RecordInfo struct {
	StartTime      Time `json:"startTime"`       // Start recording FM at this time.
	StopTime       Time `json:"stopTime"`        // Stop recording FM at this time.
	StartActionIdx []int `json:"startActionIdx"` // Indices of the actions to be executed on start.
	StopActionIdx  []int `json:"stopActionIdx"`  // Indices of the actions to be executed on stop.
}

type Time struct {
	Hh int `json:"hh"` // Hours.
	Mm int `json:"mm"` // Minutes.
	Ss int `json:"ss"` // Seconds.
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

func getList(actionElements []ActionElement, indices []int) (subList []ActionElement) {
	for _, i := range indices {
		subList = append(subList, actionElements[i])
	}

	return
}

func (config Config) GetActionList() (startActionList, stopActionList []ActionElement) {
	startActionList = getList(config.Action, config.RecordInfo.StartActionIdx)
	stopActionList = getList(config.Action, config.RecordInfo.StopActionIdx)

	return
}