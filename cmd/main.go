package main

import (
	"log"
	"net/http"

  "github.com/matsubara0507/sample-go-and-elm/front"
)

func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf8")
  w.Write(front.EmbedIndexHtml())
}
