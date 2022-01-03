package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	encodeDecode()
	marshalUnmarshal()
}

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("name :%s, age :%d", p.Name, p.Age)
}

func marshalUnmarshal() {

	const jsonData = `{"Name": "Alice", "Age": 25}`

	var p person

	err := json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("unmarshalled", p)

	bs, err := json.Marshal(p)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("marshalled", string(bs))
}

func encodeDecode() {

	log.Println("encode/decode")

	const jsonData = `
    {"Name": "Alice", "Age": 25}
    {"Name": "Bob", "Age": 22}
`

	reader := strings.NewReader(jsonData)
	writer := os.Stdout

	encoder := json.NewEncoder(writer)
	decoder := json.NewDecoder(reader)

	for {

		var m map[string]interface{}

		err := decoder.Decode(&m)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		for key, _ := range m {
			if key == "Age" {
				delete(m, key)
			}
		}

		err = encoder.Encode(&m)
		if err != nil {
			log.Fatal(err)
		}

	}

}
