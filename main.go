package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: edn <json_string>")
		os.Exit(1)
	}

	fmt.Println(os.Args[1])
}
