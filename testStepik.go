package main

import (
	"fmt"
	"math"
)

func Reminder(year int) int {
	var length, fnum int
	//var mnum1 []int
	length = int(math.Log10((float64(year)))) + 1 // длина вводимого числа

	for i := length - 1; i >= 0; i-- {
		fnum = int(float64(year)/math.Pow(10, float64(i))) % 10
		//mnum1 = append(mnum1, fnum)
		fmt.Print(fnum)
	}
	return fnum
}

func main() {
	var year int
	_, _ = fmt.Scan(&year)
	Reminder(year)

}
