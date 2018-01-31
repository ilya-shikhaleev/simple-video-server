package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"encoding/json"
)

func TestVideo(t *testing.T) {
	videoListItem := getFirstListItem(t)
	r := Router()
	ts := httptest.NewServer(r)
	defer ts.Close()

	videoUrl := "/api/v1/video/" + videoListItem.ID
	response, err := http.Get(ts.URL + videoUrl)
	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code for %s is wrong. Have: %d, want: %d.", videoUrl, response.StatusCode, http.StatusOK)
	}

	jsonContentTypeHeader := "application/json; charset=UTF-8"
	if response.Header.Get("Content-type") != jsonContentTypeHeader {
		t.Errorf("Status code for %s is wrong. Have: %d, want: %d.", videoUrl, response.StatusCode, http.StatusOK)
	}

	jsonString, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	item := &VideoItem{}
	err = json.Unmarshal(jsonString, item)
	if err != nil {
		t.Errorf("Can't parse json response with error %v", err)
	}
}

func getFirstListItem(t *testing.T) *VideoListItem {
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
	if len(items) < 1 {
		t.Errorf("Empty response %s", jsonString)
	}

	return &items[0]
}