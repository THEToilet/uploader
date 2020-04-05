package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/list", list)
	mux.HandleFunc("/viewlist", viewlist)
	mux.HandleFunc("/showwiki", showwiki)

	// http.Server構造体のポインタを宣言
	server := &http.Server{
		Addr:           ":11180"/*config.Address*/,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
