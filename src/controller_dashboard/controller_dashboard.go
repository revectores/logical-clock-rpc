package main

import (
	"fmt"
	"net/http"
	"log"
)

func hello(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "controller_dashboard.html")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.Handle("/style.css", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", headers)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
    	log.Fatal("ListenAndServe: ", err)
	}
}