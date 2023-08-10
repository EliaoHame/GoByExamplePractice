package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	y := make([][]uint8, dy)
	for a := range y {
		x := make([]uint8, dx)
		for b := range x {
			x[b] = uint8(a * b)
		}
		y[a] = x
	}
	return y
}

func main() {
	pic.Show(Pic)
}
