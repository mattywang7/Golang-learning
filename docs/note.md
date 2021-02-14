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

## Switch

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow.
In effect, the `break` statement that is needed at the end of each case in those languages is provided automatically in Go.

Go's switch cases need not be constants, and the values involved need not be integers.

Switch without a condition is the same as `switch true`. This construct can be a clean way to write long if-then-else chains.

## Defer

A `defer` statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in LIFO order.

## Pointers

The `&` operator generates a pointer to its operand.

The `*` operator denotes the pointer's underlying value.

Go has no pointer arithmetic.

## Structs

A `struct` is a collection of fields.
Struct fields are accessed using a dot.

Struct fields can be accessed through a struct pointer.

To access the field `X` of a struct when we have the struct pointer `p` we could write `(*p).X`. However, that notation is cumbersome, so the language permits us instead to write just `p.X`, without the explicit dereference.

A struct literal denotes a newly allocated struct value by listing the values of its fields.
```
var (
	v1 = Vertex{1, 2}	// has type Vertex
	v2 = Vertex{X: 1}	// Y: 0 is implicit
	v3 = Vertex{}		// X: 0 and Y: 0
	p = &Vertex{1, 2}	// has type *Vertex
)
```

## Arrays

The type `[n]T` is an array of `n` values of type `T`.

An array's length is part of its type, so arrays cannot be resized.

Go's arrays are values. An array variable denotes the entire array; it is not a pointer to the first array element (as would be the case in C).
This means that when you assign or pass around an array value you will make a copy of its contents.
One way to think about `array` is as a sort of struct but with indexed rather than named fields: a fixed-size composite value.

An array literal can be specified like so:
```
b := [2]string{"Penn", "Teller"}
```
or let the compiler to count the array elements
```
b := [...]string{"Penn", "Teller"}
```

```
var a [2]string
a[0] = "Hello"
a[1] = "World"
primes := [6]int{2, 3, 5, 7, 11, 13}
```

## Slices

An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
In practice, slices are much more common than arrays.

The type `[]T` is a lice with elements of type `T`.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:
```
a[low : high]
```

A slice does not store any data, it just describes a section of an underlying array.
Changing the elements of a slice modifies the corresponding elements of its underlying array.
Other slices that share the same underlying array will see those changes.

When slicing, you may omit the high or low bounds to use their defaults instead. The default is zero for the low bound and the length of the slice for the high bound.

A slice has both a **length** and a **capacity** (`len(s)` and `cap(s)`).
The length of a slice is the number of elements it contains. The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

Slices can be created with the built-in `make` function; this is how you create dynamically-sized arrays.
The `make` function allocates a zeroed array and returns a slice that refers to that array:
```
a := make([]int, 5)     // len(a) == 5
```
To specify a capacity, pass a third argument to `make`.
```
func make([]T, len, cap) []T
```
When called, `make` allocates an array and returns a slice that refers to that array.

### Slice internals

A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).

![slice-structure](./images/slice-structure.jpeg)

The length is the number of elements referred to by the slice. The capacity is the number of elements in the underlying array (beginning at the element referred to by the slice pointer).
```
s = s[2:4]
```

![length-capacity](./images/length-capacity.jpeg)

Slicing does not copy the slice's data. It creates a new slice value that points to the original array.
This makes slice operations as efficient as manipulating array indices. 
Therefore, modifying the elements (not the slice itself) of a re-slice modifies the elements of the original slice.

A slice cannot be grown beyond its capacity. Attempting to do so will cause a runtime panic, just as when indexing outside the bounds of a lice or array.

To increase the capacity of a slice one must create a new, larger slice and copy the contents of the original slice into it.

Go provides a built-in `append` function that's good for most purposes.
```
func append(s []T, x ...T) []T
```
The `append` function appends the elements `x` to the end of the slice `s`, and grows the slice if a greater capacity is needed.
```
a := []string{"John", "Paul"}
b := []string{"George", "Ringo", "Pete"}
a = append(a, b...)  // equivalent to "append(a, b[0], b[1], b[2])"
```

Re-slicing a slice doesn't make a copy of the underlying array. 
The full array will be kept in memory until it is no longer referenced.
Occasionally this can cause the program to hold all the data in memory when only a small piece of it is needed.
Since the slice references the original array, as long as the slice is kept around the garbage collector can't release the array.
To fix this problem, one can copy the interesting data to a new slice before returning it.