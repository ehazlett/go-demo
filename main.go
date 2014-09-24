package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	p := os.Getenv("PROVIDER")
	if p == "" {
		hostname, err := os.Hostname()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		p = hostname
	}
	fmt.Fprintf(w, "Go demo running on %s testing from auto-deploy", p)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)

}
