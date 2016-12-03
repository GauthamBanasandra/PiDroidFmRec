package main

import (
	"./jsonReader"
	"./deviceController"
	"./scheduler"
)

func main() {
	config := jsonReader.ReadJson("config.json")

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

	start, stop := scheduler.ScheduleRecording(*startInfo, *stopInfo)
	<-start
	<-stop

}
