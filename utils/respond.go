package utils

import (
	"encoding/json"
	"get-avataaar/types"
	"net/http"
)

// MarshalAndRespond : marshals and responds with default response
func MarshalAndRespond(w http.ResponseWriter, res types.Response) {
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
