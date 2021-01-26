package api

import (
	"get-avataaar/types"
	"get-avataaar/utils"
	"net/http"
)

// TransparentHandler - transparent handler function
func TransparentHandler(w http.ResponseWriter, r *http.Request) {
	res := types.Response{
		URL: utils.GenerateURL("Transparent"),
	}

	utils.MarshalAndRespond(w, res)
}
