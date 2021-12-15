package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	str := "this needs to be encoded in base 64"
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(encoded)
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(string(decoded))
}
