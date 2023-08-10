package url_controller

type CreateShortUrl struct {
	Url string `json:"url"`
}

type CreateShortUrls struct {
	Urls []string `json:"urls"`
}
