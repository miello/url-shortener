package types

import "time"

type CreateURLRequest struct {
	Url string `json:"url"`
}

type EditURLRequest struct {
	Url string `json:"url"`
}

type UserHistoryResponse struct {
	Original  string    `json:"original"`
	ShortenId string    `json:"shorten_id"`
	CreatedAt time.Time `json:"created_at"`
	Expires   time.Time `json:"expires"`
}
