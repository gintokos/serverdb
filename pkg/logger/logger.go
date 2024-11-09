package logger

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/fatih/color"
)

var logger = &CustomLogger{}

func MustSetupLogger() {
	h := NewCustomHandler("logs.txt")
	l := slog.New(h)
	logger.Logger = l
}

type CustomLogger struct {
	*slog.Logger
}

func GetLogger() *CustomLogger {
	return logger
}

func (cl *CustomLogger) Error(msg string, err error) {
	cl.Logger.Error(fmt.Sprintf("%s: %s",msg, err.Error()))
}

type CustomHandler struct {
	slog.Handler
	file *os.File
}

func NewCustomHandler(filename string) *CustomHandler {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panicln("Error on opening logs file: ", err)
	}

	handler := slog.NewTextHandler(file,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	)

	return &CustomHandler{
		file:    file,
		Handler: handler,
	}
}

func (l *CustomHandler) Close() error {
	return l.file.Close()
}

func (l *CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	l.Handler.Handle(ctx, r)

	level := r.Level.String() + ":"
	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fmt.Print(level, r.Time.Format("02.01.2006 15:04:05:"))
	fmt.Printf(" %s\n", r.Message)

	return nil
}
