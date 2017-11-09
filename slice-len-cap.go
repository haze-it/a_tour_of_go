package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
    s := []int{1,2,3,4,5,6}
    printSlice(s)

    s = s[:0]
    printSlice(s)

    s = s[:6]
    printSlice(s)

    s = s[2:4]
    printSlice(s)

    s = s[0:4]
    printSlice(s)

    s = s[:5]
    printSlice(s)
}
