package main 

import (
    "fmt"
    "strconv"
)

func main() {
    var n int 
    var count int
    //var sum2 int
    _,_ = fmt.Scan(&n)
    s:= make([]string, n)
    snum := make(map[int]int, n)
    
    for i:=0; i < len(s); i++ {
        _, _ = fmt.Scan(&s[i])
        num, _ := strconv.Atoi(s[i])
        //snum[i]+=num
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
    fmt.Println(snum)
    for snum[4] > 0 {
        count++
        snum[4]--
    }
    if snum[3] > 0 && snum[1] > 0 {
        for snum[3] + snum[1]  > 0 {
        count++
        snum[3]--
        snum[1]--
    }
    
    } else if snum[3] > 0 && snum[1] == 0 {
        for snum[3]> 0 {
        count++
        snum[3]--
    }
    
    } else if snum[2] > 0 && snum[1] == 0 {
        if snum[2] >=2 {
           for snum[2] >= 2 {
        count++
        snum[2]--
    }  
        } else if snum[2] ==1 {
            count++
            snum[2]--
            }
    } else if snum[2] > 0 && snum[1] > 0 {
        for snum[2] + snum[1] > 0 {
        count++
        snum[2]--
        snum[1]--
        }
     }else if snum[3] == 0 && snum[1] > 0 {
        for snum[1] >= 4 {
        count++
        snum[1]--
    }
     }
    
    
   
    fmt.Println(snum)
    fmt.Println(snum[1]+snum[2]+snum[3]+snum[4])
    fmt.Println(count)
    
    /*if n == 1 {2
        count++ 
    } else {
      for j:= 0; j < len(snum)-1; j++ {
        for k:= j+1; k < len(snum); k++ {
            fmt.Println(snum[j], snum[k])
            if snum[j] + snum[k] == 4 {
                fmt.Println(snum[j] + snum[k])
            } else if snum[j] + snum[k] == 3 {
                fmt.Println(snum[j] + snum[k])
            } else if snum[j] > 2 && snum[j] <= 4 {
                fmt.Println(snum[j])
            }
        }
      }
}*/

}
