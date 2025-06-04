package main

import (
	"fmt"
	"net/http"
	"time"
)

func greetwrld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do something with your life. %s", time.Now())
}

func main() {
	http.HandleFunc("/", greetwrld)
	http.ListenAndServe(":8080", nil)
}
