package main

import (
	"context"
	"net/http"
	"github.com/ilya-shikhaleev/simple-video-server/handlers"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
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
	killSignalChan := getKillSignalChan()
	srv := startServer(serverUrl)

	waitForKillSignal(killSignalChan)
	log.Info("the service is shutting down...")
	srv.Shutdown(context.Background())
}

func startServer(serverUrl string) *http.Server {
	log.WithFields(log.Fields{"serverUrl": serverUrl}).Info("starting the server")

	router := handlers.Router()
	srv := &http.Server{
		Addr:    serverUrl,
		Handler: router,
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	log.Info("the server is ready to serve requests")

	return srv
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Kill, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}
