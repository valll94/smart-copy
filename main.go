package main

import (
	"io"
	"os"
)

func main() {
	source, _ := os.Open("./test")
	defer source.Close()

	destination, _ := os.OpenFile("./asd", os.O_RDWR|os.O_CREATE, 0666)
	defer destination.Close()

	_, _ = io.Copy(destination, source)
}
