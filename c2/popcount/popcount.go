package popcount

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopcountByArray(x uint64) int {
	var ans = 0
	times := time.Now()
	for h := 0; h <= 1000000000; h++ {
		ans = 0
		for i := 0; i < 8; i++ {
			ans += int(pc[byte(x>>(i*8))])
		}
	}
	fmt.Printf("f1 cos:%fsec\n", time.Since(times).Seconds())
	return ans
}

func PopcountBybit(x uint64) int {
	var ans = 0
	times := time.Now()
	for h := 0; h <= 1000000000; h++ {
		ans = 0
		for i := 0; i < 64; i++ {
			ans += int((x >> i) & 1)
		}
	}
	fmt.Printf("f1 cos:%fsec\n", time.Since(times).Seconds())
	return ans
}

func PopcountByXor(x uint64) int {
	var ans = 0
	times := time.Now()
	for h := 0; h <= 1000000000; h++ {
		ans = 0
		tmp := x
		for tmp != 0 {
			tmp = tmp & (tmp - 1)
			ans++
		}
	}
	fmt.Printf("f1 cos:%fsec\n", time.Since(times).Seconds())
	return ans
}
