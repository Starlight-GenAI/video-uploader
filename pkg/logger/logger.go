package logger

import (
	"os"

	"log/slog"
)

var Logger *slog.Logger

func Init() {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(Logger)
}
