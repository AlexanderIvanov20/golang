package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type textField struct{}

func (tf textField) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))

	return len(bs), nil
}

func main() {
	name := os.Args[1]

	f, err := os.Open(name)
	tf := textField{}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	io.Copy(tf, f)
}
