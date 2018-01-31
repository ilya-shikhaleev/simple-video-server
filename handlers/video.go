package handlers

import (
	"net/http"
	"encoding/json"
	"io"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
)

// video is a HTTP handler function which writes a response with video information.
func video(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	log.WithField("id", id).Info("parse id")
	if id != "d290f1ee-6c54-4b01-90e6-d701748f0851" {
		http.NotFound(w, r)
		return
	}

	response := VideoItem{}
	response.ID = "d290f1ee-6c54-4b01-90e6-d701748f0851"
	response.Name = "Black Retrospetive Woman"
	response.Duration = 127
	response.Thumbnail = "/some/image.png"
	response.URL = "/some/video.mp4"

	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.WithField("err", err).Error("unmarshal error")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(b)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}
