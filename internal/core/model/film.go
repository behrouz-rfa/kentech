package model

import (
	"time"
)

// Film represents a film in the system
type Film struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
	Title       string    `json:"title"`
	Director    string    `json:"director"`
	ReleaseDate time.Time `json:"releaseDate"`
	Cast        []string  `json:"cast"`
	Genre       string    `json:"genre"`
	Synopsis    string    `json:"synopsis"`
	CreatorID   string    `json:"creatorID"`
}

type FilmInput struct {
	Title       string    `json:"title"`
	Director    string    `json:"director"`
	ReleaseDate time.Time `json:"releaseDate"`
	Cast        []string  `json:"cast"`
	Genre       string    `json:"genre"`
	Synopsis    string    `json:"synopsis"`
	CreatorID   string    `json:"creatorID"`
}

type FilmUpdateInput struct {
	Title       string    `json:"title"`
	Director    string    `json:"director"`
	ReleaseDate time.Time `json:"releaseDate"`
	Cast        []string  `json:"cast"`
	Genre       string    `json:"genre"`
	Synopsis    string    `json:"synopsis"`
	CreatorID   string    `json:"creatorID"`
}
