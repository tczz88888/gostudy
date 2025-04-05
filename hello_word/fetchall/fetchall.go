package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	stime := time.Now()
	data, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, data.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(stime).Seconds()
	ch <- fmt.Sprintf("%.2fs   %7d   %s", secs, nbytes, url)
	fmt.Printf("%s done!\n", url)
}

func main() {
	stime := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //携程
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("tot time: %.2f", time.Since(stime).Seconds())
}
