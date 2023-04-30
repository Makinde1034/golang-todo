package controller

import (
	"encoding/json"
	"get-going/structs"
	"net/http"
)
func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Name: "Bryan",
				Body: "Senior Software Engineer",
			}
			json.NewEncoder(w).Encode(data)
		}
	}
}

