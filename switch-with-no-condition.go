package main

import (
    "fmt"
    "time"
)

func main(){
    t := time.Now()
    var condition string
    switch {
        case t.Hour() < 12:
            condition = "Good morning"
        case t.Hour() < 17:
            condition = "Good afternooon"
        default:
            condition = "Good evening"
    }
    fmt.Println(condition)
}
