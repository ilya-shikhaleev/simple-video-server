package handlers

import (
	"net/http"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
)

// list is a HTTP handler function which writes a response with list of videos.
func list(w http.ResponseWriter, _ *http.Request) {
	var response []VideoListItem
	record := VideoListItem{}
	record.ID = "d290f1ee-6c54-4b01-90e6-d701748f0851"
	record.Name = "Black Retrospetive Woman"
	record.Duration = 127
	record.Thumbnail = "/some/image.png"
	response = append(response, record)

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
