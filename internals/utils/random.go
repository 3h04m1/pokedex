package utils

import (
	"math/rand"
)

func RandomBool(n int) bool {
	catchRate := float64(n) / 100
	randomFactor := rand.Float64()
	return randomFactor <= catchRate
}
