package helper

import "time"

func CurrentTimestamp() int64 {
	return time.Now().UTC().Unix()
}
