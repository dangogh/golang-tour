package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ry := make([][]uint8, dy, dx)

	for y := 0; y < dy; y++ {
		ry[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			ry[y][x] = uint8(x * y)

		}

	}
	return ry
}

func main() {
	pic.Show(Pic)
}
