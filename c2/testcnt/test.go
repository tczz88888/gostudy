package main

import (
	"bufio"
	"c2/popcount"
	"fmt"
	"os"
	"strconv"
)

func show(x string) {
	f, err := strconv.ParseUint(x, 10, 64)
	if err != nil {
		fmt.Printf("popcount: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d\n%d\n%d\n", popcount.PopcountByArray(f), popcount.PopcountBybit(f), popcount.PopcountByXor(f))
}
func main() {
	len := len(os.Args)
	if len == 1 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			show(input.Text())
		}
	} else {
		for _, arg := range os.Args[1:] {
			show(arg)
		}
	}
}
