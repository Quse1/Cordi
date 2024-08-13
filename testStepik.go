// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    
    )

func alphabet (t []string) string {
    var res string
    var alf []string = []string{"_", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
    
    for i:=0; i < len(alf); i++ {
        for j:=i; j < i+1; j++ {
            if alf[i] == t[j] {
                res = t[j]
            }
        }
    }
    return res
}

func main() {
    input := bufio.NewScanner(os.Stdin)
    _ = input.Scan()
    text := input.Text()
    t := strings.Split(text, "")
    fmt.Println(alphabet(t))
    
    
}
