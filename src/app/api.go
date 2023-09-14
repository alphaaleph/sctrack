package app

import (
	"encoding/json"
	"github.com/alphaaleph/sctrack/src/models"
	"io/ioutil"
	"net/http"
)

// WriteJSONResponse encodes the inconing object
func WriteJSONResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

// WriteErrorResponse adds a return code and message to an error response
func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	WriteJSONResponse(w, code, models.Error{
		Code:    code,
		Message: message,
	})
}

// ExtractBodyJSON extracts the incoming object data into a selected struct
func ExtractBodyJSON(r *http.Request, body interface{}) error {
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
