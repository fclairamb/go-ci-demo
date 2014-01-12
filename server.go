package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

const DEFAULT_PORT = 5000

func main() {
	port := flag.Int("port", DEFAULT_PORT, "Port to use")
	flag.Parse()

	log.Printf("Server: Version=%s/%d", BUILD_GIT_COMMIT, BUILD_NUMBER)

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		out := fmt.Sprint("This is build ", BUILD_NUMBER, " from commit ", BUILD_GIT_COMMIT)
		log.Println(out)
		w.Write([]byte(out))
	})

	http.HandleFunc("/cgo", func(w http.ResponseWriter, r *http.Request) {
		out := fmt.Sprint("Number of cgo calls: ", runtime.NumGoroutine())
		log.Println(out)
		w.Write([]byte(out))
	})

	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.OpenFile(LOGS_FILE, os.O_RDONLY, 0666)
		if err != nil {
			w.WriteHeader(404)
			log.Println("File could not be opened !")
			return
		}
		defer f.Close()
		reader := bufio.NewReader(f)
		copied, err := io.Copy(w, reader)
		log.Println("Copied", copied, "bytes of logs.")
	})

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "There's no favicon.")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		out := fmt.Sprintf("Hi there, I love \"%s\" !", strings.Replace(r.URL.Path[1:], "/", "<SLASH>", -1))
		log.Println(out)
		w.Write([]byte(out))
	})

	log.Printf("Listening on port %d...", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
