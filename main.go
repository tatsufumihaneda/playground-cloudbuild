package main

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", HealthcheckHandler)
}

func main() {
	port := "8080"
	http.ListenAndServe(":"+port, nil)
}

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK -CI BUILD-"))
}
