package main

import (
	"math/rand"
	"time"
)

func scaleFloat(value float64, min float64, max float64) float64 {
	rand.Seed(time.Now().UnixNano())

	if value < min || value > max {
		return min + rand.Float64()*(max-min)
	}
	return 0
}
