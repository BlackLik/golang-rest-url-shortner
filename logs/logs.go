package logs

import (
	"os"

	"golang.org/x/exp/slog"
)

const (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARNING"
	LevelError = "ERROR"
)

// SetLevel sets the logging level for the application.
//
// level: The desired logging level. Valid values are LevelDebug, LevelInfo,
// LevelWarn, and LevelError.
// Return type: None.
func SetLevel(level string) {

	levelSlog := slog.LevelDebug

	switch level {
	case LevelDebug:
		levelSlog = slog.LevelDebug
	case LevelInfo:
		levelSlog = slog.LevelInfo
	case LevelWarn:
		levelSlog = slog.LevelWarn
	case LevelError:
		levelSlog = slog.LevelError
	default:
		levelSlog = slog.LevelDebug
	}

	var h = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: levelSlog,
	})
	slog.SetDefault(slog.New(h))
}
