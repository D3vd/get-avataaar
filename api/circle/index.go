package api

import (
	"get-avataaar/types"
	"get-avataaar/utils"
	"net/http"
)

// CircleHandler - circle handler function
func CircleHandler(w http.ResponseWriter, r *http.Request) {
	res := types.Response{
		URL: utils.GenerateURL("Circle"),
	}

	utils.MarshalAndRespond(w, res)
}
