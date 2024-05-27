package main

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
)

type T struct {
	val []int
}

func fill(rowlen int) T {
	val := make([]int, rowlen)
	for j, _ := range val {
		val[j] = 0
	}
	t := T{val}
	return t
}
func  trect(collen int, rowlen int) []T {
	vt := make([]T, collen)
	for i, _ := range vt {
		vt[i] = fill(rowlen)
	}
	return vt
}

func (t *T) put(idx int,intval int){
	t.val[idx]=intval
}
func (t *T) square() {
	for i,_ := range t.val {
	t.val[i]=t.val[i]*t.val[i]
	}
}
func get(ta []T, colidx, rowidx int) int {
	return ta[colidx].val[rowidx]
}
/*
H W
a1
a2
a3
a4
aN
*/
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	firstline := strings.Split(stdin.Text(), " ")
	N,_ := strconv.Atoi(firstline[0])
	va := trect(N,1)
	for i:=0; i<=N-1;i++{
		stdin.Scan()
		lineval,_ := strconv.Atoi(stdin.Text())
		va[i].put(0,lineval)
	}
	fmt.Println(va)
}
