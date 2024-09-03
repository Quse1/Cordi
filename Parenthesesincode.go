/* Скобки расставлены не правильно. В первую очередь необходимо найти закрывающую скобку, для которой нет
соответствующей открывающей скобки, либо она закрывает не соответсвуюшую ей открывающую скобку.
Либо найти первую открывающую скобку, для которой нет соответсвуюшуей ей закрывающейся скобки*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func CorrectParenthesesBrackets(s string) (int, string) {
	stack := make([]rune, 0)
	match := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	for i, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else if ch == ')' || ch == ']' || ch == '}' {
			if len(stack) == 0 || match[stack[len(stack)-1]] != ch {
				return i + 1, "closing" // индекс последней закрывающейся скобки
			}
			stack = stack[:len(stack)-1] //  открывающиеся  скобки без закрывающихся
		}
	}
	if len(stack) > 0 {
		return len(stack), "opening" // индекс последней открывающейся скобки

	}

	return 0, "" // все парные скобки открывающейся и закрывающейся
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	_ = input.Scan()
	text := input.Text()
	idx, types := CorrectParenthesesBrackets(text)

	if idx == 0 && types == "" {
		fmt.Println("Success")
	} else if types == "closing" {
		fmt.Println(idx)
	} else if types == "opening" {
		fmt.Println(idx)
	}

}
