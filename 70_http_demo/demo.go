package main

import (
	"log"
	"net/http"
)

type server int

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL.Path)
	w.Write([]byte("Hello World"))
}

func main() {

	var s server

	http.ListenAndServe("localhost:9999", &s)

}
