package origin

import (
	"net/http"
)

func GetOrigin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte("{\"result\":\"OK\"}"))
}