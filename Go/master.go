package main

import (
	"./jsonReader"
	"./deviceController"
	"./scheduler"
)

func main() {
	config := jsonReader.ReadJson("config.json")
	<-scheduler.ScheduleRecording(config.RecordInfo.StartTime, config.RecordInfo.StopTime, func() {
		deviceController.UnlockDevice(config.DeviceInfo.Ip, config.DeviceInfo.DevicePwd)
	})
}
