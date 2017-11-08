package main

import (
    "fmt"
    "time"
)

func main(){
    today := time.Now().Weekday()
    var week string

    switch time.Saturday {
        case today + 0:
            week = "Today."
        case today + 1:
            week = "Tomorrow."
        case today + 2:
            week = "In two days.."
        default:
            week = "Too far away."
    }
    fmt.Println("When's Saturday?")
    fmt.Println(week)
}
