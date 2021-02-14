package main

import "fmt"

func main() {
	a := []string{"John", "Paul"}
	b := []string{"George", "Ringo", "Pete"}
	a = append(a, b...)  // equivalent to "append(a, b[0], b[1], b[2])"
	fmt.Println(a)
}