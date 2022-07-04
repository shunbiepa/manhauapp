package utils

import (
	"encoding/json"
	"errors"
	"time"
)

// GetZeroTimeOut 获取距离零点还有多少分钟
func GetZeroTimeOut() int {
	todayLast := time.Now().Format("2006-01-02") + " 23:59:59"
	todayLastTime, _ := time.ParseInLocation("2006-01-02 15:04:05", todayLast, time.Local)
	remainSecond := time.Duration(todayLastTime.Unix()-time.Now().Local().Unix()) * time.Second
	return int(remainSecond.Minutes())
}

func GeCreatedRange(createdRange string) (err error, startTime time.Time, endTime time.Time) {
	var s []string
	err = json.Unmarshal([]byte(createdRange), &s)
	if err != nil {
		return
	}
	if len(s) != 2 {
		err = errors.New("创建时间错误")
		return
	}

	startTime, err = time.Parse(time.RFC3339, s[0])
	endTime, err = time.Parse(time.RFC3339, s[1])

	return
}

func GetTimeofMonthStr() string {

	str := ""

	now := time.Now()
	str = str + now.Format("2006-01-02") + ","
	for i := 1; i < 30; i++ {
		if i == 29 {
			str = str + now.AddDate(0, 0, i).Format("2006-01-02")

		} else {
			str = str + now.AddDate(0, 0, i).Format("2006-01-02") + ","
		}
	}

	return str
}
