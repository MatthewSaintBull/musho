package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
)

var m = martini.Classic()

//ERRORMESSAGE response struct
type ERRORMESSAGE struct {
	MESSAGE string
}

var urls = make(map[string]URL)

func fillUrls(id string, url string) {
	urls[id] = URL{FULLURL: url}
}

func main() {
	getURL()
	addURL()
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Origin"},
	}))
	m.Run()
}
