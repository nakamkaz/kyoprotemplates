package main

/*
# N
# a1
# a2
# ..
# aN
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
	narr := strings.Split(firstline," ")
	N, _ := strconv.Atoi(narr[0])
	arr := make([]int, N)
	for i:=0; i <= N-1; i++ {
		stdin.Scan()
		arr[i],_ = strconv.Atoi(stdin.Text())
	}
	fmt.Println(arr)

}
