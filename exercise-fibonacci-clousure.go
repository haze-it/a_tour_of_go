package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x1, x2 := 0, 1

	return func() int {
		result := x1
		x1 = x2
		x2 += result
		
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Println(f())
	}
}

