package main

import (
	"fmt"
	"net/http"
	"time"
)

var Gc int

func main() {
	Gc = 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Gc = Gc + 1
		time.Sleep(1 * time.Second)
		fmt.Fprintf(w, "Welcome!You are %d \n", Gc)
		Gc = Gc - 1
	})
	http.ListenAndServe(":8080", nil)
}
