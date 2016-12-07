package logger

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
)

/**
 * Base struct
 */
type Logger struct {
	Meta map[string]string

	transports []*Transport
}

/**
 * Transport methods
 */
func (l *Logger) AddTransport(t ...*Transport) *Logger {
	for _, transport := range t {
		l.transports = append(l.transports, transport)
	}

	return l
}

/**
 * Level methods
 */
func (l *Logger) Debug(event string, message string) error {
	return l.Log("Debug", event, message)
}

func (l *Logger) Debugf(event string, messageFormat string, params ...interface{}) error {
	return l.Log("Debug", event, fmt.Sprintf(messageFormat, params...))
}

func (l *Logger) Info(event string, message string) error {
	return l.Log("Info", event, message)
}

func (l *Logger) Infof(event string, messageFormat string, params ...interface{}) error {
	return l.Log("Info", event, fmt.Sprintf(messageFormat, params...))
}

func (l *Logger) Warn(event string, message string) error {
	return l.Log("Warning", event, message)
}

func (l *Logger) Warnf(event string, messageFormat string, params ...interface{}) error {
	return l.Log("Warning", event, fmt.Sprintf(messageFormat, params...))
}

func (l *Logger) Error(event string, message string) error {
	return l.Log("Error", event, message)
}

func (l *Logger) Errorf(event string, messageFormat string, params ...interface{}) error {
	return l.Log("Error", event, fmt.Sprintf(messageFormat, params...))
}

/**
 * Common log method
 */
func (l *Logger) Log(level string, event string, message string) error {
	entry := NewEntry(
		level,
		event,
		message,
		l.Meta,
	)

	var wg sync.WaitGroup

	errBuff := []error{}
	var errs error

	e := make(chan error, 1)
	done := make(chan bool, 1)

	for _, t := range l.transports {
		wg.Add(1)

		go func(transport *Transport) {
			err := transport.log(entry)

			if err != nil {
				e <- err
			}

			wg.Done()
		}(t)
	}

	go func() {
		wg.Wait()

		done <- true
	}()

out:
	for {
		select {
		case err := <-e:
			errBuff = append(errBuff, err)
		case <-done:
			if len(errBuff) > 0 {
				var buf bytes.Buffer

				for i, v := range errBuff {
					buf.WriteString(fmt.Sprintf("Error %d: %v", i, v))
				}

				errs = errors.New(buf.String())
			}

			break out
		}
	}

	if len(l.Meta) > 0 {
		l.Meta = make(map[string]string)
	}

	return errs
}

func (l *Logger) Logf(level string, event string, message string, params ...interface{}) error {
	m := fmt.Sprintf(message, params...)

	return l.Log(level, event, m)
}

/**
 * Instantiation
 */
func New() *Logger {
	l := &Logger{
		make(map[string]string),
		[]*Transport{},
	}

	return l
}
