package helper

import (
	"fmt"
	"time"
)

func GetTimestamp() int64 {
	return time.Now().Unix()
}

func GetMsTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetTimeString() string {
	now := time.Now()
	return fmt.Sprintf("%s%d", now.Format("20060102150405"), now.UnixNano()%1e9)
}
