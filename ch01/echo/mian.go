package main

import (
	"fmt"
	"os"
)

func main() {
	var s, step string
	for i := 1; i < len(os.Args); i++ {
		s += step + os.Args[i];
		step = " ";
	}
	fmt.Println("输出结果：" + s);
}