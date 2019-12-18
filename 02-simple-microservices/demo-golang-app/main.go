package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Server starting at: http://localhost:8080")
	http.Handle("/", loggingMiddleware(http.HandlerFunc(handler)))
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "learn-kubernetes-demo-golang-app")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Infof("uri: %s", req.RequestURI)
		next.ServeHTTP(w, req)
	})
}
