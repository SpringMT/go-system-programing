package main

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var reader io.Reader = strings.NewReader(" テストデータ ")
var readCloser io.ReadCloser = ioutil.NopCloser(reader)

func main() {
	/*
		for {
			buffer := make([]byte, 5)
			size, err := os.Stdin.Read(buffer)
			if err == io.EOF {
				fmt.Println("EOF")
				break
			}
			fmt.Printf("size=%d input='%s'\n", size, string(buffer))
		}
	*/
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
