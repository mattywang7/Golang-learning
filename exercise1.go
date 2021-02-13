package main

import "fmt"

// Given a number x, we want to find the number z for which z^2 is most nearly x.

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		fmt.Println(z)
		z -= (z * z - x) / (2 * z)
	}
	fmt.Println("-------------")
	return z
}

func main() {
	//fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(14))
}
