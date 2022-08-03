package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func createObject(response_writer http.ResponseWriter, request *http.Request) {

    /*
        Handle with obj request using POST.
    */

    if request.Method != "POST" {
        response_writer.Header().Set("Allow", "POST")
        http.Error(response_writer, "Method not Allowed.", 405)
        return
    }

    response_writer.Write([]byte("Creating object"))
    // or
    // fmt.Fprintf(response_writer, "Creating object")
}

func objectQuery(resw http.ResponseWriter, req *http.Request) {
    /*
        Object Query by id in the url.
    */

    str_id := req.URL.Query().Get("id")
    id, err := strconv.Atoi(str_id)

    if (id < 0 || err != nil) {
        http.NotFound(resw, req)
        return
    }

    str := "UserID: " + strconv.Itoa(id)
    //resw.Write([]byte(str))
    fmt.Fprintf(resw, str)
}

func home(response_writer http.ResponseWriter, request *http.Request) {
    if request.URL.Path != "/" {
        http.NotFound(response_writer, request)
        return
    }

    files := []string {
        "./ui/html/home.page.tmpl",
        "./ui/html/base.layout.tmpl",
        "./ui/html/footer.partial.tmpl",
    }

    template_set, err := template.ParseFiles(files...)

    if err != nil {
        log.Println(err.Error())
        http.Error(response_writer, "Internal Error.", 500)
    }

    err = template_set.Execute(response_writer, nil)

    if err != nil {
        log.Println(err.Error())
        http.Error(response_writer, "Internal Error.", 500)
    }
}
