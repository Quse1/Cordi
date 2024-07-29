package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	// Создание массива
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, n)
	}

	// Заполнение массива
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			if i+j == n-1 {
				arr[i][j] = 1 // Элемент на побочной диагонали
			} else if i+j < n-1 {
				arr[i][j] = 0 // Элемент выше побочной диагонали
			} else {
				arr[i][j] = 2 // Элемент ниже побочной диагонали
			}
		}
	}
	// Вывод массива
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}
