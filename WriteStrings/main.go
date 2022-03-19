package main

import (
	"io"
	"log"
	"os"
)

func main() {

	values := []string{"this is first", "this is second", "this is third", "this is forth"}

	f, createErr := os.Create("WriteStrings/newfile.txt")
	if createErr != nil {
		log.Fatalln(createErr)
	}

	defer f.Close()

	for _, v := range values {
		io.WriteString(f, v)
		//Write to new line.
		io.WriteString(f,"\n")
	}

}
