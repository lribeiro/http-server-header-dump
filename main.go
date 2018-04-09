package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {	
	header := w.Header()
	csp := []string{"default-src: 'self'", "font-src: 'fonts.googleapis.com'", "frame-src: 'none'"}
	header := http.Header{
    		"Content-Type": {"text/html; charset=UTF-8"},
	}
	header.Set("Content-Security-Policy", strings.Join(csp, "; "))
	r.Write(w)
	fmt.Fprintf(w, "RemoteAddr: %s", r.RemoteAddr)
}

func main() {
	port := flag.Int("port", 8080, "Listen on port")
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Println("Starting http server on port:", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}

