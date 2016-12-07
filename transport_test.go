package logger

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

type TestFaultyFormat struct {
}

func (t *TestFaultyFormat) Format(e *Entry) (string, error) {
	return "", errors.New("Something went wrong")
}

func TestNewTransport(t *testing.T) {
	var testTransport = NewTransport(os.Stdout, DefaultStringFormat)

	expected := "*logger.Transport"
	actual := fmt.Sprintf("%s", reflect.TypeOf(testTransport))

	if expected != actual {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}

func TestTransportWithFaultyFormat(t *testing.T) {
	var testTransport = NewTransport(os.Stdout, &TestFaultyFormat{})
	var testMeta = make(map[string]string)
	var entry = NewEntry("Info", "SomeEvent", "Message...", testMeta)

	err := testTransport.log(entry)

	if err == nil {
		t.Error("expected error from TestFaultyFormat")
	}
}
