package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	reader := strings.NewReader("read this text")

	buf := make([]byte, 10)

	for {

		n, err := reader.Read(buf)

		if err == io.EOF {
			break
		}

		fmt.Printf("reading first %d : %s\n", n, buf[:n])

	}

}
