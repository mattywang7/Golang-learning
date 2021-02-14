package main

import "fmt"

func main() {
	s := []byte{'r', 'o', 'a', 'd'}
	t := make([]byte, len(s), (cap(s) + 1) * 2)
	for i := range s {
		t[i] = s[i]
	}
	s = t
	fmt.Println(cap(s))
	fmt.Println(len(s))
}
