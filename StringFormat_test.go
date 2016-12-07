package logger

import (
	"testing"
)

func TestStringFormat(t *testing.T) {
	var testStringFormat = NewStringFormat("[%s]", " %s: ", "%s", "\n%s=", "%s")

	var testMeta = make(map[string]string)
	testMeta["foo"] = "bar"
	testMeta["key"] = "value"

	var tests = []struct {
		Level   string
		Event   string
		Message string
		Meta    map[string]string

		Expected string
	}{
		{"Info", "EventName", "Message...", testMeta, "[Info] EventName: Message...\nfoo=bar\nkey=value"},
		{"Debug", "EventName", "blah...", testMeta, "[Debug] EventName: blah...\nfoo=bar\nkey=value"},
		{"Error", "EventName", "blah...", testMeta, "[Error] EventName: blah...\nfoo=bar\nkey=value"},
	}

	for _, item := range tests {
		t.Run(item.Level, func(t *testing.T) {
			entry := NewEntry(item.Level, item.Event, item.Message, item.Meta)

			expected := item.Expected
			actual, err := testStringFormat.Format(entry)

			if err != nil {
				t.Errorf("error: %s", err)
			}

			if expected != actual {
				t.Errorf("expected %s, actual %s", expected, actual)
			}
		})
	}
}

func TestStringFormatWithNoEntry(t *testing.T) {
	var testStringFormat = NewStringFormat("[%s]", " %s: ", "%s", "\n%s=", "%s")

	str, err := testStringFormat.Format(nil)

	if err != nil {
		t.Errorf("error: %s", err)
	}

	if "" != str {
		t.Errorf("expected %s, actual %s", "", str)
	}
}

func TestStringFormatWithNoLevel(t *testing.T) {
	var testStringFormat = NewStringFormat("", " %s: ", "%s", "%s", "%s")
	var testMeta = make(map[string]string)
	testMeta["key"] = "value"

	entry := NewEntry("Info", "SomeEvent", "Message...", testMeta)

	expected := entry.String()
	actual, err := testStringFormat.Format(entry)

	if err != nil {
		t.Errorf("error: %s", err)
	}

	if expected != actual {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}
