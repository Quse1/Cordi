package main

import (
	"fmt"
)

func main() {
	s := make([]string, 2)
	var lcase []string
	var idx1, idx2 int
	for i := 0; i < len(s); i++ {
		_, _ = fmt.Scan(&s[i])
	}

	if s[0] == s[1] {
		fmt.Println(s[0])
		return
	} else {
		for i := 97; i < 123; i++ {
			lcase = append(lcase, string(i))
		}

		for j := 0; j < len(lcase); j++ {
			switch lcase[j] {
			case s[0]:
				idx1 = j
			case s[1]:
				idx2 = j
			}
		}
	}
	if idx1 < idx2 {
		for _, k := range lcase[idx1 : idx2+1] {
			fmt.Print(k, " ")
		}
	} else if idx1 > idx2 {
		for _, k := range lcase[idx2 : idx1+1] {
			fmt.Print(k, " ")
		}
	}
}
