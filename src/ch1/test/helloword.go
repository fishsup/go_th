package main

import (
	"fmt"
	"os"
)

func main() {
	var result = "Hello word"
	if len(os.Args) > 1 {
		fmt.Println(result, os.Args[1])
	} else {
		fmt.Println(result)
	}
	os.Exit(0)
}
