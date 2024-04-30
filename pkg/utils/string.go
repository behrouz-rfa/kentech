package utils

import (
	"errors"
	"regexp"
)

func GetUUIDFromURL(imgURL string) (string, error) {
	// Parse URL
	// Regex to match UUID portion
	uuidRegex := regexp.MustCompile(`[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`)

	// Find UUID match
	match := uuidRegex.FindString(imgURL)

	if match == "" {
		return "", errors.New("couldnt find the image")
	}

	return match, nil
}
