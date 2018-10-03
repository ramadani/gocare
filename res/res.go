package res

import (
	"encoding/json"
	"net/http"
)

// ResponseData json wrapper
type ResponseData struct {
	Data interface{} `json:"data"`
}

// ResponseError json wrapper
type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Data wrapper
func Data(data interface{}) ResponseData {
	return ResponseData{data}
}

// Error wrapper
func Error(code, msg string) ResponseError {
	return ResponseError{code, msg}
}

// JSON response
func JSON(w http.ResponseWriter, data interface{}, statusCode int) {
	result, err := json.Marshal(data)
	if err != nil {
		Fail(w, "internal-server-error", "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(result)
}

// Fail response
func Fail(w http.ResponseWriter, code string, msg string, statusCode int) {
	result, err := json.Marshal(Data(Error(code, msg)))
	if err != nil {
		Fail(w, "internal-server-error", "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(result)
}
