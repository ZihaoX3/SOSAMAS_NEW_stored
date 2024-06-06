package utils

import (
	"math/rand"
)

func GenerateRandomCoordinates() Coordinates {
	return Coordinates{X: rand.Float64() * GridWidth, Y: rand.Float64() * GridHeight}
}

func GenerateRandomFloat(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func Contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func ContainsString(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
