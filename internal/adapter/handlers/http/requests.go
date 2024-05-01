package http

import "time"

// listUsersRequest represents the request body for listing films
type listFilmsRequest struct {
	Page  uint64    `form:"page" binding:"required,min=0" example:"0"`
	Limit uint64    `form:"limit" binding:"required,min=5" example:"10"`
	Title *string   `form:"title" example:"Name of movie"`
	Genre *string   `form:"genre" example:"Name of genre"`
	From  *DataTime `form:"from" example:"2021-02-18"`
	To    *DataTime `form:"to" example:"2021-02-18"`
}

// updateFilmRequest represents the request body for updating a film
type updateFilmRequest struct {
	Title       *string   `json:"title"`
	Director    *string   `json:"director"`
	ReleaseDate *DataTime `json:"releaseDate" example:"2021-02-18"`
	Cast        []*string `json:"cast"`
	Genre       *string   `json:"genre"`
	Synopsis    *string   `json:"synopsis"`
	CreatorID   *uint64   `json:"creatorID"`
}

// createFilmRequest represents the request body for create a film
type createFilmRequest struct {
	Title       string   `json:"title"`
	Director    string   `json:"director"`
	ReleaseDate DataTime `json:"releaseDate"example:"2021-02-18"`
	Cast        []string `json:"cast"`
	Genre       string   `json:"genre"`
	Synopsis    string   `json:"synopsis"`
}

type DataTime time.Time

const DateFormat = "2006-01-02"

func (d *DataTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*d = DataTime(time.Time{})
		return
	}
	now, err := time.Parse(`"`+DateFormat+`"`, string(data))
	if err != nil {
		return err
	}
	*d = DataTime(now)
	return
}
