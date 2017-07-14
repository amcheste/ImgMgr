package main

import (
    "fmt"
    "log"
    "net/http"
)

type Album struct {
    Name string `json:"name"`
}

type Albums []Album


func main() {
    fmt.Println("Starting ImgMgr server...")

    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}

