package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/", index)
	mux.HandleFunc("/list", list)

	// http.Server構造体のポインタを宣言
	server := &http.Server{
		Addr:    ":11180",
		Handler: mux,
	}
	server.ListenAndServe()
}
