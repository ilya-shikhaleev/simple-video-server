package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"encoding/json"
)

func TestList(t *testing.T) {
	w := httptest.NewRecorder()
	list(w, nil)
	response := w.Result()

	listUrl := "/api/v1/list"
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code for %s is wrong. Have: %d, want: %d.", listUrl, response.StatusCode, http.StatusOK)
	}

	jsonContentTypeHeader := "application/json; charset=UTF-8"
	if response.Header.Get("Content-type") != jsonContentTypeHeader {
		t.Errorf("Status code for %s is wrong. Have: %d, want: %d.", listUrl, response.StatusCode, http.StatusOK)
	}

	jsonString, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	items := make([]VideoListItem, 10)
	err = json.Unmarshal(jsonString, &items)
	if err != nil {
		t.Errorf("Can't parse json response with error %v", err)
	}
}
