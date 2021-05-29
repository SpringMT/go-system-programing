package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("2_1.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	fmt.Fprintf(writer, "数字 %d 文字列 %s 浮動小数点 %f \n", 123, "あああ", 0.12345)
	file.Close()
}
