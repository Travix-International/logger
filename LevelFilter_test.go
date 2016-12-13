package logger

import "testing"

func TestFilterByMinimumLevel(t *testing.T) {

	levelFilter := NewLevelFilter("Debug")
	filterFunc := FilterByMinimumLevel(levelFilter)

	if filterFunc == nil {
		t.Error("Nil func constructed, this is not expected")
	}

	if !filterFunc(&Entry{Level: "Debug"}) {
		t.Error("Expected that entry at debug level would be allowed, but it isn't")
	}
}

func TestNewLevelFilterWithExactExistingLevel(t *testing.T) {
	minLevel := LogLevel("Info")
	levelFilter := NewLevelFilter(minLevel)
	levelFilter.init()

	if levelFilter.MinLevel != "Info" {
		t.Errorf("Minimum level is not set to %s as expected", minLevel)
	}
}

// TestNewLevelFilterWithCasedExistingLevel harnasses against the case where the minimum
// level is initialized with a different casing.
func TestNewLevelFilterWithCasedExistingLevel(t *testing.T) {
	minLevel := LogLevel("INFO")
	levelFilter := NewLevelFilter(minLevel)
	levelFilter.init()

	if levelFilter.MinLevel != "INFO" {
		t.Errorf("Minimum level is not set to %s as expected", minLevel)
	}

	if levelFilter.filter(&Entry{Level: "Debug"}) {
		t.Error("Expected that entry at debug level would be disallowed, but it isn't")
	}

	if !levelFilter.filter(&Entry{Level: "Info"}) {
		t.Error("Expected that entry at debug level would be allowed, but it isn't")
	}
}

func TestFilter(t *testing.T) {
	minLevel := LogLevel("Info")
	levelFilter := NewLevelFilter(minLevel)

	if levelFilter.filter(&Entry{Level: "Debug"}) {
		t.Error("Expected that entry at debug level would be disallowed, but it isn't")
	}

	if !levelFilter.filter(&Entry{Level: "Info"}) {
		t.Error("Expected that entry at debug level would be allowed, but it isn't")
	}

	if !levelFilter.filter(nil) {
		t.Error("Expected that nil entry would be allowed, but it isn't")
	}
}

func TestSetMinLevel(t *testing.T) {
	minLevel := LogLevel("Info")
	levelFilter := NewLevelFilter(minLevel)

	if levelFilter.MinLevel != "Info" {
		t.Errorf("Minimum level is not set to %s as expected", minLevel)
	}

	levelFilter.SetMinLevel("Warning")
	if levelFilter.MinLevel != "Warning" {
		t.Errorf("Minimum level is not set to %s as expected", minLevel)
	}
}
