package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	input := myscanner.Text()

	stringNumbers := strings.Split(input, " ")
	rowsCounts, _ := strconv.Atoi(stringNumbers[0])
	columnsCount, _ := strconv.Atoi(stringNumbers[1])

	a := make([][]int, rowsCounts)
	for i := 0; i < rowsCounts; i++ {
		a[i] = make([]int, columnsCount+1)
		for j := 0; j < columnsCount; j++ {
			if i == 0 && j == 0 || i == 0 || j == 0 {
				a[i][j] = 1
			} else {
				a[i][j] = a[i][j-1] + a[i-1][j]
			}

		}

	}

	for i := 0; i < rowsCounts; i++ {
		for j := 0; j < columnsCount; j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
}
