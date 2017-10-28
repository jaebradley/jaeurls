package config

import (
	"net/url"
	"strings"
	"errors"
)

func CreateUrl(key string) url.URL {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = "localhost"
	u.Path = "api/v1/jae" + key
	return u
}

func ParseUrl(u url.URL) (string, error) {
	var key string
	if isValidUrl(u) {
		i := strings.Index(u.Path, "/api/v1/jae")
		if i < 0 {
			return key, errors.New("Unable to identify key")
		}
		return u.Path[i + len("/api/v1/jae"):len(u.Path)], nil
	}

	return key, errors.New("Invalid URL")
}

func isValidUrl(u url.URL) bool {
	return strings.HasPrefix(u.Path, "/api/v1/jae")
}
