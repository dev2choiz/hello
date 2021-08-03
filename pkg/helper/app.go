package helper

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, cod int, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(cod)
	_ = json.NewEncoder(w).Encode(data)
}

func OkHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	Json(w, http.StatusOK, map[string]interface{}{"status": "ok"})
}
