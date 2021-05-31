package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var reader io.Reader = strings.NewReader(" テストデータ ")
var readCloser io.ReadCloser = ioutil.NopCloser(reader)

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}

func readChunk(file *os.File) []io.Reader {
	var chunks []io.Reader
	file.Seek(8, 0)
	var offset int64 = 8
	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		fmt.Printf("%d\n", int64(length))
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))
		offset, _ = file.Seek(int64(length+8), 1)
		//fmt.Printf("%d\n", offset)
	}
	return chunks
}

func textChunk(text string) io.Reader {
	byteText := []byte(text)
	var buffer bytes.Buffer

}

func main() {
	file, err := os.Open("Lenna.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunk(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
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
	/*
		file, err := os.Open("main.go")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		io.Copy(os.Stdout, file)
	*/
	/*
		conn, err := net.Dial("tcp", "ascii.jp:80")
		if err != nil {
			panic(err)
		}
		conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
		res, err := http.ReadResponse(bufio.NewReader(conn), nil)
		fmt.Println(res.Header)
		defer res.Body.Close()
		io.Copy(os.Stdout, res.Body)
		// time.Sleep(time.Second * 20)
		reader := strings.NewReader("Example of io.SectionReader\n")
		sectionReader := io.NewSectionReader(reader, 14, 7)
		io.Copy(os.Stdout, sectionReader)
		// 32 ビットのビッグエンディアンのデータ(10000)
		data := []byte{0x0, 0x0, 0x27, 0x10}
		var i int32
		// エンディアンの変換
		binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
		fmt.Printf("data: %d\n", i)
		fmt.Printf("data: %d\n", data)
	*/
}
