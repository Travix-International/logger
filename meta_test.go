package logger

import (
	"testing"
)

func TestMetaSet(t *testing.T) {
	m := NewMeta()
	m.Set("foo", "bar")

	expected := "bar"
	actual := m.Fields["foo"]

	if expected != actual {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}

func TestMetaGet(t *testing.T) {
	m := NewMeta()
	m.Set("foo", "bar")

	expected := "bar"
	actual := m.Get("foo")

	if expected != actual {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}

func TestMetaRemove(t *testing.T) {
	m := NewMeta()
	m.Set("foo", "bar")
	m.Remove("foo")

	expected := ""
	actual := m.Get("foo")

	if expected != actual {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}

func TestMetaGetFields(t *testing.T) {
	var fixtures = []struct {
		key   string
		value string
	}{
		{"foo", "bar"},
		{"baz", "qux"},
		{"hello", "world"},
	}

	m := NewMeta()

	for _, fixture := range fixtures {
		m.Set(fixture.key, fixture.value)
	}

	fields := m.GetFields()

	for _, fixture := range fixtures {
		expected := fixture.value
		actual := fields[fixture.key]

		if expected != actual {
			t.Errorf("expected %s, actual %s", expected, actual)
		}
	}
}
