package log

import (
	"encoding/json"
	"net/http"
)

type LogCollector struct{}

func NewLogCollector() *LogCollector {
	return &LogCollector{}
}

func (l *LogCollector) GetLogLines(w http.ResponseWriter, r *http.Request) {
	// Load log lines from file
	logLines := []byte(`2023-02-20T14:30:00.000Z INFO my-app my-container started
2023-02-20T14:30:01.000Z INFO my-app my-container running
2023-02-20T14:30:02.000Z ERROR my-app my-container error occurred`)

	// Return log lines as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logLines)
}