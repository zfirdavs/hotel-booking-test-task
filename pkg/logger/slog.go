package logger

import (
	"io"
	"log/slog"
)

type JSONLogger struct {
	log *slog.Logger
}

func NewJSONLogger(w io.Writer) *JSONLogger {
	return &JSONLogger{
		slog.New(slog.NewJSONHandler(w, nil)),
	}
}

func (j *JSONLogger) Info(msg string, v ...any) {
	j.log.Info(msg, v...)
}

func (j *JSONLogger) Error(msg string, v ...any) {
	j.log.Error(msg, v...)
}
