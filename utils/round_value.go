package utils

import "math"

func RoundValue(value float64) float64 {
	return math.Round(100*value) / 100
}
