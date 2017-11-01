# jaeurls

## Introduction

I've read a number of articles on [Go](https://golang.org/) ([for](https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65) [example](https://hackernoon.com/the-beauty-of-go-98057e3f0a7d)) but never used it. Similarly, I have never used [MongoDB](https://www.mongodb.com/) before, so when I stumbled across an article on how to create a URL shortener, I figured creating my own URL shortener using Go and Mongo.

### API
It's pretty simple - there are two API endpoints

#### Create a shortened URL
`POST` an `HTTP` request to the `api/v1/` endpoint with a URL in the request's body.
```bash
curl -X POST -d '{"url": "http://google.com"}' https://jaeurls.herokuapp.com/api/v1/
```
In the returned response payload, there will be a `Url` field that contains the shortened URL.

#### Redirect to your input URL
Make a `GET` request (using a browser, for example) with the shortened URL and (hopefully) you'll be redirected to your input URL!

### Implementation
I keep track of an ID for each input URL. I also use a set of characters as part of my hashing strategy. Given this set of characters and an ID for a given input URL, I hash the ID into some string of characters, which are now used to produce the shortened URL.

When I receive a `GET` request, I parse out the hash, translate the hash into an ID, lookup the ID in my persistence layer, get the URL associated with that ID, and then redirect to that URL.
