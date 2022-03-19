package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("ReadFileLinByLine/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := sc.Text()
		fmt.Println(line)
	}
}
