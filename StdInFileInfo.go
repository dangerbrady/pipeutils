package main

import (
	"os"
	"fmt"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", fi.Mode())
}
