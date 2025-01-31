package gogenz

import (
	"encoding/json"
	"net/http"
)

// APIResponse represents the standard structure of an API response
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// JSONResponse s a standard JSON response
func JSONResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	response := APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Failed to  response")
	}
}

// ErrorResponse s a JSON error response
func ErrorResponse(w http.ResponseWriter, status int, message string) {
	response := APIResponse{
		Status:  "error",
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to  error response", http.StatusInternalServerError)
	}
}

// SuccessResponse s a standard success response without data
func SuccessResponse(w http.ResponseWriter, status int, message string) {
	JSONResponse(w, status, message, nil)
}

// CreatedResponse s a 201 Created response
func CreatedResponse(w http.ResponseWriter, message string, data interface{}) {
	JSONResponse(w, http.StatusCreated, message, data)
}

// BadRequestResponse s a 400 Bad Request response
func BadRequestResponse(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusBadRequest, message)
}

// UnauthorizedResponse s a 401 Unauthorized response
func UnauthorizedResponse(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusUnauthorized, message)
}

// NotFoundResponse s a 404 Not Found response
func NotFoundResponse(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusNotFound, message)
}

// InternalServerErrorResponse s a 500 Internal Server Error response
func InternalServerErrorResponse(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusInternalServerError, message)
}

// MethodNotAllowedResponse s a 405 Method Not Allowed response
func MethodNotAllowedResponse(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusMethodNotAllowed, message)
}

// ParseJSONBody parses the JSON body of the request and stores it in the given object
func ParseJSONBody(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	if err != nil {
		BadRequestResponse(w, "Invalid JSON")
		return false
	}
	return true
}
