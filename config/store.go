package config

import (
	"gopkg.in/mgo.v2"
	"log"
	"net/url"
	"gopkg.in/mgo.v2/bson"
)

func CreateStore() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

type urlKey struct {
	Id string		`json:"id" bson:"_id"`
	Key uint64		`json:"key" bson:"key"`
}

type keyedUrl struct {
	Id uint64		`json:"id" bson:"_id"`
	Url url.URL		`json:"url" bson:"url"`
}

func getUpdatedUrlKey(s *mgo.Session) (uint64, error) {
	session := s.Copy()
	defer session.Close()

	var currentUrlKey urlKey
	c := session.DB("test").C("id")
	change := mgo.Change{
		Update: bson.M{"$inc": bson.M{"key": 1}},
		Upsert: true,
		ReturnNew: true,
	}
	_, err := c.Find(bson.M{"_id": "id"}).Apply(change, &currentUrlKey)
	return currentUrlKey.Key, err
}

func insertUrl(s *mgo.Session, id uint64, u url.URL) error {
	session := s.Copy()
	defer session.Close()

	c := session.DB("test").C("urls")

	e := c.Insert(keyedUrl{Id: id, Url: u})

	if e != nil {
		log.Println("Failed to insert url", e)
	}

	return e
}

func getUrl(s *mgo.Session, id uint64) (url.URL, error) {
	var ku keyedUrl

	session := s.Copy()
	defer session.Close()

	c := session.DB("test").C("urls")

	e := c.FindId(id).One(&ku)

	if e != nil {
		log.Println("Failed to get url", e)
	}

	return ku.Url, e
}