package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	last := 0
	current := 0
	return func() int {
		sum := last + current
		if sum == 0 {
			current = 1
		} else {
			last = current
			current = sum
		}
		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
