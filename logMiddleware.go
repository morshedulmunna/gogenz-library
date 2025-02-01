package gogenz

import (
	"log"
	"net/http"
	"os"
	"time"
)

// LoggingMiddleware logs each incoming request with method, URL, and response time
func LoggingMiddleware(next http.Handler) http.Handler {
	logger := log.New(os.Stdout, "HTTP: ", log.LstdFlags)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Capture the start time
		start := time.Now()

		// Log request details
		logger.Printf("Request received: %s %s", r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Calculate and log the response time
		duration := time.Since(start)
		logger.Printf("Request processed: %s %s in %v", r.Method, r.URL.Path, duration)
	})
}

func SetupLogging() *log.Logger {
	// Create 'logs' directory if it doesn't exist
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		log.Fatal("Error creating logs directory: ", err)
	}

	// Generate the log file name based on the current date
	logFileName := "logs/" + time.Now().Format("2006-01-02") + ".log"

	// Create/Open the log file in append mode
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}

	// Set up the logger to write to the log file
	logger := log.New(file, "", log.LstdFlags)
	log.SetOutput(file) // Log to the file by default

	return logger
}
