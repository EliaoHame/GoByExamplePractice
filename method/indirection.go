package main

import "fmt"

type InterVertex struct {
	X, Y float64
}

func (v InterVertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *InterVertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := InterVertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &InterVertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
