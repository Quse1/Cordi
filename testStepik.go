package main 

import (
    "fmt"
    "strconv"
)

func main() {
    var n int 
    var n2 int
    var n31 int
    var n21 int
    var count int
    //var sum int 
    _,_ = fmt.Scan(&n)
    s:= make([]string, n)
    snum := make([]int, n)
    
    for i:=0; i < len(s); i++ {
        _, _ = fmt.Scan(&s[i])
        num, _ := strconv.Atoi(s[i])
        snum[i] = num
    }
    
    if n == 1 {
        count++ 
    } else {
      for j:= 0; j < len(snum)-1; j++ {
        for k:= j+1; k < len(snum); k++ {
            //fmt.Println(snum[j], snum[k])
            if snum[j] == 2 && snum[k] == 2 {
                count++
                fmt.Println(count)
                //fmt.Println(count)
                if snum[j] == snum[k] {
                    n2 = snum[j] + snum[k]
                    if n2 == 4 {
                        count--
                        //fmt.Println(count)
                    }
            }
        } else if snum[j] == 3 && snum[k] == 1 || snum[j] == 1 && snum[k] == 3 {
                n31 = snum[j] + snum[k]
                if n31 == 4 {
                    count++
                    //fmt.Println(count)
                }
        } else if snum[j] == 3 && snum[k] == 3 {
                    count++
                    //fmt.Println(count)
        } else if snum[j] == 4 && snum[k] == 4 {
                count++
                //fmt.Println(count)
        } else if snum[j] == 2 && snum[k] == 1 || snum[j] == 1 && snum[k] == 2 {
            n21 = snum[j] + snum[k]
            if n21 == 3 {
                count++
            }
        }
    }
}  
    }
    
fmt.Println(count)

}
