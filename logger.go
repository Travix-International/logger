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
	return l.Log("Debug", event, message, map[string]string{})
}

func (l *Logger) DebugWithMeta(event string, message string, meta map[string]string) error {
	return l.Log("Debug", event, message, meta)
}

func (l *Logger) Info(event string, message string) error {
	return l.Log("Info", event, message, map[string]string{})
}

func (l *Logger) InfoWithMeta(event string, message string, meta map[string]string) error {
	return l.Log("Info", event, message, meta)
}

func (l *Logger) Warn(event string, message string) error {
	return l.Log("Warning", event, message, map[string]string{})
}

func (l *Logger) WarnWithMeta(event string, message string, meta map[string]string) error {
	return l.Log("Warning", event, message, meta)
}

func (l *Logger) Error(event string, message string) error {
	return l.Log("Error", event, message, map[string]string{})
}

func (l *Logger) ErrorWithMeta(event string, message string, meta map[string]string) error {
	return l.Log("Error", event, message, meta)
}

/**
 * Common log method
 */
func (l *Logger) Log(level string, event string, message string, meta map[string]string) error {
	entry := NewEntry(
		level,
		event,
		message,
		meta,
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

	return errs
}

/**
 * Instantiation
 */
func New() *Logger {
	l := &Logger{
		[]*Transport{},
	}

	return l
}
