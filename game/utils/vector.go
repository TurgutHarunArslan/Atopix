package utils

import "math"

type Vector struct {
	X, Y float32
}

func (v1 *Vector) DistanceFrom(v2 Vector) float32 {
	dx := float64(v2.X - v1.X)
	dy := float64(v2.Y - v1.Y)
	return float32(math.Sqrt(dx*dx + dy*dy))
}