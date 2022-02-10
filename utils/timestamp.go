package utils

import (
	"strconv"
	"time"
)

func TimeStamp() string {
	return strconv.Itoa(int(time.Now().Unix()))
}
