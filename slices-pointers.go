package main

import "fmt"

func main(){
    numbers := [4]string{
        "one", "two", "three", "faur",
    }
    fmt.Println(numbers)

    a := numbers[0:2]
    b := numbers[1:3]
    fmt.Println(a,b)

    b[0] = "XXX"
    fmt.Println(numbers)

    numbers[2] = "3"
    fmt.Println(a,b)
}
