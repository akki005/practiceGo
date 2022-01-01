package main

import (
	"fmt"
	"io"
	"strings"
)

type MyCapReader struct {
	reader io.Reader
}

func (m MyCapReader) Read(bytes []byte) (int, error) {

	n, err := m.reader.Read(bytes)

	if err != nil {
		return 0, err
	}

	str := string(bytes)

	str = strings.ToUpper(str)

	for i := 0; i < len(str); i++ {
		bytes[i] = str[i]
	}

	return n, nil

}

func main() {
	str := "abcdefghijk"
	buf := make([]byte, len(str))
	strReader := strings.NewReader(str)
	readableStream := MyCapReader{strReader}
	readableStream.Read(buf)
	fmt.Println(string(buf))
}
