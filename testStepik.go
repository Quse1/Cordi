package main 

import (
    "fmt"
    "strconv"
)

func main() {
    var n int
    var count int
    _,_ = fmt.Scan(&n)
    s:= make([]string, n)
    snum := make(map[int]int, n)
    
       for i:=0; i < len(s); i++ {
        _, _ = fmt.Scan(&s[i])
        num, _ := strconv.Atoi(s[i])
        switch num {
            case 1:
            snum[num]++
            case 2:
            snum[num]++
            case 3:
            snum[num]++
            case 4:
            snum[num]++
        }
    } 
    if n == 1 {
        count = 1
    } else {
        //fmt.Println(snum)
        for snum[4] > 0 {
        count++
        snum[4]--
    }
    if snum[3] > 0  {
        for snum[3] > 0 {
        count++
        snum[3]--
        if snum[1] > 0 {
            snum[1]--
        }
    }
    }
    if snum[2] >= 2 {
        for snum[2] > 1 {
            count+= snum[2] / 2 
            snum[2]-=2
           }
    }
    if snum[2] > 0  {
        for snum[2] > 0  {
        count++
        snum[2]--
        if snum[1] > 0  {
            for snum[1] > 0 {
                snum[1]--
            }
            
        }
        }
     }
     if snum[1] >= 4 {
        count += snum[1] / 4
     }
     
     if snum[1] <= 2 && snum[1] != 0   {
        count++
        snum[1] = 0
        if snum[2] > 0 {
            snum[2]--
        }
     }
     
     if snum[1] == 3 {
         count++
         snum[1] = 0
     }
     
    }
    //fmt.Println(snum)
    fmt.Println(count)
}
