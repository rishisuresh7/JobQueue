package handler

import "net/http"

func NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("resource not found"))
	}
}
