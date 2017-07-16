package main


import(
    "fmt"
    "io"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

var (
    server  *httptest.Server
    reader  io.Reader
    url     string
)

func init(){
    server = httptest.NewServer(main.Handlers()
}
