package eslog

import "log/slog"

// New creates a new Logger with the given non-nil Handler.
func New(h Handler) *Logger {
	if h == nil {
		panic("nil Handler")
	}
	return slog.New(h)
}
