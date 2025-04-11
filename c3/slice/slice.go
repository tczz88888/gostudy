package main

import "fmt"

func main() {
	a := [...]int64{0: 3, 1: 2, 20: 20}
	b := a[0:5]
	a[0] = 1
	fmt.Println(a)
	fmt.Println(b)

}
