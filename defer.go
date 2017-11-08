package main

import "fmt"

func main(){
    defer fmt.Println("world")
    defer fmt.Println("^o^")
    defer fmt.Println("^_^")


    fmt.Println("Hello")
}
