package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	Handler()
}

//GetUpdates main function to get updates from telegram and procced
func Handler() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Accepted connection from: %s Proto: %s Method: %s URL: %s", r.RemoteAddr, r.Proto, r.Method, r.URL)
		log.WithFields(log.Fields{
			"function":   "RootHandler",
			"RemoteAddr": r.RemoteAddr,
			"UserAgent":  r.UserAgent(),
			"Proto":      r.Proto,
			"Method":     r.Method,
			"URL":        r.URL.Path,
		}).Info()
		w.WriteHeader(200)
		w.Write([]byte(msg))
	})
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
