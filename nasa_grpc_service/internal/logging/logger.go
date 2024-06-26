package logging

import (
	"log/slog"
	"os"
)

func New(environment string, f *os.File) *slog.Logger {
	var logger *slog.Logger

	switch environment {
	case "testing":
		logger = setUpTestingLogger(f)
	case "production":
		logger = setUpProductionLogger(f)
	default:
		logger = setUpDefaultLogger()
	}

	return logger
}

func setUpTestingLogger(f *os.File) *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(
			f,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		),
	)
}

func setUpProductionLogger(f *os.File) *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(
			f,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			},
		),
	)
}

func setUpDefaultLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		),
	)
}
