package logger

type Log interface {
	Info(msg string, v ...any)
	Error(msg string, v ...any)
}
