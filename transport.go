package logger

type ITransport interface {
	Log(entry Entry) error
}
