package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ",
		float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	oldz, z := x, x/2.0
	ii := 0
	for math.Abs(z-oldz) > 1e-15 {
		oldz, z = z, z-(math.Pow(z, 2)-x)/(2*z)
		ii++
	}
	fmt.Println("Took ", ii, " iterations")
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
