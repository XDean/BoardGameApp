package main

import (
	"fmt"
)

func main() {
	a := make(map[int]int)
	a[0] -= 1
	fmt.Println(a)
}

type A struct {
	i int
}

func (a A) copy() A {
	return a
}
