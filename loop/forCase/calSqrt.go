package main

import (
	"fmt"
	"math"
)

const inf = 0.000000000000001

func sqrt(x float64) float64 {
	z := 1.0
	tmp := 0.0
	for {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if math.Abs(z-tmp) < inf {
			break
		}
		tmp = z
	}
	return z
}

func main() {
	fmt.Println(sqrt(2))
}
