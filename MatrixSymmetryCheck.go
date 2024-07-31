/*Чтобы проверить, является ли двумерный массив (матрица) симметричным относительно главной диагонали,
нужно убедиться, что для всех элементов матрицы выполняется условие:a[i][j]=a[j][i]*/

package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	// Создание массива
	var flag bool = true
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			_, _ = fmt.Scan(&a[i][j])
		}
	}

	// Заполнение массива
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] != a[j][i] {
				flag = false
			}

		}
	}
	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
