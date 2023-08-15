package main

import (
	"fmt"
	"math"
)

type NewVertex struct {
	X, Y float64
}

func Abs(v NewVertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v NewVertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := NewVertex{3, 4}
	Scale(v, 10)
	fmt.Println(Abs(v))
}
