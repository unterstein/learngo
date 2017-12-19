package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var separator, result string = "", ""
	for i := 1; i < len(os.Args); i++ {
		result += separator + os.Args[i]
		separator = ", "
	}
	fmt.Println(result)

	s, sep := "", "" // weird declaration, but ok
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ", "
	}
	fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], ", "))
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[2:])
	fmt.Println(os.Args[0:])
}