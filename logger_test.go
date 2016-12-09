package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// Formatter
type testLoggerFormat struct {
}

func (t *testLoggerFormat) Format(e *Entry) (out string, err error) {
	var buf bytes.Buffer
	buf.WriteString(e.Level + " " + e.Event + " " + e.Message)

	if len(e.Meta) > 0 {
		for k, v := range e.Meta {
			buf.WriteString(fmt.Sprintf(" %s=", k))
			buf.WriteString(fmt.Sprintf("%s", v))
		}
	}

	return buf.String(), nil
}

// Writer
type testLoggerWriter struct {
	lastLogs []string
}

func (t *testLoggerWriter) clearLogs() {
	t.lastLogs = make([]string, 0)
}

func (t *testLoggerWriter) Write(p []byte) (n int, err error) {
	t.lastLogs = append(t.lastLogs, string(p))

	return len(p), nil
}

// Transport
var testWrite = &testLoggerWriter{}
var testFormat = &testLoggerFormat{}
var testTransport = NewTransport(testWrite, testFormat)

func TestLoggerLog(t *testing.T) {
	var tests = []struct {
		Method  string
		Level   string
		Event   string
		Message string
	}{
		{"Debug", "Debug", "EventName", "Message..."},
		{"Info", "Info", "EventName", "Message..."},
		{"Warn", "Warning", "EventName", "Message..."},
		{"Error", "Error", "EventName", "Message..."},
	}

	log := New()
	log.AddTransport(testTransport)

	for _, item := range tests {
		t.Run(item.Method, func(t *testing.T) {
			var err error

			switch item.Method {
			case "Debug":
				err = log.Debug(item.Event, item.Message)
			case "Info":
				err = log.Info(item.Event, item.Message)
			case "Warn":
				err = log.Warn(item.Event, item.Message)
			case "Error":
				err = log.Error(item.Event, item.Message)
			default:
				err = log.Log(item.Level, item.Event, item.Message, map[string]string{})
			}

			if err != nil {
				t.Errorf("error: %s", err)
			}

			lastLog := strings.Split(testWrite.lastLogs[0], " ")
			testWrite.clearLogs()

			if lastLog[0] != item.Level {
				t.Errorf("expected %s, actual %s", item.Level, lastLog[0])
			}

			if lastLog[1] != item.Event {
				t.Errorf("expected %s, actual %s", item.Event, lastLog[1])
			}

			if lastLog[2] != item.Message {
				t.Errorf("expected %s, actual %s", item.Message, lastLog[2])
			}
		})
	}
}

func TestLoggerLogWithMeta(t *testing.T) {
	var tests = []struct {
		Method  string
		Level   string
		Event   string
		Message string
		Meta    map[string]string
	}{
		{"Debug", "Debug", "EventName", "Message...", map[string]string{"key": "value"}},
		{"Info", "Info", "EventName", "Message...", map[string]string{"key": "value"}},
		{"Warn", "Warning", "EventName", "Message...", map[string]string{"key": "value"}},
		{"Error", "Error", "EventName", "Message...", map[string]string{"key": "value"}},
	}

	log := New()
	log.AddTransport(testTransport)

	for _, item := range tests {
		t.Run(item.Method, func(t *testing.T) {
			var err error

			switch item.Method {
			case "Debug":
				err = log.DebugWithMeta(item.Event, item.Message, item.Meta)
			case "Info":
				err = log.InfoWithMeta(item.Event, item.Message, item.Meta)
			case "Warn":
				err = log.WarnWithMeta(item.Event, item.Message, item.Meta)
			case "Error":
				err = log.ErrorWithMeta(item.Event, item.Message, item.Meta)
			default:
				err = log.Log(item.Level, item.Event, item.Message, item.Meta)
			}

			if err != nil {
				t.Errorf("error: %s", err)
			}

			lastLog := strings.Split(testWrite.lastLogs[0], " ")
			testWrite.clearLogs()

			if lastLog[0] != item.Level {
				t.Errorf("expected %s, actual %s", item.Level, lastLog[0])
			}

			if lastLog[1] != item.Event {
				t.Errorf("expected %s, actual %s", item.Event, lastLog[1])
			}

			if lastLog[2] != item.Message {
				t.Errorf("expected %s, actual %s", item.Message, lastLog[2])
			}

			if lastLog[3] != "key=value" {
				t.Errorf("expected %s, actual %s", "key=value", lastLog[3])
			}
		})
	}
}

func TestLoggerWithDefaultMeta(t *testing.T) {
	log := New()
	log.AddTransport(testTransport)

	log.SetDefaultMeta(map[string]string{
		"defaultKey": "defaultValue",
	})

	err := log.Info("SomeEvent", "Message")

	if err != nil {
		t.Errorf("error: %s", err)
	}

	lastLog := strings.Split(testWrite.lastLogs[0], " ")
	testWrite.clearLogs()

	if lastLog[0] != "Info" {
		t.Errorf("expected %s, actual %s", "Info", lastLog[0])
	}

	if lastLog[1] != "SomeEvent" {
		t.Errorf("expected %s, actual %s", "SomeEvent", lastLog[1])
	}

	if lastLog[2] != "Message" {
		t.Errorf("expected %s, actual %s", "Message", lastLog[2])
	}

	if lastLog[3] != "defaultKey=defaultValue" {
		t.Errorf("expected %s, actual %s", "defaultKey=defaultValue", lastLog[3])
	}
}
