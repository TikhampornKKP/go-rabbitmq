package simple

import (
	"log"
	"time"
)

type TimeRange struct {
	Start string
	End   string
}

var (
	AcceptedTopic = "my-exchange"
	now           = time.Now()
	StartTime     = time.Date(now.Year(), now.Month(), now.Day(), 22, 0, 0, 0, now.Location())
	EndTime       = time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, now.Location())
)

func IsTimeClose(now time.Time) bool {
	if EndTime.Before(StartTime) {
		EndTime = EndTime.Add(24 * time.Hour)
	}

	if now.After(StartTime) && now.Before(EndTime) {
		log.Println("Current time is within the interval of 22:00 to 06:00")
		return true
	}

	log.Println("Current time is outside the interval of 22:00 to 06:00")
	return false
}
