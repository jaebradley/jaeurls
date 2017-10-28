package config

import "net/url"

func CreateUrl(key string) url.URL {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = "localhost"
	u.Path = "api/v1/jae" + key
	return u
}
