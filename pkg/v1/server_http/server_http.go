package server_http

import (
	"fmt"
	"log"
	"net/http"
)

func Run(branch, commithash, version string) {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("App version: %s, accepted connection from: %s Proto: %s Method: %s URL: %s", version, r.RemoteAddr, r.Proto, r.Method, r.URL)
		log.Println(msg)
		w.WriteHeader(200)
		w.Write([]byte(msg))
	})
	log.Printf("Branch: %s, CommitHash: %s, Version: %s \n", branch, commithash, version)
	log.Println("HTTP server started on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
