package logger

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type jsonEntry struct {
	Level   string            `json:"level"`
	Event   string            `json:"event"`
	Message string            `json:"message"`
	Meta    map[string]string `json:"meta,omitempty"`
}

var (
	ErrNoLevel   = errors.New("No Level!")
	ErrNoEvent   = errors.New("No Event!")
	ErrNoMessage = errors.New("No Message!")
)

func (j jsonEntry) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if len(j.Level) == 0 {
		return nil, ErrNoLevel
	}
	if len(j.Event) == 0 {
		return nil, ErrNoEvent
	}
	if len(j.Message) == 0 {
		return nil, ErrNoMessage
	}
	_, err := buf.WriteString(fmt.Sprintf("{\"level\":\"%s\", \"event\":\"%s\", \"message\":\"%s\"", j.Level, j.Event, j.Message))
	if err != nil {
		return nil, err
	}

	if len(j.Meta) > 0 {
		_, err = buf.WriteString(",")
		if err != nil {
			return nil, err
		}
		i := len(j.Meta)
		for k, v := range j.Meta {
			_, err = buf.WriteString(fmt.Sprintf(" \"%s\":\"%s\"", k, v))
			if err != nil {
				return nil, err
			}

			i = i - 1
			if i > 0 {
				_, err = buf.WriteString(",")
				if err != nil {
					return nil, err
				}
			}
		}
	}

	_, err = buf.WriteString("}")
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type JSONFormat struct{}

var DefaultJSONFormat = &JSONFormat{}

func (j *JSONFormat) Format(e *Entry) (string, error) {
	if e == nil {
		return "", nil
	}

	if j == nil {
		return "JSONFormat not initialized", nil
	}

	item := jsonEntry{
		Level:   e.Level,
		Event:   e.Event,
		Message: e.Message,
		Meta:    e.Meta,
	}

	out, err := json.Marshal(item)
	return string(out), err
}
