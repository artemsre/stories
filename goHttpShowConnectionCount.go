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
		fmt.Fprintf(w, "Welcome!You are %d \n", Gc)
		fmt.Printf("We have simultaneous connections: %d \n", Gc)
		time.Sleep(1 * time.Second)
		Gc = Gc - 1
	})
	http.ListenAndServe(":8080", nil)
}
