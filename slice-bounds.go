package main

import "fmt"

func main() {
    s := []int{1,2,3,4,5,6}

    s = s[1:4]
    fmt.Println(s)

    s = s[:]
    fmt.Println(s)
}
