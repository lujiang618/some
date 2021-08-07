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

func GetLastMonthScope() (time.Time, time.Time) {
	now := time.Now()
	lastMonthFirstDay := now.AddDate(0, -1, -now.Day()+1)
	lastMonthEndDay := lastMonthFirstDay.AddDate(0, 1, -1)
	return lastMonthFirstDay, lastMonthEndDay
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

func GetCurrentWeekDay() int {
	return GetWeekDay(time.Now())
}

func GetWeekDay(date time.Time) int {
	day := int(date.Weekday())
	if day == 0 {
		day = 7
	}
	return day
}

func GetCurrentDayOfYear() int {
	return GetDayOfYear(time.Now())
}

func GetLastMonthDayOfYear() int {
	_, dateEnd := GetLastMonthScope()

	return GetDayOfYear(dateEnd)

}
func GetDayOfYear(date time.Time) int {
	return date.YearDay()
}

func GetCurrentYear() int {
	return time.Now().Year()
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
