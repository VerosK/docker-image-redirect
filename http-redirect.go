package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var redirect_target string
var redirect_code int = 302

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s\n", r.URL.Path)
	http.Redirect(w, r, redirect_target, redirect_code)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm alive")
}

func main() {
	var has_target bool
	redirect_target, has_target = os.LookupEnv("REDIRECT_TARGET")
	if !has_target {
		panic("Variable REDIRECT_TARGET is not defined")
	}
	_, has_permanent := os.LookupEnv("REDIRECT_PERMANENT")
	if has_permanent {
		redirect_code = http.StatusMovedPermanently
	} else {
		redirect_code = http.StatusFound
	}

	http.HandleFunc("/", handle)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
