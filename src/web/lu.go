package main

import (
    "flag"
	"log"
	"net/http"
)

func main() {
    /* getting args configuration parameters
     * setting addr flag. */
    addr := flag.String("addr", "localhost:4000", "HTTP Network Address")
    flag.Parse() // Getting the value for addr variable. Now addr points to the real value

    mux := http.NewServeMux()

    /* Here we are chaining handlers in a mux (a special handler that calls
    * another handlers.
    */
    mux.HandleFunc("/", home)
    mux.HandleFunc("/create_object", createObject)
    mux.HandleFunc("/user", objectQuery)

    // creating fileserver which deals 
    fileServer := http.FileServer(http.Dir("./ui/static"))
    // set mux to handle with /static prefix
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    // Start listen on :4000
    // addr is a pointer to the network address value. We need to dereference
    // it using * operator
    log.Printf("Start listening on port %s\n", *addr)
    err := http.ListenAndServe(*addr, mux)

    // call os.exit(1) immediately after writing the message
    log.Fatal(err)
}
