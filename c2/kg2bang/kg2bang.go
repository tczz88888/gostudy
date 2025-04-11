package main

import (
	"bufio"
	"c2/weightconv"
	"fmt"
	"os"
	"strconv"
)

func show(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "kg2bang: %v\n", err)
		os.Exit(1)
	}
	kg := weightconv.Kg(t)
	bang := weightconv.Bang(t)
	fmt.Printf("%s->%s  %s->%s\n", kg, weightconv.Kg2Bang(kg), bang, weightconv.Bang2Kg(bang))
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
