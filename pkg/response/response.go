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
type SuccessResponse struct {
	Message string `json:"message"`
}

type ValidResponse interface {
	ErrorResponse | interface{}
	IsError() bool
}

func (svc *ErrorResponse) IsError() bool {
	if len(svc.ErrorMessage) > 0 {
		return true
	}
	return true
}

func (svc *SuccessResponse) IsError() bool {
	return false
}

func NewErrorBuilder(msg string, status int) ValidResponse {
	svc := ErrorResponse{}
	svc.ErrorMessage = msg
	svc.ErrorStatus = http.StatusInternalServerError
	return &svc
}

func NewSuccessBuilder(msg string) ValidResponse {
	svc := SuccessResponse{}
	svc.Message = msg
	return &svc
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

// Response
func Response[T ValidResponse](w http.ResponseWriter, payload T) {
	result, err := json.Marshal(payload)
	if err != nil {
		errorBuilder := ErrorResponse{}
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		Error(w, errorBuilder)
		return
	}

	if payload.IsError() {
		w.Header().Set(common.ContentTypeKey, common.ContentTypeValue)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(result))
		return
	}
	w.Header().Set(common.ContentTypeKey, common.ContentTypeValue)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
