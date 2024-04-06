package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

var (
	port uint64
)

type Message struct {
	Message string `json:"message"`
}

func ping(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Message{"pong"})
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Uint64Var(&port, "port", 0, "A port number in the range 1-65535, both inclusive")
	flag.Parse()
	if port < 1 || port > 65535 {
		s := fmt.Sprintf("Expected input port number %d to be within the closed range 1-65535", port)
		panic(s)
	}
	http.HandleFunc("/ping", ping)
	addr := fmt.Sprintf(":%d", port)
	e := http.ListenAndServe(addr, nil)
	check(e)
}
