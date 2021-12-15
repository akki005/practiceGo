package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	h := md5.New()
	text := "this is string"
	io.WriteString(h, text)
	final := h.Sum(nil)
	fmt.Printf("%s -> %x\n", text, final)
	fileToMd5()
}

func fileToMd5() {
	filePath := "/home/akash/InterviewPrep/LanguagePractice/GO/hash/main.go"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s -> %x\n", filePath, h.Sum(nil))
}
