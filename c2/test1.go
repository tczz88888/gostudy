package main

import (
	"c2/tempconv"
	"fmt"
)

func main() {
	fmt.Println(tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.BoilingC)
	fmt.Println(tempconv.FreezingC)
	tp := tempconv.Celsius(123.3)
	fmt.Println(tp)
}
