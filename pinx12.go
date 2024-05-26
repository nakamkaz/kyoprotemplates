package main

/*
## N
## a1 a2 a3 a4 a5 ... aN
*/
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	stdin.Scan()
	secondline := strings.Split(stdin.Text(), " ")
	arr := make([]int, N)
	for i, val := range secondline {
		arr[i], _ = strconv.Atoi(val)
	}
	fmt.Println(arr)

}
