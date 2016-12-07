package logger

import (
	"testing"
)

func TestJSONFormat(t *testing.T) {
	var testJSONFormat = NewJSONFormat()

	var testMeta = make(map[string]string)
	testMeta["key"] = "value"

	var tests = []struct {
		Level   string
		Event   string
		Message string
		Meta    map[string]string

		Expected string
	}{
		{"Info", "EventName", "Message...", testMeta, `{"level":"Info","event":"EventName","message":"Message...","key":"value"}`},
		{"Debug", "EventName", "blah...", testMeta, `{"level":"Debug","event":"EventName","message":"blah...","key":"value"}`},
		{"Error", "EventName", "blah...", testMeta, `{"level":"Error","event":"EventName","message":"blah...","key":"value"}`},
	}

	for _, item := range tests {
		t.Run(item.Level, func(t *testing.T) {
			entry := NewEntry(item.Level, item.Event, item.Message, item.Meta)

			expected := item.Expected
			actual, err := testJSONFormat.Format(entry)

			if err != nil {
				t.Errorf("error: %s", err)
			}

			if expected != actual {
				t.Errorf("expected %s, actual %s", expected, actual)
			}
		})
	}
}

func TestJSONFormatForErrors(t *testing.T) {
	var testJSONFormat = NewJSONFormat()
	var testMeta = make(map[string]string)

	t.Run("ErrNoLevel", func(t *testing.T) {
		entry := NewEntry("", "SomeEvent", "Message...", testMeta)

		_, err := testJSONFormat.Format(entry)

		if err == nil {
			t.Error("expected error ErrNoLevel")
		}
	})

	t.Run("ErrNoEvent", func(t *testing.T) {
		entry := NewEntry("Info", "", "Message...", testMeta)

		_, err := testJSONFormat.Format(entry)

		if err == nil {
			t.Error("expected error ErrNoEvent")
		}
	})

	t.Run("ErrNoMessage", func(t *testing.T) {
		entry := NewEntry("Info", "SomeEvent", "", testMeta)

		_, err := testJSONFormat.Format(entry)

		if err == nil {
			t.Error("expected error ErrNoMessage")
		}
	})
}

func TestJSONFormatWithNoEntry(t *testing.T) {
	var testJSONFormat = NewJSONFormat()

	str, err := testJSONFormat.Format(nil)

	if err != nil {
		t.Errorf("error: %s", err)
	}

	if "" != str {
		t.Errorf("expected %s, actual %s", "", str)
	}
}

func TestJSONFormatWithCustomKeys(t *testing.T) {
	var testJSONFormat = NewJSONFormat().
		SetLevelKey("levelName").
		SetEventKey("eventName").
		SetMessageKey("messageBody")

	var testMeta = make(map[string]string)

	var tests = []struct {
		Level   string
		Event   string
		Message string
		Meta    map[string]string

		Expected string
	}{
		{"Info", "EventName", "Message...", testMeta, `{"levelName":"Info","eventName":"EventName","messageBody":"Message..."}`},
		{"Debug", "EventName", "blah...", testMeta, `{"levelName":"Debug","eventName":"EventName","messageBody":"blah..."}`},
		{"Error", "EventName", "blah...", testMeta, `{"levelName":"Error","eventName":"EventName","messageBody":"blah..."}`},
	}

	for _, item := range tests {
		t.Run(item.Level, func(t *testing.T) {
			entry := NewEntry(item.Level, item.Event, item.Message, item.Meta)

			expected := item.Expected
			actual, err := testJSONFormat.Format(entry)

			if err != nil {
				t.Errorf("error: %s", err)
			}

			if expected != actual {
				t.Errorf("expected %s, actual %s", expected, actual)
			}
		})
	}
}
