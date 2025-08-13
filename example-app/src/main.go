package main

import (
	"flag"
	"io"
	"log/slog"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from Go\n")
}

func main() {
	var (
		address string
	)
	flag.StringVar(&address, `a`, `:8080`, `the TCP address for the server to listen on`)
	flag.Parse()

	slog.Info(`starting web server`, `address`, address)
	http.HandleFunc("/", handler)
	http.ListenAndServe(address, nil)
}
