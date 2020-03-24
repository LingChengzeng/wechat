// time
package tools

import (
	"time"
)

/*
月份 1,01,Jan,January
日　 2,02,_2
时　 3,03,15,PM,pm,AM,am
分　 4,04
秒　 5,05
年　 06,2006
周几 Mon,Monday
时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
时区字母缩写 MST
*/
// This function returns the local Time corresponding to the Unix time
func Now() int64 {
	return time.Now().Unix()
}

// Format date string to unix time,like the function of PHP with strtotime
func Strtotime(date string) int64 {
	if date == "" {
		return 0
	}

	timer, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return 0
	}

	return timer.Unix()
}

// This function returns the unix time corresponding to date string.
// like 2006-01-02 15:04:05
// if the unixTime equals 0 then the unixTime equals local time.
func Timetostr(unixTime int64) string {
	if unixTime == 0 {
		unixTime = time.Now().Unix()
	}
	timer := time.Unix(unixTime, 0)
	return timer.Format("2006-01-02 15:04:05")
}

func GetDate() string {
	timer := time.Unix(time.Now().Unix(), 9)
	return timer.Format("2006-01-02")
}

// 2006-01-02 15:04:05 to 2006-01-02
func DateTimeToDate(dateTime string) string {
	timer, err := time.Parse("2006-01-02 15:04:05", dateTime)
	if err != nil {
		return ""
	}

	return timer.Format("2006-01-02")
}
