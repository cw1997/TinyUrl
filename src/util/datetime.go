// datetime.go
// author:昌维 [github.com/cw1997]
// date:2017-05-11 01:07:16
package util

import (
	"time"
)

func GetDatetime() string {
	//获取时间戳
	timestamp := time.Now().Unix()
	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05 PM")
}

func GetTimestamp() int64 {
	//获取时间戳
	return time.Now().Unix()
}
