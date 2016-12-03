package main

import (
	"./jsonReader"
	"./deviceController"
	"./scheduler"
	"fmt"
)

func main() {
	config := jsonReader.ReadJson("config.json")

	actions := deviceController.GetCommands(config.Action)

	var startInfo = &scheduler.ScheduleAction{
		Time:config.RecordInfo.StartTime,
		Action:func() {
			deviceController.UnlockDevice(config.DeviceInfo.Ip, config.DeviceInfo.DevicePwd)
			deviceController.ExecuteCommands(actions[:2])
		},
	}

	var stopInfo = &scheduler.ScheduleAction{
		Time:config.RecordInfo.StopTime,
		Action:func() {
			fmt.Println("Recording stopped")
		},
	}

	start, stop := scheduler.ScheduleRecording(*startInfo, *stopInfo)
	<-start
	<-stop
}
