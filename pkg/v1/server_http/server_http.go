package server_http

import (
	"fmt"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Accepted connection from: %s Proto: %s Method: %s URL: %s", r.RemoteAddr, r.Proto, r.Method, r.URL)
		log.Println(msg)
		w.WriteHeader(200)
		w.Write([]byte(msg))
	})
	log.Println("HTTP server started on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
