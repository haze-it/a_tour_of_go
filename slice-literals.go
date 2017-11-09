package main

import "fmt"

func main(){
    q := []int{1,2,3,4,5,6}
    fmt.Println(q)

    r := []bool{true,true,false,false,true,false}
    fmt.Println(r)

    s := []struct {
    i int
    b bool
    }{
        {2,true},
        {2,true},
        {2,true},
        {2,false},
        {2,true},
    }
    fmt.Println(s)
}
