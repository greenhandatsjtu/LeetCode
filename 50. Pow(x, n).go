package main

import (
	"fmt"
)

func quickMul(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n > 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func main() {
	res := myPow(0.55356, 0)
	fmt.Print(res)
}
