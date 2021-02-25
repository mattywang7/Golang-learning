package main

import "fmt"

// Implement a fibonacci function that returns a function (a closure) that returns successive
// fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

func fibonacci() func() int {
	n1, n2 := -1, 1
	return func() int {
		result := n1 + n2
		n1 = n2
		n2 = result
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
