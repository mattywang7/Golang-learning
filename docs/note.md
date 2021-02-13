# Golang syntax

## For

Go has only one looping construct, the `for` loop.

No parentheses and the braces are always required.

The init and post statements are optional. C's `while` is spelled `for` in Go.

```go
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```

## If

Like `for`, the `if` statement can start with a short statement to execute before the condition.
```go
package main
import "math"
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
```


