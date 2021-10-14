package util

import "time"

func GetMondayTimeStamp(now time.Time) int64 {
	offset := int64(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	nowTS := now.Unix()
	todayTS := nowTS - nowTS%86400

	return todayTS + offset*86400
}

func GetThis15MinTimeStamp(now time.Time) int64 {
	tomorrowTs := now.Unix()
	nextTs := tomorrowTs - tomorrowTs%86400
	return nextTs + int64(now.Hour()*3600+now.Minute()/15*900)
}
