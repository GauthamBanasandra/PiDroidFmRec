package scheduler

import (
	"../jsonReader"
	"time"
	"strings"
	"strconv"
	"fmt"
)

type ScheduleAction struct {
	Time   jsonReader.Time
	Action func() ()
}

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

func GetTimeNow() (hh, mm, ss int) {
	timeSplice := strings.Split(time.Now().Format("15:04:05"), ":")
	hh, _ = strconv.Atoi(timeSplice[0])
	mm, _ = strconv.Atoi(timeSplice[1])
	ss, _ = strconv.Atoi(timeSplice[2])

	return
}

func ScheduleAlarm(scheduler ScheduleAction) (endSignal chan string) {
	endSignal = make(chan string)
	go func() {
		hh, mm, ss := GetTimeNow()
		alarmTime := GetDiffSeconds(jsonReader.Time{hh, mm, ss}, scheduler.Time)

		// Debug.
		fmt.Println("starting recording after ", alarmTime, "s")

		time.AfterFunc(time.Duration(alarmTime) * time.Second, func() {
			scheduler.Action()
			endSignal <- "finished"
			close(endSignal)
		})
	}()

	return
}

func ScheduleRecording(start, stop ScheduleAction) (endStartSignal, endStopSignal chan string) {
	endStartSignal = ScheduleAlarm(start)
	endStopSignal = ScheduleAlarm(stop)

	return
}