package main

import (
	"net/http"
	"github.com/ilya-shikhaleev/simple-video-server/handlers"
	"fmt"
)

func main() {
	router := handlers.Router()
	fmt.Println(http.ListenAndServe(":8000", router))
}
