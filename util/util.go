package util

import (
	"math/rand"
	"time"
)

func RollDSix() int {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(6 - 1 + 1) + 1
}