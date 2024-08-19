package main

import (
    "fmt"
    _"strings"
)

func main() {
    //var n int
    var s string
   // _, _ = fmt.Scan(&n)
    _, _ = fmt.Scan(&s)
    var count int
    
    for i:=0; i < len(s); i+=2 {
        end := i + 2
        if len(s) < end {
            end = len(s)
        }
        //fmt.Println(s[i:end])
        if s[i] != 'x' {
            if s[i:end] == "xx" {
            fmt.Println(s[i:end])
            count = 0
            }
        }
    }
    fmt.Println(count)
}
