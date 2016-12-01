package logger

type Entry struct {
	Level   string
	Event   string
	Message string
	Meta    map[string]string
}

func (e *Entry) ToJSON() string {
	return ""
}

func (e *Entry) ToString() string {
	str := "[" + e.Level + "] " + e.Event + ": " + e.Message

	if len(e.Meta) > 0 {
		for k, v := range e.Meta {
			str += "\n" + k + ": " + v
		}
	}

	return str
}

func NewEntry(level string, event string, message string, meta map[string]string) Entry {
	entry := Entry{
		Level:   level,
		Event:   event,
		Message: message,
		Meta:    meta,
	}

	return entry
}
