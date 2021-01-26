package main

import (
	"net/http"

	"get-avataaar/api"
	circle "get-avataaar/api/circle"
	transparent "get-avataaar/api/transparent"
)

func main() {
	http.HandleFunc("/", api.Handler)
	http.HandleFunc("/circle", circle.CircleHandler)
	http.HandleFunc("/transparent", transparent.TransparentHandler)

	http.ListenAndServe(":1337", nil)
}
