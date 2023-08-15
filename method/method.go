package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v1 Vertex) Abs() float64 {
	return math.Sqrt(v1.X*v1.X + v1.Y*v1.Y)
}

func (v1 Vertex) Sum() float64 {
	return v1.X + v1.Y
}

func normalFunction(V Vertex) float64 {
	fmt.Println("Vertex: ", V)
	return 0.0
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(v.Sum())
	normalFunction(v)
}
