package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, fileOpenError := os.Open("ReadEntireFileAtOnce/input.txt")
	if fileOpenError != nil {
		log.Fatalln(fileOpenError)
	}

	entireFile, fileReadError := io.ReadAll(f)
	if fileReadError != nil {
		log.Fatalln(fileReadError)
	}

	fmt.Println(string(entireFile))

}
