package main

import "net/http"

func handlerReadniess(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, struct{}{})
}
