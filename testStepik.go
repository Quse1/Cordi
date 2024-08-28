package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	symb := map[string]int {
	    "(": 0,
	    ")": 0,
	    "[": 0,
	    "]": 0,
	    "{": 0,
	    "}": 0,
	}
	var count int
	input := bufio.NewScanner(os.Stdin)
	_ = input.Scan()
	text := input.Text()
	//fmt.Println(text)
	
	for i:=0; i < len(text); i++ {
	    switch string(text[i]) {
	        case "(":
	        symb["("]++
	        case ")":
	        symb[")"]++
	        case "[":
	        symb["["]++
	        case "]":
	        symb["]"]++
	        case "{":
	        symb["{"]++
	        case "}":
	        symb["}"]++
	    }
	}
	
	
	if symb["("] == symb[")"] {
	    count++
	}
	if symb["["] == symb["]"] {
	    count++
	}
	if symb["{"] == symb["}"] {
	    count++
	}
	
	if count == 3 {
	    fmt.Println("Success")
	}

}
