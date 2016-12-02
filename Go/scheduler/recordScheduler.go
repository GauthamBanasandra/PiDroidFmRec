package scheduler

import (
	"../jsonReader"
	"time"
	"strings"
	"strconv"
	"fmt"
)

func GetSeconds(time jsonReader.Time) int {
	return time.Hh * 60 * 60 + time.Mm * 60 + time.Ss
}

func GetDiffSeconds(fromTime, toTime jsonReader.Time) int {
	fromSec := GetSeconds(fromTime)
	toSec := GetSeconds(toTime)
	diff := toSec - fromSec

	if diff < 0 {
		return diff + 24 * 60 * 60
	} else {
		return diff
	}
}

func ScheduleRecording(startTime, stopTime jsonReader.Time, callback func() ()) (endRecSignal chan string) {
	endRecSignal = make(chan string)
	go func() {
		timeSplice := strings.Split(time.Now().Format("15:04:05"), ":")
		hh, _ := strconv.Atoi(timeSplice[0])
		mm, _ := strconv.Atoi(timeSplice[1])
		ss, _ := strconv.Atoi(timeSplice[2])

		startAlarm := GetDiffSeconds(jsonReader.Time{hh, mm, ss}, startTime)
		fmt.Println("starting after ", startAlarm)
		time.AfterFunc(time.Duration(startAlarm) * time.Second, func() {
			callback()
			endRecSignal <- "finished recording"
			close(endRecSignal)
		})
	}()
	return
}
