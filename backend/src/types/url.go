package types

import "time"

type CreateURLRequest struct {
	Url     string `json:"url"`
	Expires string `json:"expires"`
}

type EditURLRequest struct {
	Url     string `json:"url"`
	Expires string `json:"expires"`
}

type UserHistoryResponse struct {
	Original   string    `json:"original"`
	ShortenId  string    `json:"shorten_id"`
	RawExpires string    `json:"raw_expires"`
	CreatedAt  time.Time `json:"created_at"`
	Expires    time.Time `json:"expires"`
}
