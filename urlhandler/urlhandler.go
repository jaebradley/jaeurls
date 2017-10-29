package urlhandler

import (
	"net/url"
	"strings"
	"errors"
	"os"
)

func CreateUrl(key string) url.URL {
	u := url.URL{}
	u.Scheme = os.Getenv("SCHEME")
	u.Host = os.Getenv("HOST")
	u.Path = os.Getenv("PREFIX") + key
	return u
}

func ParseUrl(u url.URL) (string, error) {
	var key string
	var path = "/" + os.Getenv("PREFIX")
	if isValidUrl(u) {
		i := strings.Index(u.Path, path)
		if i < 0 {
			return key, errors.New("unable to identify key")
		}
		return u.Path[i + len(path):len(u.Path)], nil
	}

	return key, errors.New("invalid URL")
}

func isValidUrl(u url.URL) bool {
	var path = "/" + os.Getenv("PREFIX")
	return strings.HasPrefix(u.Path, path)
}
