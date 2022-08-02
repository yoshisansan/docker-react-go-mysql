package api

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, dictionary map[string]interface{}) bool {
	data, err := json.Marshal(dictionary)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return true
}