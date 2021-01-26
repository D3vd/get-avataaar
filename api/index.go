package api

import (
	"get-avataaar/types"
	"get-avataaar/utils"
	"net/http"
)

// Handler - main handler function
func Handler(w http.ResponseWriter, r *http.Request) {
	res := types.Response{
		URL: utils.GenerateURL(),
	}

	utils.MarshalAndRespond(w, res)
}
