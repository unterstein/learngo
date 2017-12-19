package main

import (
	"fmt"
	"os"
)

func main() {
	var separator, result string
	for i := 1; i < len(os.Args); i++ {
		result += separator + os.Args[i]
		separator = ", "
	}
	fmt.Println(result)
}