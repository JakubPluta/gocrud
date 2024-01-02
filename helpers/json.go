package helpers

import (
	"encoding/json"
	"net/http"
)

func ReadRequest(r *http.Request, v interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	ErrorPanic(err)
}

func WriteResponse(w http.ResponseWriter, v interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(v)
	ErrorPanic(err)
}
