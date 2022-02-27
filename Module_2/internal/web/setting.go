package web

import (
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Setting(w http.ResponseWriter, r *http.Request) {
	// 1 - Receive the client request and write the header from the request to the response header
	for key, values := range r.Header {
		for _, value := range values {
			log.Printf("Header key: %s, Header value: %s \n", key, value)
			w.Header().Set(key, value)
		}
	}

	// 2 - Reads the VERSION configuration from the current system's environment variables and writes the response header
	os.Setenv("VERSION", "v1.0.0")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	// 3 - Server side records access logs including client IP, HTTP return code, output to the server side standard output
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}

	log.Printf("clientip is : %s, Response Code is %d", ip, 200)
}