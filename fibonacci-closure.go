package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci(size int) func(int) int {
	sl := make([]int, size)
	return func(x int) int {
		if x < 2 {
			sl[x] = x
			return x
		} else {
			sl[x] = sl[x-1] + sl[x-2]
			return sl[x]
		}
	}
}

func main() {
	size := 10
	f := fibonacci(size)
	for i := 0; i < size; i++ {
		fmt.Println(f(i))
	}
}
