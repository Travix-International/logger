package logger

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestEntry(t *testing.T) {
	meta := make(map[string]string)
	meta["key"] = "value"
	entry := NewEntry("Info", "SomeEvent", "Message...", meta)

	t.Run("TypeOf", func(t *testing.T) {
		expected := "*logger.Entry"
		actual := fmt.Sprintf("%s", reflect.TypeOf(entry))

		if expected != actual {
			t.Errorf("expected %s, actual %s", expected, actual)
		}
	})

	t.Run("Fields", func(t *testing.T) {
		if entry.Level != "Info" {
			t.Errorf("expected %s, actual %s", "Info", entry.Level)
		}

		if entry.Event != "SomeEvent" {
			t.Errorf("expected %s, actual %s", "SomeEvent", entry.Event)
		}

		if entry.Message != "Message..." {
			t.Errorf("expected %s, actual %s", "Message...", entry.Message)
		}
	})

	t.Run("String", func(t *testing.T) {
		str := entry.String()

		expectedPrefix := "[Info] SomeEvent: Message..."
		if strings.HasPrefix(str, expectedPrefix) != true {
			t.Errorf("expected prefix: %s, actual string %s", expectedPrefix, str)
		}
	})
}
