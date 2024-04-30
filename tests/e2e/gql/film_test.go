//go:build e2e

package gql

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/goccy/go-json"
)

func (s *GqlTestSuite) TestCreateFilm() {
	inputReq := `{
		  "firstname": "Jon",
		  "lastname": "doa",
		  "password": "12345678",
		  "username": "usermovie"
		}`

	req, err := http.NewRequest("POST", "/api/v1/users/register", strings.NewReader(inputReq))

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	s.api.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
	var data userResponseData
	err = json.Unmarshal(w.Body.Bytes(), &data)
	if err != nil {
		s.T().Fatal(err)
	}
	s.Equal(data.Data.Username, "usermovie")

	filmInput := `{
		  "cast": [
			"string"
		  ],
		  "director": "string",
		  "genre": "string",
		  "releaseDate": "2021-02-18T21:54:42.123Z",
		  "synopsis": "string",
		  "title": "Avatar"
		}`
	req, err = http.NewRequest("POST", "/api/v1/films", strings.NewReader(filmInput))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+data.Data.JwtToken)

	w = httptest.NewRecorder()
	s.api.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
	var data2 filmResponseData
	err = json.Unmarshal(w.Body.Bytes(), &data2)
	if err != nil {
		s.T().Fatal(err)
	}
	s.Equal(data2.Data.Title, "Avatar")
}
