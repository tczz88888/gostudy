package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"
)

func t1() {
	fmt.Println(os.Args[0])
}

func t2() {
	for idx, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", idx, arg)
	}
}

func t3() {
	var strArray [5000]string
	for i := 0; i < 5000; i++ {
		strArray[i] = string(bytes.Repeat([]byte{'1'}, 5000))
	}
	var times = time.Now()
	var s1, seq string
	for _, arg := range strArray[0:5000] {
		s1 += seq + arg
		seq = " "
	}
	var time1 = time.Since(times)
	times = time.Now()
	var s2 = strings.Join(strArray[0:5000], " ")
	var time2 = time.Since(times)
	fmt.Println(s1 == s2)

	fmt.Println(time1.Seconds())
	fmt.Println(time2.Seconds())
}

func main() {
	t1()
	t2()
	t3()
}
