package main

import (
	"net/http"
	"github.com/ilya-shikhaleev/simple-video-server/handlers"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
}
func main() {
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	}
	defer file.Close()

	serverUrl := ":8000"
	log.WithFields(log.Fields{
		"serverUrl": serverUrl,
	}).Info("starting the server")
	router := handlers.Router()
	log.Info("the server is ready to serve requests")
	log.Fatal(http.ListenAndServe(":8000", router))
}
