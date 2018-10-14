package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("ch03/3.4.2/file.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
