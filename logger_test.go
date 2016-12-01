package logger

import (
	"log"
	"testing"
)

type FakeTransport struct {
}

func (transport *FakeTransport) Log(entry Entry) error {
	log.Println(entry.ToString())

	return nil
}

func NewFakeTransport() *FakeTransport {
	transport := &FakeTransport{}

	return transport
}

func TestLoggerLog(t *testing.T) {
	fakeTransport := NewFakeTransport()

	testLogger := New()
	testLogger.AddTransport(fakeTransport)

	err := testLogger.Debug("DebugEvent", "debug message...")

	if err != nil {
		t.Errorf("error: %s", err)
	}
}
