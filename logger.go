package logger

/**
 * Base struct
 */
type Logger struct {
	meta       Meta
	transports []ITransport
}

/**
 * Transport methods
 */
func (l *Logger) AddTransport(transport ITransport) *Logger {
	l.transports = append(l.transports, transport)

	return l
}

/**
 * Level methods
 */
func (l *Logger) Debug(event string, message string) error {
	return l.Log("Debug", event, message)
}

func (l *Logger) Info(event string, message string) error {
	return l.Log("Info", event, message)
}

func (l *Logger) Warn(event string, message string) error {
	return l.Log("Warning", event, message)
}

func (l *Logger) Error(event string, message string) error {
	return l.Log("Error", event, message)
}

func (l *Logger) Exception(event string, err error, message string) error {
	// @TODO: how to catch error message and stacktrace from `err`?

	// @TODO: l.meta.Set("exceptionmessage", "...")
	// @TODO: l.meta.Set("exceptiondetails", "...")

	return l.Log("Error", event, message)
}

/**
 * Common log method
 */
func (l *Logger) Log(level string, event string, message string) error {
	entry := Entry{
		Level:   level,
		Event:   event,
		Message: message,
		Meta:    l.meta.GetFields(),
	}

	for i := 0; i < len(l.transports); i++ {
		transport := l.transports[i]

		err := transport.Log(entry)

		if err != nil {
			l.meta = l.Meta()

			return err
		}
	}

	l.meta = l.Meta()

	return nil
}

/**
 * Meta methods
 */
func (l *Logger) WithMeta(meta Meta) *Logger {
	l.meta = meta

	return l
}

func (l *Logger) Meta() Meta {
	return NewMeta()
}

/**
 * Instantiation
 */
func New() Logger {
	l := Logger{}

	l.meta = l.Meta()
	l.transports = make([]ITransport, 0)

	return l
}
