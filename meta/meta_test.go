package meta

import (
	"testing"
)

func TestAdd(t *testing.T) {
	m := New()
	m.Set("foo", "bar")

	expected := "bar"
	actual := m.Fields["foo"]

	if expected != actual {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}
