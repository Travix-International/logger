package logger

import (
	"io"
)

type Transport struct {
	destination io.Writer
	formatter   Formatter
}

type Formatter interface {
	Format(entry *Entry) (string, error)
}

func NewTransport(w io.Writer, f Formatter) *Transport {
	return &Transport{
		w,
		f,
	}
}

func (t *Transport) log(e *Entry) (err error) {
	whatToWrite, err := t.formatter.Format(e)

	if err != nil {
		return err
	}

	_, err = t.destination.Write([]byte(whatToWrite))

	return err
}
