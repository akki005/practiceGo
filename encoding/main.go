package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "rzp_test_14CharRZPId000:passworkC"
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(encoded)
	decoded, _ := base64.StdEncoding.DecodeString("cnpwX3Rlc3Rfd2hpdGVsaXN0bWVyY2g6cGFzc3dvcmQy04")
	fmt.Println(string(decoded))
}
