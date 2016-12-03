package scheduler

import (
	"../jsonReader"
	"time"
	"strings"
	"strconv"
	"fmt"
)

// Structure representing the action to be scheduled.
type ScheduleAction struct {
	Time   jsonReader.Time
	Action func() ()
}

// Converts time in hh-mm-ss format to seconds.
func GetSeconds(time jsonReader.Time) int {
	return time.Hh * 60 * 60 + time.Mm * 60 + time.Ss
}

/*
Returns the difference between the given 2 times(A, B).
If A-B < 0, then it adds 24*60*60 (total seconds in a day) and returns.
*/
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

// Returns the current time in hh-mm-ss format.
func GetTimeNow() (hh, mm, ss int) {
	timeSplice := strings.Split(time.Now().Format("15:04:05"), ":")
	hh, _ = strconv.Atoi(timeSplice[0])
	mm, _ = strconv.Atoi(timeSplice[1])
	ss, _ = strconv.Atoi(timeSplice[2])

	return
}

/*
Schedules an alarm with the corresponding action to be executed.
Also, it returns a channel to signal its triggering.
*/
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

// Schedules start and stop of recording.
func ScheduleRecording(start, stop ScheduleAction) (endStartSignal, endStopSignal chan string) {
	endStartSignal = ScheduleAlarm(start)
	endStopSignal = ScheduleAlarm(stop)

	return
}