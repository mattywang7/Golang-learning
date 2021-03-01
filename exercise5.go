package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// SqrtImproved should return a non-nil error when given a negative number.
// Return a ErrNegativeSqrt value when given a negative number.
func SqrtImproved(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z * z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	value, err := SqrtImproved(-10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
}
