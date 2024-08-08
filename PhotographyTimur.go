package main

import "fmt"

func main() {
	var row, col int
	_, _ = fmt.Scan(&row)
	_, _ = fmt.Scan(&col)
	var color int
	var bkwt int
	a := make([][]string, row)
	for i := 0; i < row; i++ {
		a[i] = make([]string, col)
		for j := 0; j < col; j++ {
			_, _ = fmt.Scan(&a[i][j])
		}
	}
	//fmt.Println(a)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			//fmt.Printf("%T", a[i][j], " ")
			if a[i][j] == "W" || a[i][j] == "B" || a[i][j] == "G" {
				bkwt--
			} else if a[i][j] == "C" || a[i][j] == "M" || a[i][j] == "Y" {
				color++
			}

		}
	}

	if color == 0 {
		fmt.Println("#Black&White")
	} else if color > bkwt {
		fmt.Println("#Color")
	}
}
