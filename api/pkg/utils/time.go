package utils

import (
	"time"
)

func GetCurrentTimeStamp() time.Time {
	location, _ := time.LoadLocation("Timezone")
	return time.Now().In(location)
}

func GetMonthScope(date time.Time) (string, string) {
	currentYear, currentMonth, _ := date.Date()
	currentLocation := date.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return firstOfMonth.Format(TimeStdFormat), lastOfMonth.Format(TimeStdDateFormat)
}

func GetCurrentMonthScope() (string, string) {
	now := time.Now()
	return GetMonthScope(now)
}

func GetLastMonthStartEnd() (int64, int64) {
	now := time.Now()
	lastMonthFirstDay := now.AddDate(0, -1, -now.Day()+1)
	lastMonthStart := time.Date(lastMonthFirstDay.Year(), lastMonthFirstDay.Month(), lastMonthFirstDay.Day(), 0, 0, 0, 0, now.Location()).Unix()
	lastMonthEndDay := lastMonthFirstDay.AddDate(0, 1, -1)
	lastMonthEnd := time.Date(lastMonthEndDay.Year(), lastMonthEndDay.Month(), lastMonthEndDay.Day(), 23, 59, 59, 0, now.Location()).Unix()
	return lastMonthStart, lastMonthEnd
}

func GetCurrentYearMonth() string {
	now := time.Now()
	return now.Format(TimeStdYMFormat)
}

func GetLastYearMonth() string {
	now := time.Now()
	lastMonth := now.AddDate(0, -1, -now.Day()+1)
	return lastMonth.Format(TimeStdYMFormat)
}

/**
获取本周周一的日期
*/
func GetCurrentWeekScope() (string, string) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := now.AddDate(0, 0, offset)
	weekEndDate := weekStartDate.AddDate(0, 0, 6)
	return weekStartDate.Format(TimeStdDateFormat), weekEndDate.Format(TimeStdDateFormat)
}

/**
获取上周的周一日期
*/
func GetLastWeekScope() (string, string) {
	thisWeekMonday, _ := GetCurrentWeekScope()
	TimeMonday, _ := time.Parse("2006-01-02", thisWeekMonday)
	lastWeekMonday := TimeMonday.AddDate(0, 0, -7)
	lastEndDate := lastWeekMonday.AddDate(0, 0, 6)

	return lastWeekMonday.Format(TimeStdDateFormat), lastEndDate.Format(TimeStdDateFormat)
}
