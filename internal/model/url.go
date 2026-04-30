package model

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"` 
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}