//go:build e2e

package gql

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/goccy/go-json"
)

func (s *GqlTestSuite) TestCreateUser() {
	inputReq := `{
  "firstname": "Jon",
  "lastname": "doa",
  "password": "12345678",
  "username": "jondoa"
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
	s.Equal(data.Data.Username, "jondoa")
}

func (s *GqlTestSuite) TestUpdateUser() {
	inputReq := `{
		  "firstname": "Jon",
		  "lastname": "doa",
		  "password": "12345678",
		  "username": "jondoa2"
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
	s.Equal(data.Data.Username, "jondoa2")

	inputUpdateReq := `{
		  "firstname": "Jon3",
		  "lastname": "doa3",
		  "username": "jondoa3"
		}`

	req, err = http.NewRequest("PUT", fmt.Sprintf("/api/v1/users/%s", data.Data.ID), strings.NewReader(inputUpdateReq))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+data.Data.JwtToken)

	w = httptest.NewRecorder()
	s.api.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
	var data2 userResponseData
	err = json.Unmarshal(w.Body.Bytes(), &data2)
	if err != nil {
		s.T().Fatal(err)
	}
	s.Equal(data2.Data.Username, "jondoa3")
	s.Equal(data2.Data.Firstname, "Jon3")
}
