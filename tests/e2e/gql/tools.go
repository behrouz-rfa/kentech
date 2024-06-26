//go:build e2e

package gql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/behrouz-rfa/kentech/internal/core/model"
)

type gqlQuery = map[string][]string

func query(s string) (*http.Request, error) {
	req, err := http.NewRequest("POST", "/gql", strings.NewReader(
		fmt.Sprintf(`{
		"query": "query { %s }"
	}`, s),
	))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "valid_token")

	return req, err
}

type responseData struct {
	Data   interface{} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type userResponseData struct {
	Data model.User `json:"data"`
}

type filmResponseData struct {
	Data model.Film `json:"data"`
}

type requestBody struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func mutationReq(reqBodyBytes []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", "/gql", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "valid_token")
	return req, nil
}

func createData(query string, vars map[string]interface{}) []byte {
	requestBody := &requestBody{
		Query:     query,
		Variables: vars,
	}
	reqBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	return reqBodyBytes
}
