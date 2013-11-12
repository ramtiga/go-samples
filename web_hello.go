package main

import (
    "flag"
    "net/http"
    "log"
    "fmt"
)

var port = flag.String("port", ":8080", "server port number")

func main() {
    flag.Parse()
    http.Handle("/", http.HandlerFunc(hello))
    err := http.ListenAndServe(*port, nil)
    if err != nil {
        log.Printf("ListenAndServe: %v", err)
    }
}

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Hello World!!!")
}
