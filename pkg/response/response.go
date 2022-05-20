package response

import (
	"encoding/json"
	"net/http"

	"github.com/MakMoinee/go-mith/pkg/common"
)

type ErrorResponse struct {
	ErrorStatus  int    `json:"errorStatus"`
	ErrorMessage string `json:"errorMessage"`
}

// Success() - returns success response
func Success(w http.ResponseWriter, payload interface{}) {
	result, err := json.Marshal(payload)
	if err != nil {
		errorBuilder := ErrorResponse{}
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		Error(w, errorBuilder)
		return
	}
	w.Header().Set(common.ContentTypeKey, common.ContentTypeValue)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

// Error() - returns error response
func Error(w http.ResponseWriter, payload ErrorResponse) {
	result, _ := json.Marshal(payload)
	w.Header().Set(common.ContentTypeKey, common.ContentTypeValue)
	w.WriteHeader(payload.ErrorStatus)
	w.Write([]byte(result))
}
