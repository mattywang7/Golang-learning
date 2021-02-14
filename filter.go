package main

// Filter returns a new slice holding only the elements of s that satisfy fn()
func Filter(s []int, fn func(int) bool) []int {
	var p []int  // p == nil, len(p) == 0
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}
