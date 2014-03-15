package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	oldz, z := x, x/2.0
	ii := 0
	for math.Abs(z-oldz) > 1e-15 {
		oldz, z = z, z-(math.Pow(z, 2)-x)/(2*z)
		ii++
	}
	fmt.Println("Took ", ii, " iterations")
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
