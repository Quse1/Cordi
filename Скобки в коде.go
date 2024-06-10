package main

import (
	"fmt"
)

func main() {
	var s string = "{"
	var count int = 1

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

	if count%2 == 0 {
		fmt.Println(len(s))
	} else {
		fmt.Println("Success")
	}
}
