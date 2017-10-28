package config

import (
	"gopkg.in/mgo.v2"
	"net/url"
)

func InsertUrl(session *mgo.Session, url url.URL) (string, error) {
	var encodedKey string
	key, err := getUpdatedUrlKey(session)

	if err != nil {
		panic(err)
	}

	e := insertUrl(session, key, url)

	if e == nil {
		encodedKey = Encode(key)
	}


	return encodedKey, e
}
