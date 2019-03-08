package util

import (
	"time"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func StringToTime(dateStr string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, dateStr)
	return t, err
}


func UnixToTime(timeStamp int64) time.Time {
	return time.Unix(timeStamp, 0)
}

func DiffDay(date1, date2 time.Time) int {
	diff := date1.Sub(date2)
	return int(diff.Hours() / 24)
}

func DiffYear(date1, date2 time.Time) int {
	if date1.IsZero() || date2.IsZero() {
		return 0
	}
	diff := date1.Sub(date2)
	return int(diff.Hours() / 24 / 365.25)
}

func ExchangeNullTime (time time.Time) (nullTime mysql.NullTime) {
	return mysql.NullTime{
			Time: time,
			Valid: !time.IsZero(),
	}
}

func StartTime(date time.Time) (time.Time, error) {
	return StringToTime(fmt.Sprintf("%d-%02d-%02dT00:00:00Z", date.Year(), int(date.Month()), date.Day()))
}

func EndTime(date time.Time) (time.Time, error) {
	return StringToTime(fmt.Sprintf("%d-%02d-%02dT23:59:59Z", date.Year(), int(date.Month()), date.Day()))
}

func RequestStartTime(year, month, day int) (time.Time, error) {
	return StringToTime(fmt.Sprintf("%d-%02d-%02dT08:00:00Z", year, month, day))
}

func RequestEndTime(year, month, day int) (time.Time, error) {
	return StringToTime(fmt.Sprintf("%d-%02d-%02dT07:59:59Z", year, month, day))
}