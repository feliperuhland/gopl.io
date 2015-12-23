package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("File: " + os.Args[0])
	for index, line := range os.Args[1:] {
		fmt.Println(index, line)
	}
}
