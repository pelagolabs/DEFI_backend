package util

import "time"

const (
	TimeLayer     = "2006-01-02 15:04:05"
	TimeHourLayer = "2006-01-02 15"
	DateLayer     = "2006-01-02"
	MonthLayer    = "2006-01"
)

func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return getZeroTime(d)
}

func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

func GetTimeStrFixedByTZ(oriTimeStr string, tzDiff int) string {
	if tzDiff == 0 {
		return oriTimeStr
	}

	oriTime, _ := time.Parse(TimeLayer, oriTimeStr)
	return oriTime.Add(time.Duration(tzDiff*-1) * time.Hour).Format(TimeLayer)
}

func getZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
