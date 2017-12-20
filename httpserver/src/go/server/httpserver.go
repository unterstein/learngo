package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "(-:")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":" + os.Args[1], nil)
}