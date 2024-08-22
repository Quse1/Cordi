package main

import (
	"fmt"
)

func main() {

	var t string
	_, _ = fmt.Scan(&t)
	nt := make([]string, 0)

	var dash bool = false

	var s [3]string = [3]string{".", "-.", "--"}

	for i := 0; i < len(t); i++ {
		if string(t[i]) == "-" && !dash {
			dash = true
		} else if string(t[i]) == "-" && dash {
			nt = append(nt, "--")
			dash = false
		} else if string(t[i]) == "." && dash {
			nt = append(nt, "-.")
			dash = false
		} else if string(t[i]) == "." && !dash {
			nt = append(nt, ".")
		}

	}

	for _, v := range nt {
		for i, k := range s {
			if v == k {
				fmt.Print(i)
			}

		}
	}
}
