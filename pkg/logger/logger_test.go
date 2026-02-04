package logger

import (
	"testing"
)

func TestInit(t *testing.T) {
	Init("info")

	if Log == nil {
		t.Fatal("Expected logger to be initialized")
	}
}

func TestGetLogger(t *testing.T) {
	logger := GetLogger()

	if logger == nil {
		t.Fatal("Expected non-nil logger")
	}
}

func TestLogLevels(t *testing.T) {
	testCases := []string{"debug", "info", "warn", "error"}

	for _, level := range testCases {
		Init(level)
		if Log == nil {
			t.Errorf("Logger should be initialized for level: %s", level)
		}
	}
}
