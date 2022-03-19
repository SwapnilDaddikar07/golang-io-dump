package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("/Users/swapnildaddikar/Workspace/go-practise/io/ReadFileLinByLine/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := sc.Text()
		fmt.Println(line)
	}
}
