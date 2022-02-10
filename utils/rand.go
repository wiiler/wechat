package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func RandNumber() string {
	return fmt.Sprintf("%05v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000))
}

func OrderNo() string {
	return fmt.Sprintf("%s%s", time.Now().Format("20060102150405"), RandNumber())
}
