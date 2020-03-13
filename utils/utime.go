package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func CurrentDate() string {
	return time.Now().Format("20060102")
}

func CurrentDateByPlace(place string) string {
	return time.Now().Format(fmt.Sprintf("2006%s01%s02", place, place))
}

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func FormatMonth(t time.Time) string {
	return t.Format("2006-01")
}

func FormatMonthChinese(t time.Time) string {
	return t.Format("2006年01月")
}

func ParseTime(date string) (time.Time, error) {
	local, _ := time.LoadLocation("Local")
	return time.ParseInLocation("2006-1-2 15:04:05", replaceTimeClear(date), local)

}

func ParseDate(date string) (time.Time, error) {
	local, _ := time.LoadLocation("Local")
	return time.ParseInLocation("2006-1-2", replaceTimeClear(date), local)
}

func replaceTimeClear(date string) string {
	date = strings.Replace(date, "/", "-", -1)
	date = strings.Replace(date, ".", "-", -1)
	date = strings.Replace(date, "-0", "-", -1)
	return date
}

//转换excel中日期字符串（有可能为数字格式的字符串）
func ParseExcelDate(date string) (d *time.Time, err error) {
	if date != "" {
		var date2 time.Time
		if !IsValidNumber(date) {
			date2, err = ParseDate(date)
			if err != nil {
				return
			}
			d = &date2
			return
		} else {
			date2, err = ParseDate("1900-1-1")
			if err != nil {
				return
			}
			var days int
			days, err = strconv.Atoi(date)
			if err != nil {
				return
			}
			date2 = date2.AddDate(0, 0, days-2)
			d = &date2
			return
		}
	}
	return
}
