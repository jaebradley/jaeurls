package eventhandler

import (
	"gopkg.in/mgo.v2"
	"net/url"
	"github.com/jaebradley/jaeurls/store"
	"github.com/jaebradley/jaeurls/keyhandler"
)

func InsertUrl(session *mgo.Session, url url.URL) (string, error) {
	var encodedKey string
	k, err := store.GetUpdatedUrlKey(session)

	if err != nil {
		panic(err)
	}

	e := store.InsertUrl(session, k, url)

	if e == nil {
		encodedKey = keyhandler.Encode(k)
	}

	return encodedKey, e
}

func IdentifyUrl(session *mgo.Session, k string) (url.URL, error) {
	return store.GetUrl(session, keyhandler.Decode(k))
}
