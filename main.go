package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "Go demo running on %s", hostname)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
