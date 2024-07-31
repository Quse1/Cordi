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
		a[i] = make([]int, columnsCount)
		for j := 0; j < columnsCount; j++ {
			a[i][j] = (i + 1) * (j + 1)
		}
	}

	for i := 0; i < rowsCounts; i++ {
		for j := 0; j < columnsCount; j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
}
