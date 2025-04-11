package main

import (
	"fmt"
)

func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}

func main() {
	var input int
	fmt.Scan(&input)
	fmt.Println(Signum(input))
}
