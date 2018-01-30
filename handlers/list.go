package handlers

import (
	"fmt"
	"net/http"
)

// list is a HTTP handler function which writes a response with list of videos.
func list(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(
		w,
		`[{
			"id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
			"name": "Black Retrospetive Woman",
			"duration": 127,
			"thumbnail": "/some/image.png"
		}]`,
	)
}
