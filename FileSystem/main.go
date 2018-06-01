package main

import (
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
func main() {
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(":8085", nil); err != nil {
		panic(err)
	}
}

//Created by Zazingo --- June 1, 2018
