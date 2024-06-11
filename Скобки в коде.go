package main

import (
	"fmt"
)

func main() {
	var s string = "foo(bar);"
	var count int = 1
	res := int(math.Pow(10, 5))
	if 1 <= len(s) && len(s) <= res {
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case '[':
				count++
			case ']':
				count++
			case '(':
				count++
			case ')':
				count++
			case '{':
				count++
			case '}':
				count++
			}
		}

	}

	if count%2 == 0 {
		fmt.Println(len(s))
	} else {
		fmt.Println("Success")
	}

}
