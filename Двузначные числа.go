/*Напишите программу, которая выводит все двузначные числа, которые равны удвоенному произведению своих цифр.
В ответе запишите найденные числа через запятую без пробелов.*/

package main

import "fmt"

func main() {
	for i := 10; i < 99; i++ {
		for j := 1; j < 9; j++ {
			for k := 0; k < 9; k++ {
				if i*j == (j*10 + k) {
					fmt.Print(j*10+k, ",")
				}
			}
		}
	}
}
