package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type UniqueStringSet map[string]struct{}
type StringSetMap map[string]UniqueStringSet

func (m StringSetMap) Add(key string, value string) {
	if _, exists := m[key]; !exists {
		m[key] = make(UniqueStringSet)
	}
	m[key][value] = struct{}{}
}

func (m StringSetMap) Get(key string) []string {
	set, exists := m[key]
	if !exists {
		return nil
	}
	result := make([]string, 0, len(set))
	for item := range set {
		result = append(result, item)
	}
	return result
}

func (m StringSetMap) Count(key string, value string) bool {
	result := m.Get(key)
	return slices.Contains(result, value)
}

// func main() {
// 	counts := make(StringSetMap)
// 	for _, filename := range os.Args[1:] {
// 		data, err := os.ReadFile(filename)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
// 			continue
// 		} else {
// 			for line := range strings.SplitSeq(string(data), "\n") {
// 				counts.Add(line, filename)
// 			}
// 		}
// 	}
// 	for key := range counts {
// 		fmt.Println(key, "         ", counts.Get(key))
// 	}
// }

func countLines(f *os.File, counts map[string]int, bfiles StringSetMap) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		val := input.Text()
		counts[val]++
		bfiles.Add(f.Name(), val)
	}
}

func main() {
	counts := make(map[string]int)
	bfiles := make(StringSetMap)
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg) //俩返回值，一个是文件指针，一个是err，如果打开错误err就不等于nil
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts, bfiles)
		f.Close()
	}
	for line, n := range counts {
		if n >= 1 {
			fmt.Printf("%s\t%d\n", line, n)
			var result, seq string
			for key := range bfiles {
				if bfiles.Count(key, line) {
					//	fmt.Println("sb............")
					result += seq + key
					seq = " "
				}
			}
			fmt.Println(result)
		}
	}
}
