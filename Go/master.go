package main

import (
	"./jsonReader"
	"./scheduler"
	"fmt"
)

func main() {
	config := jsonReader.ReadJson("config.json")
	//deviceController.UnlockDevice(config.DeviceInfo.Ip, config.DeviceInfo.DevicePwd)
	<-scheduler.ScheduleRecording(config.RecordInfo.StartTime, config.RecordInfo.StopTime, func() {
		fmt.Println("alarm received")
	})
}
