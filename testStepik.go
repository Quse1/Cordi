package main 

import (
    "fmt"
    "strconv"
)

func main() {
    var n int 
    var n2 int
    var n3 int
    //var count int
    //var sum int 
    _,_ = fmt.Scan(&n)
    s:= make([]string, n)
    var snum []int
    for i:=0; i < len(s); i++ {
        _, _ = fmt.Scan(&s[i])
        num, _ := strconv.Atoi(s[i])
        snum = append(snum, num)
    }
    
    for j:= 0; j < len(snum); j++ {
        for k:= j+1; k < len(snum); k++ {
            //fmt.Println(snum[j], snum[k])
            if snum[j] < 3 && snum[k] < 3 && snum[j] != 1 && snum[k] != 1   {
                if snum[j] == snum[k] {
                    n2 = snum[j] + snum[k]
                    //fmt.Println(snum[j], snum[k])
                    fmt.Println(n2)
            }
        } else if snum[j] == 3 && snum[k] == 1 || snum[j] == 1 && snum[k] == 3 {
                n3 = snum[j] + snum[k]
                fmt.Println(n3)
            }
}
}
}
