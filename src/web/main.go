package main

import (
    "flag"
	"log"
	"net/http"
    "os"
)

type Application struct {
    errorLogger     *log.Logger
    infoLogger      *log.Logger
}

type Config struct {
    Addr        string
    StaticDir   string
}

func main() {
    /* Open the infolog file */
    f_info, err := os.OpenFile("/tmp/go-serv-info.log", os.O_WRONLY|os.O_CREATE, 0666)
    f_error, err := os.OpenFile("/tmp/go-serv-error.log", os.O_WRONLY|os.O_CREATE, 0666)

    /* Creating loggers */
    infoLogger := log.New(f_info, "INFO\t", log.Ldate|log.Ltime)
    errorLogger := log.New(f_error, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

    /* getting args configuration parameters
     * setting addr flag. */
    /*
    addr := flag.String("addr", "localhost:4000", "HTTP Network Address")
    flag.Parse() // Getting the value for addr variable. Now addr points to the real value
    */

    cfg := new(Config)
    flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP Network Address")
    flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to the static files directory")
    flag.Parse()

    mux := http.NewServeMux()

    /* Creating the application struct */
    app := &Application{
        errorLogger:    errorLogger,
        infoLogger:     infoLogger,
    }

    /* Here we are chaining handlers in a mux (a special handler that calls
    * another handlers.
    */
    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/create_object", app.createObject)
    mux.HandleFunc("/user", app.objectQuery)

    // creating fileserver which deals 
    fileServer := http.FileServer(http.Dir(cfg.StaticDir))
    // set mux to handle with /static prefix
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    // Start listen on :4000
    // addr is a pointer to the network address value. We need to dereference
    // it using * operator
    infoLogger.Printf("Start listening on port %s\n", cfg.Addr)

    /* Here we can implements a http.server struct */
    server := &http.Server {
        Addr:       cfg.Addr,
        ErrorLog:   errorLogger,
        Handler:    mux,
    }

    err = server.ListenAndServe()

    // call os.exit(1) immediately after writing the message
    errorLogger.Fatal(err)
}
