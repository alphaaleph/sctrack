package handlers

import (
	"encoding/json"
	"github.com/alphaaleph/sctrack/server/src/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

// writeJSONResponse encodes the incoming object
func writeJSONResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

// writeErrorResponse adds a return code and message to an error response
func writeErrorResponse(w http.ResponseWriter, code int, message string) {
	writeJSONResponse(w, code, models.DBError{
		Message: message,
	})
}

// extractBodyJSON extracts the incoming object data into a selected struct
func extractBodyJSON(r *http.Request, body interface{}) error {
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, &body)
	if err != nil {
		return err
	}

	return nil
}

// nUint converts string to uint
func nUint(value string) uint {
	v, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0
	}
	return uint(v)
}
