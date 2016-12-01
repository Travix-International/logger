package logger

/**
 * Base struct
 */
type Logger struct {
	meta Meta
	// transports [] // map/array of Transports
}

/**
 * Transport methods
 */
func (l *Logger) AddTransport() *Logger {

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
	// entry := Entry{
	// 	Level:   level,
	// 	Event:   event,
	// 	Message: message,
	// 	Meta:    l.meta.GetFields(),
	// }

	// loop through transports
	// call transport.log(entry)

	// empty local meta
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

	return l
}
