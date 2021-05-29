package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// https://golang.org/pkg/encoding/csv/#example_Writer
func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")

	fileZip, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writerZip := gzip.NewWriter(fileZip)
	writerZip.Header.Name = "test.txt"
	io.WriteString(writerZip, "gzip.Writer example\n")
	writerZip.Close()

	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", " ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world"})

	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", " ヘッダーも追加できます ")
	request.Write(os.Stdout)
}
