package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Please provice source and destination files")
	}
	sourcepath := os.Args[1]
	destinationpath := os.Args[2]
	source, err := os.Open(sourcepath)
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()

	destination, _ := os.OpenFile(destinationpath, os.O_RDWR|os.O_CREATE, 0666)
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Copying from", sourcepath, "to", destinationpath, "has been successfull!")
	}
}
