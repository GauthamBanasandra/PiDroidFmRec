package main

import (
	"./jsonReader"
	"./deviceController"
	"./scheduler"
)

// Driver program.
func main() {
	// Read the configuration file.
	config := jsonReader.ReadJson("config.json")

	// Get the start and stop actions.
	startActions, stopActions := config.GetActionList()

	var startInfo = &scheduler.ScheduleAction{
		Time:config.RecordInfo.StartTime,
		Action:func() {
			deviceController.UnlockDevice(config.DeviceInfo.Ip, config.DeviceInfo.DevicePwd)
			deviceController.ExecuteCommands(deviceController.GetCommands(startActions))
		},
	}

	var stopInfo = &scheduler.ScheduleAction{
		Time:config.RecordInfo.StopTime,
		Action:func() {
			deviceController.UnlockDevice(config.DeviceInfo.Ip, config.DeviceInfo.DevicePwd)
			deviceController.ExecuteCommands(deviceController.GetCommands(stopActions))
		},
	}

	// Schedule the start and stop actions.
	start, stop := scheduler.ScheduleRecording(*startInfo, *stopInfo)

	// Wait till both the alarms and their handlers finish.
	<-start
	<-stop

}
