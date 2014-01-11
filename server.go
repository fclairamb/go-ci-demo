package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const DEFAULT_PORT = 5000

func main() {
	port := flag.Int("port", DEFAULT_PORT, "Port to use")
	flag.Parse()

	log.Printf("Server: Version=%s/%d", BUILD_GIT_COMMIT, BUILD_VERSION)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, I love %s!", strings.Replace(r.URL.Path[1:], "/", "<SLASH>", -1))
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is build", BUILD_VERSION, "from commit", BUILD_GIT_COMMIT)
	})

	log.Printf("Listening on port %d...", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
