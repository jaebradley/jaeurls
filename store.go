package jaeurls

import (
	"gopkg.in/mgo.v2"
	"log"
	"net/url"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
}

type urlKey struct {
	id string
	key uint64
}

type keyedUrl struct {
	id uint64
	url url.URL
}

func getUpdatedUrlKey(s *mgo.Session) uint64 {
	session := s.Copy()
	defer session.Close()

	currentUrlKey := urlKey{}

	c := session.DB("test").C("id")
	change := mgo.Change{
		Update: bson.M{"$inc": bson.M{"key": 1}},
		Upsert: true,
		ReturnNew: true,
	}
	_, err := c.Find(bson.M{"_id": "id"}).Apply(change, &currentUrlKey)
	if err != nil {
		log.Println("Failed to increment counter", err)
	} else {
		return currentUrlKey.key
	}
}

func insertUrl(s *mgo.Session, id uint64, url url.URL) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("test").C("urls")

	e := c.Insert(keyedUrl{id: id, url: url})

	if e != nil {
		log.Print("Failed to insert url", e)
	} else {
		log.Print("Succeeded to insert url")
	}
}

func InsertUrl(s *mgo.Session, url url.URL) {
	var key uint64 = getUpdatedUrlKey(s)
	insertUrl(s, key, url)
}
