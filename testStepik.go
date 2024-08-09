package main 

import (
    "fmt"
    "strconv"
)

func main() {
    var n int
    var count int
    var count1 int
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
        for snum[4] > 0 {
        count++
        snum[4]--
    }
    if snum[3] > 0 && snum[1] > 0 {
        for snum[3] + snum[1] > 0 {
        count++
        snum[3]--
        snum[1]--
        
    }
    
    } else if snum[3] > 0 && snum[1] <= 0 {
        count+=snum[3]
        snum[3]--
    }
    if snum[2] >= 2 {
        for snum[2] > 1 {
            count+= snum[2] / 2 
            snum[2]-=2
           }
    } else if snum[2] > 0 && snum[1] > 0 {
        for snum[2] + snum[1] > 0 {
        count++
        snum[2]--
        snum[1]--
        }
        
     } else if snum[3] <= 0 && snum[2] <= 0 && snum[1] >= 4 {
        count += snum[1] / 4
        
     } 
     if snum[2] == 1 && snum[3] <= 0 && snum[4] <= 0 && snum[1] <= 0 {
         count+=snum[2]
         snum[2]-- 
     }
     
     if snum[1] <= 3 && snum[3] <= 0 && snum[4] <= 0 && snum[2] <= 0 {
          //fmt.Println(snum[1])
            count1 = snum[1]
            snum[1] =0
         if count1 == 3 || count1 == 2 || count1 ==1 {
             count++
         }
     }
    }
    fmt.Println(count)
}
