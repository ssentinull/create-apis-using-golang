package utils

import (
	"math/rand"
	"time"
)

func GenerateID() int64 {
	return time.Now().UnixNano() + int64(rand.Intn(10000))
}
