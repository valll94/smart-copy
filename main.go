package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var FILESIZE int64

var Sourcepath string
var Destinationpath string

var BYTESWRITTEN int64

func main() {
	BYTESWRITTEN = 0

	if len(os.Args) != 3 {
		log.Fatal("Please provice source and destination files")
	}
	Sourcepath = os.Args[1]
	Destinationpath := os.Args[2]
	if sourceFileStat, err := os.Stat(Sourcepath); err != nil {
		log.Fatal("Source file ", Sourcepath, " do not exist")
	} else {
		FILESIZE = sourceFileStat.Size()
	}

	if _, err := os.Stat(Destinationpath); err == nil {
		log.Fatal("Destination file ", Destinationpath, " exist")
	}
	source, err := os.Open(Sourcepath)
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()

	destination, _ := os.OpenFile(Destinationpath, os.O_RDWR|os.O_CREATE, 0666)
	defer destination.Close()
	go func() {
		for {
			time.Sleep(time.Second)
			BYTESWRITTEN, _ := os.Stat(Destinationpath)
			progress := float64(BYTESWRITTEN.Size()) / float64(FILESIZE) * 100
			fmt.Printf("\rProgress: %.2f%%", progress)

		}
	}()
	_, err = io.Copy(destination, source)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("\nCopying from", Sourcepath, "to", Destinationpath, "has been successfull!")
	}

}
