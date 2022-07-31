package utils

import (
	"net/http"
)

func ParseToJson(w http.ResponseWriter, json []byte){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}