package main

/*
## N
## 1 3
## 3 5
## 7 98
## ..
## 3 0
*/
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Tp struct {
	Ia []int
}

func dout(msg ...interface{}) {
	if os.Getenv("DEBUG") == "1" {
		log.Print(msg)
	}
}
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	firstline := stdin.Text()
	narr := strings.Split(firstline, " ")
	N, _ := strconv.Atoi(narr[0])
	arr := make([]Tp, N)
	for i := 0; i <= N-1; i++ {
		stdin.Scan()
		line := strings.Split(stdin.Text(), " ")
		loIa := make([]int, 2)
		for j, val := range line {
			loIa[j], _ = strconv.Atoi(val)
		}
		arr[i] = Tp{Ia: loIa}
	}
	fmt.Println(arr)

}
