package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func Popcount(array [64]byte) int {
	ans := 0
	for i := 0; i < 64; i++ {
		for array[i] != 0 {
			ans++
			array[i] = array[i] & (array[i] - 1)
		}
	}
	return ans
}

var op = flag.Int("S", 256, "选择sha的方式,默认256,可选384,512")
var val = flag.String("V", " ", "选择生成摘要的字符串")

func main() {
	flag.Parse()
	var c [64]byte
	if *op == 256 {
		cg := sha256.Sum256([]byte(*val))
		copy(c[:], cg[:])
	} else if *op == 384 {
		cg := sha512.Sum384([]byte(*val))
		copy(c[:], cg[:])
	} else if *op == 512 {
		cg := sha512.Sum512([]byte(*val))
		copy(c[:], cg[:])
	} else {
		os.Exit(1)
	}
	ans := Popcount(c)
	fmt.Printf("sha%d: %8x\npopcount:%d\n", (*op), c, ans)
}
