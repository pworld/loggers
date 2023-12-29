package loggers

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

// testWriter is a helper type to capture log output
type testWriter struct {
	buf bytes.Buffer
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	return tw.buf.Write(p)
}

func (tw *testWriter) String() string {
	return tw.buf.String()
}

// createTestLogger creates a Logger with a testWriter to capture output
func createTestLogger() (*Logger, *testWriter) {
	tw := new(testWriter)
	logger := &Logger{
		StandardLogger: log.New(tw, "", 0),
	}
	return logger, tw
}

// TestInfoLog tests the Info method of Logger
func TestInfoLog(t *testing.T) {
	logger, tw := createTestLogger()

	testMessage := "Test info message"
	logger.Info(testMessage)

	if !strings.Contains(tw.String(), testMessage) {
		t.Errorf("Expected %q to be in log output", testMessage)
	}
}

// TestErrorLog tests the Error method of Logger
func TestErrorLog(t *testing.T) {
	logger, tw := createTestLogger()

	testMessage := "Test error message"
	logger.Error(testMessage)

	if !strings.Contains(tw.String(), testMessage) {
		t.Errorf("Expected %q to be in log output", testMessage)
	}
}

// TestDebugLog tests the Debug method of Logger
func TestDebugLog(t *testing.T) {
	logger, tw := createTestLogger()

	testMessage := "Test debug message"
	logger.Debug(testMessage)

	if !strings.Contains(tw.String(), testMessage) {
		t.Errorf("Expected %q to be in log output", testMessage)
	}
}

// TestWarnLog tests the Warn method of Logger
func TestWarnLog(t *testing.T) {
	logger, tw := createTestLogger()

	testMessage := "Test warn message"
	logger.Warn(testMessage)

	if !strings.Contains(tw.String(), testMessage) {
		t.Errorf("Expected %q to be in log output", testMessage)
	}
}

// TestFatalLog tests the Fatal method of Logger
// Note: Testing Fatal can be tricky since it exits the program.
// You might want to refactor Logger to make it more testable for Fatal logs.
//func TestFatalLog(t *testing.T) {
//	logger, tw := createTestLogger()
//
//	testMessage := "Test fatal message"
//	// Normally, logger.Fatal would exit the program. You might need to adjust your logger to handle this in tests.
//	logger.Fatal(testMessage)
//
//	if !strings.Contains(tw.String(), testMessage) {
//		t.Errorf("Expected %q to be in log output", testMessage)
//	}
//}
