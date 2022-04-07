package main

/**
https://coderwall.com/p/zyxyeg/golang-having-fun-with-os-stdin-and-shell-pipes
I am told that this should work with something like `date | [program]` and not curl since it takes a bit,
	but I was not able to make this example work

**/

import (
	"os"
	"fmt"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Size() > 0 {
		fmt.Println("There is something to read")
	} else {
		fmt.Println("stdin is empty")
	}
}
