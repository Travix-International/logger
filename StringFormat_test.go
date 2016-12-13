package logger

import (
	"strings"
	"testing"
)

func TestStringFormat(t *testing.T) {
	var testStringFormat = NewStringFormat("[%s]", " %s: ", "%s", "\n%s=", "%s", "\nend\n")

	var testMeta = make(map[string]string)
	testMeta["foo"] = "bar"
	testMeta["key"] = "value"

	var tests = []struct {
		Level   string
		Event   string
		Message string
		Meta    map[string]string

		Expected     string
		ExpectedMeta [2]string
	}{
		{"Info", "EventName", "Message...", testMeta, "[Info] EventName: Message...", [2]string{"foo=bar", "key=value"}},
		{"Debug", "EventName", "blah...", testMeta, "[Debug] EventName: blah...", [2]string{"foo=bar", "key=value"}},
		{"Error", "EventName", "blah...", testMeta, "[Error] EventName: blah...", [2]string{"foo=bar", "key=value"}},
	}

	for _, item := range tests {
		t.Run(item.Level, func(t *testing.T) {
			entry := NewEntry(item.Level, item.Event, item.Message, item.Meta)

			expected := item.Expected
			actual, err := testStringFormat.Format(entry)

			if err != nil {
				t.Errorf("error: %s", err)
			}

			if strings.HasPrefix(actual, expected) != true {
				t.Errorf("expected to begin with %s, actual %s", expected, actual)
			}

			for _, m := range item.ExpectedMeta {
				if strings.Index(actual, m) == -1 {
					t.Errorf("expected to find meta %s, actual %s", m, actual)
				}
			}

			if strings.HasSuffix(actual, testStringFormat.lineSuffix) != true {
				t.Errorf("expected to end with %s, actual %s", testStringFormat.lineSuffix, actual)
			}
		})
	}
}

func TestStringFormatWithNoEntry(t *testing.T) {
	var testStringFormat = NewStringFormat("[%s]", " %s: ", "%s", "\n%s=", "%s", "")

	str, err := testStringFormat.Format(nil)

	if err != nil {
		t.Errorf("error: %s", err)
	}

	if "" != str {
		t.Errorf("expected %s, actual %s", "", str)
	}
}

func TestStringFormatWithNoLevel(t *testing.T) {
	var testStringFormat = NewStringFormat("", " %s: ", "%s", "%s", "%s", "")
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
