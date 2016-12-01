package console

import (
	"testing"

	"github.com/Travix-International/logger"
)

func TestConsoleLog(t *testing.T) {
	c := New()

	meta := make(map[string]string)
	meta["key"] = "value"
	meta["someOtherKey"] = "value here"

	entry := logger.NewEntry("Info", "SomeEvent", "message here...", meta)

	err := c.Log(entry)

	if err != nil {
		t.Errorf("error: %s", err)
	}
}
