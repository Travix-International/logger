package console

import (
	"log"

	"github.com/Travix-International/logger"
)

type Console struct {
}

func (transport *Console) Log(entry logger.Entry) error {
	log.Println(entry.ToString())

	return nil
}

func New() Console {
	transport := Console{}

	return transport
}
