package main

import (
	"log"
	"net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/create_object", createObject)
    mux.HandleFunc("/user", objectQuery)

    // Start listen on :4000
    log.Println("Start listening on port 4000")
    err := http.ListenAndServe("localhost:4000", mux)

    // call os.exit(1) immediately after writing the message
    log.Fatal(err)
}
