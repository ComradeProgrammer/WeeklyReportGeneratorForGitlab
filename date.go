package main

import (
	"time"
)

func GetLastWeekISO8601() (string, string) {
	currentTime := time.Now()
	today := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local)

	minus := -24 * int(today.Weekday())
	var lastDayOfLastWeek = today.Add(time.Duration(int(time.Hour) * minus))
	var firstDayOfLastWeek = lastDayOfLastWeek.Add(time.Hour * -168)
	lastDayOfLastWeek=lastDayOfLastWeek.Add(time.Second * (24*3600 - 1))

	return firstDayOfLastWeek.Format("2006-01-02T15:04:05Z"), lastDayOfLastWeek.Format("2006-01-02T15:04:05Z")

}
