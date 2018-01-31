package handlers

type VideoListItem struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
}

type VideoItem struct {
	VideoListItem
	URL string `json:"url"`
}
