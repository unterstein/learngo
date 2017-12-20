package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, error := http.Get(url)
		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", error)
			os.Exit(1)
		}
		body, error := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if error != nil {
			fmt.Fprintf(os.Stderr, "reading: %v\n", error)
			os.Exit(1)
		}
		fmt.Printf("%s\n", body)
	}
}