package main

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
)

func getURL() {
	m.Get("/url/:id", func(params martini.Params) (int, string) {
		id := params["id"]
		if urls[id] == (URL{}) {
			error := ERRORMESSAGE{}
			error.MESSAGE = "id not found"
			message, err := json.Marshal(error)
			if err == nil {
				return http.StatusNotFound, string(message)
			}
		}
		res, err := json.Marshal(urls[id])
		if err != nil {
			return http.StatusBadRequest, "error"
		}
		return http.StatusFound, string(res)
	})
}

func addURL() {
	m.Post("/url/:id", binding.Json(URL{}), func(url URL, params martini.Params) (int, string) {
		par := params["id"]
		res := string(par)
		if url != (URL{}) {
			fillUrls(res, url.FULLURL)
		}
		return http.StatusCreated, "created successfully"
	})
}
