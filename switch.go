package main

import (
    "fmt"
    "runtime"
)

func main(){
    var os_name string
    switch os := runtime.GOOS; os {
        case "darwin":
            os_name = "OS_X"
        case "linux":
            os_name = "Linux"
        default:
            os_name = string(os)
    }
    fmt.Print("Go runs on ")
    fmt.Println(os_name)
}
