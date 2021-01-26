package main

import (
	"net/http"

	"get-avataaar/api"
)

func main() {
	http.HandleFunc("/", api.Handler)

	http.ListenAndServe(":1337", nil)
}
