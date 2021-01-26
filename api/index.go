package api

import (
	"fmt"
	"net/http"
)

// Handler - main handler function
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World!</h1>")
}
