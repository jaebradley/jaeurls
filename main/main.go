package main

import (
	"github.com/jaebradley/jaeurls/config"
)

func main()  {
	session := config.CreateStore()
	defer session.Close()

	config.StartRouter(session)
}
