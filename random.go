package main

import (
	"math/rand"
	"time"
)

func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return 1 + rand.Intn(60)
}
