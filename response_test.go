package gogenz

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestJSONResponse tests the JSONResponse function
func TestJSONResponse(t *testing.T) {
	// Create a test response recorder
	rr := httptest.NewRecorder()

	// Define a dummy data
	data := map[string]string{"key": "value"}

	// Call the JSONResponse function
	JSONResponse(rr, http.StatusOK, "Request was successful", data)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expected := `{"status":"success","message":"Request was successful","data":{"key":"value"}}`
	if rr.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, rr.Body.String())
	}
}

// TestErrorResponse tests the ErrorResponse function
func TestErrorResponse(t *testing.T) {
	// Create a test response recorder
	rr := httptest.NewRecorder()

	// Call the ErrorResponse function
	ErrorResponse(rr, http.StatusBadRequest, "Bad Request")

	// Check the status code
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}

	// Check the response body
	expected := `{"status":"error","message":"Bad Request"}`
	if rr.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, rr.Body.String())
	}
}

// TestParseJSONBody tests the ParseJSONBody function
func TestParseJSONBody(t *testing.T) {
	// Create a valid JSON body
	validBody := `{"name": "John"}`
	req, err := http.NewRequest("POST", "/test", bytes.NewBufferString(validBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a test response recorder
	rr := httptest.NewRecorder()

	var result map[string]string
	if !ParseJSONBody(rr, req, &result) {
		t.Errorf("Expected JSON parsing to succeed, but it failed")
	}

	// Check if the body was correctly parsed
	if result["name"] != "John" {
		t.Errorf("Expected name to be John, got %s", result["name"])
	}

	// Create an invalid JSON body
	invalidBody := `{"name": "John"`
	req, err = http.NewRequest("POST", "/test", bytes.NewBufferString(invalidBody))
	if err != nil {
		t.Fatal(err)
	}

	// Call ParseJSONBody again to test invalid input
	if ParseJSONBody(rr, req, &result) {
		t.Errorf("Expected JSON parsing to fail, but it succeeded")
	}
}

// TestSuccessResponse tests the SuccessResponse function
func TestSuccessResponse(t *testing.T) {
	// Create a test response recorder
	rr := httptest.NewRecorder()

	// Call the SuccessResponse function
	SuccessResponse(rr, http.StatusOK, "Success without data")

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expected := `{"status":"success","message":"Success without data","data":null}`
	if rr.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, rr.Body.String())
	}
}

// TestCreatedResponse tests the CreatedResponse function
func TestCreatedResponse(t *testing.T) {
	// Create a test response recorder
	rr := httptest.NewRecorder()

	// Define dummy data
	data := map[string]string{"key": "value"}

	// Call the CreatedResponse function
	CreatedResponse(rr, "Resource created", data)

	// Check the status code
	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, rr.Code)
	}

	// Check the response body
	expected := `{"status":"success","message":"Resource created","data":{"key":"value"}}`
	if rr.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, rr.Body.String())
	}
}
