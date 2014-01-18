package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/codegangsta/martini"
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

	m := martini.Classic()

	log.Printf("Server: Version=%s/%d", BUILD_GIT_COMMIT, BUILD_NUMBER)

	m.Get("/version", func() string {
		return fmt.Sprint("This is build ", BUILD_NUMBER, " from commit ", BUILD_GIT_COMMIT)
	})

	m.Get("/go", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintln("Version:", runtime.Version())))
		w.Write([]byte(fmt.Sprintln("CGO calls:", runtime.NumCgoCall())))
		{
			var mem runtime.MemStats
			runtime.ReadMemStats(&mem)
			w.Write([]byte(fmt.Sprintln("MEM: Footprint: ", mem.Alloc, ", Mallocs: ", mem.Mallocs, ", Frees: ", mem.Frees, ", GCs: ", mem.NumGC, "Next GC in ", (mem.NextGC - mem.HeapAlloc), " bytes")))
		}
	})

	m.Get("/logs", func(w http.ResponseWriter, r *http.Request) {
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

	m.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "There's no favicon.")
	})

	m.Get("/", func(w http.ResponseWriter, r *http.Request) {
		out := fmt.Sprintf("Hi there, I love \"%s\" !", strings.Replace(r.URL.Path[1:], "/", "<SLASH>", -1))
		log.Println(out)
		w.Write([]byte(out))
	})

	log.Printf("Listening on port %d...", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), m)
	if err != nil {
		log.Fatal(err)
	}
}
