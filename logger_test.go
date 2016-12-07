package logger

import (
	"bytes"
	"errors"
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
				err = log.Log(item.Level, item.Event, item.Message)
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

func TestLoggerLogf(t *testing.T) {
	var tests = []struct {
		Method  string
		Level   string
		Event   string
		Message string
		Replace string
	}{
		{"Debug", "Debug", "EventName", "Message...%s", "Here"},
		{"Info", "Info", "EventName", "Message...%s", "Here"},
		{"Warn", "Warning", "EventName", "Message...%s", "Here"},
		{"Error", "Error", "EventName", "Message...%s", "Here"},
		{"CustomLevel", "CustomLevel", "EventName", "Message...%s", "Here"},
	}

	log := New()
	log.AddTransport(testTransport)

	for _, item := range tests {
		t.Run(item.Method, func(t *testing.T) {
			var err error

			switch item.Method {
			case "Debug":
				err = log.Debugf(item.Event, item.Message, item.Replace)
			case "Info":
				err = log.Infof(item.Event, item.Message, item.Replace)
			case "Warn":
				err = log.Warnf(item.Event, item.Message, item.Replace)
			case "Error":
				err = log.Errorf(item.Event, item.Message, item.Replace)
			default:
				err = log.Logf(item.Level, item.Event, item.Message, item.Replace)
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

			expectedMessage := fmt.Sprintf(item.Message, item.Replace)
			if lastLog[2] != expectedMessage {
				t.Errorf("expected %s, actual %s", expectedMessage, lastLog[2])
			}
		})
	}
}

func TestLoggerExceptions(t *testing.T) {
	log := New()
	log.AddTransport(testTransport)

	customErr := errors.New("Something wrong here")

	err := log.Exceptionf("SomeEvent", customErr, "message...%s", "here")

	if err != nil {
		t.Errorf("error: %s", err)
	}

	expected := "Error SomeEvent message...here errormessage=Something wrong here"
	actual := testWrite.lastLogs[0]

	if expected != actual {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}
