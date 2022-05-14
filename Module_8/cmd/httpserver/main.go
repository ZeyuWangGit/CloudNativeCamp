package main

import (
	"Module_2/internal/web"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", web.Setting)
	mux.HandleFunc("/healthz", web.Healthz)

	log.Info("Go server Listening on: 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
