package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//HTTPServer : manages the state of the HTTP Server
type HTTPServer struct {
	tracker *Tracker
}

func startHTTPServer(t *Tracker, port int) {
	server := HTTPServer{t}

	r := mux.NewRouter()

	r.HandleFunc("/status", server.statusHandler).Methods("GET")

	r.HandleFunc("/job/{id}", server.jobStatusHandler).Methods("GET")
	r.HandleFunc("/job", server.jobAddHandler).Methods("POST")

	http.Handle("/", r)

	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}

func (server *HTTPServer) jobStatusHandler(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK

	// vars := mux.Vars(r)
	// jobID := vars["id"]
	w.Write([]byte("Endpoint is WIP"))

	w.WriteHeader(status)
}

func (server *HTTPServer) jobAddHandler(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK

	// server.tracker.AddJob()
	w.Write([]byte("Endpoint is WIP"))

	w.WriteHeader(status)
}

func (server *HTTPServer) statusHandler(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK

	w.Write([]byte("Endpoint is WIP"))

	w.WriteHeader(status)
}
