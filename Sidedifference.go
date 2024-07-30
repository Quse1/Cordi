package main

import "fmt"

func main() {
    var n int
    _, _ = fmt.Scan(&n)
    // Создание массива
    a := make([][]int, n)
    for i := 0; i < len(a); i++ {
        a[i] = make([]int, n)
    }
    
    // Заполнение массива
    for i:=0; i < n; i++  {
        for j:=0; j < n; j++ {
            if i + j == n - 1 {
                a[i][j] = 1 // Элемент на побочной диагонали
            } else if i + j < n - 1 {
                a[i][j] = 0 // Элемент выше побочной диагонали
            } else {
                a[i][j] = 2 // Элемент ниже побочной диагонали
            }
    }
}
// Вывод массива
 for i:=0; i < n; i++  {
        for j:=0; j < n; j++ {
            fmt.Print(a[i][j], " ")
        }
        fmt.Println()
    }
}
