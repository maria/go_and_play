package main

import (
    "fmt"
    "net/http"
)

type String string

type Struct struct {
    
    Greeting string
    Punct string
    Who string
}


func (s String) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, s)
}

func main() {
    http.Handle("/string", String("I go through go like there's not tomorrow"))
    http.ListenAndServe("localhost:4000", nil)
}
