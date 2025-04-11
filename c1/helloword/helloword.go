package main //申明当前文件属于哪个包，入口文件必须声明为main包

import (
	"fmt" //导入包，和python的import，c的#include一样
	"os"
)

func main() {
	fmt.Println("Hello World!") //调用包的方法
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
