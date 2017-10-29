package store

import (
	"gopkg.in/mgo.v2"
	"net/url"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func CreateStore() *mgo.Session {
	session, err := mgo.Dial(os.Getenv("MONGODB_URL"))
	if err != nil {
		panic(err)
	}
	return session
}

type UrlKey struct {
	Id string		`json:"id" bson:"_id"`
	Key uint64		`json:"key" bson:"key"`
}

type KeyedUrl struct {
	Id uint64		`json:"id" bson:"_id"`
	Url url.URL		`json:"url" bson:"url"`
}

func GetUpdatedUrlKey(s *mgo.Session) (uint64, error) {
	session := s.Copy()
	defer session.Close()

	var currentUrlKey UrlKey
	c := session.DB(os.Getenv("DATABASE_NAME")).C(os.Getenv("ID_COLLECTION_NAME"))
	change := mgo.Change{
		Update: bson.M{"$inc": bson.M{"key": 1}},
		Upsert: true,
		ReturnNew: true,
	}
	_, err := c.Find(bson.M{"_id": "id"}).Apply(change, &currentUrlKey)
	return currentUrlKey.Key, err
}

func InsertUrl(s *mgo.Session, id uint64, u url.URL) error {
	session := s.Copy()
	defer session.Close()

	c := session.DB(os.Getenv("DATABASE_NAME")).C(os.Getenv("URLS_COLLECTION_NAME"))

	e := c.Insert(KeyedUrl{Id: id, Url: u})

	return e
}

func GetUrl(s *mgo.Session, id uint64) (url.URL, error) {
	var ku KeyedUrl

	session := s.Copy()
	defer session.Close()

	c := session.DB(os.Getenv("DATABASE_NAME")).C(os.Getenv("URLS_COLLECTION_NAME"))

	e := c.FindId(id).One(&ku)

	return ku.Url, e
}