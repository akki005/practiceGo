package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"time"
)

func main() {
	tokenBytes := make([]byte, 8)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		log.Fatal("getToken %w", err)
	}
	tokenHex := hex.EncodeToString(tokenBytes)
	tokenString := tokenHex + strconv.FormatInt(time.Now().Unix(), 10)

	h := md5.New()
	io.WriteString(h, tokenString)
	md5New := h.Sum(nil)

	md5Hex := hex.EncodeToString(md5New)

	md5HexR1 := regexp.MustCompile("\\+").ReplaceAllString(md5Hex, "-")
	md5HexR2 := regexp.MustCompile("\\/").ReplaceAllString(md5HexR1, "_")
	md5HexR3 := regexp.MustCompile("=").ReplaceAllString(md5HexR2, "")

	fmt.Println(md5HexR3)
}
