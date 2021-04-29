package utils

import (
	"fmt"
	"time"
)

func GetDateTime() string {
	timeObj := time.Now()
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second) //2020-04-26 17:36:04
}

func GetDate() string {
	timeObj := time.Now()
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
