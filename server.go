package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const DEFAULT_PORT = 5000

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", strings.Replace(r.URL.Path[1:], "/", "<SLASH>", -1))
}

func main() {
	port := flag.Int("port", DEFAULT_PORT, "Port to use")
	flag.Parse()

	http.HandleFunc("/", handler)

	log.Printf("Listening on port %d...", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
