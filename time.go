package common

import (
	"errors"
	"fmt"
	"time"
)

var months = [...]string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

// 返回字符串格式的本地时间
func Now() (string, error) {
	now := time.Now().Local()
	year, mon, day := now.Date()
	hour, min, sec := now.Clock()
	for i, v := range months {
		if mon.String() == v {
			return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n", year, i+1, day, hour, min, sec), nil
		}
	}
	return "", errors.New("month convert error")
}

//返回包含时区的本地时间
func NowIncludeZone() time.Time {
	time.LoadLocation("CST") //设置Local时区为UT
	return time.Now()
}
