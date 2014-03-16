package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := x / 2
	z = z - ((cmplx.Pow(z, complex(3, 0)) - x) / (3 * cmplx.Pow(z, complex(2, 0))))
	return z
}

func main() {
	fmt.Println(Cbrt(2))
}
