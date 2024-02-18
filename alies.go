package eslog

import (
	"log/slog"

	"github.com/lmittmann/tint"
	slogzerolog "github.com/samber/slog-zerolog"
	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type (
	Logger  = slog.Logger
	Handler = slog.Handler
	Level   = slog.Level
	Leveler = slog.Leveler
	Record  = slog.Record
	Attr    = slog.Attr
	Value   = slog.Value
)

type (
	ZapHandler        = zapslog.Handler
	ZapLogger         = zap.Logger
	ZapHandlerOptions = zapslog.HandlerOptions
)

type (
	ZeroLogHandler        = slogzerolog.ZerologHandler
	ZeroLogHandlerOptions = slogzerolog.Option
)

type (
	LumberJackLogger = lumberjack.Logger
)

type (
	TintHandlerOptions = tint.Options

	TintHandler struct {
		slog.Handler
	}
)

const (
	TimeKey    = slog.TimeKey
	LevelKey   = slog.LevelKey
	MessageKey = slog.MessageKey
	SourceKey  = slog.SourceKey
)

const (
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
	// LevelFatal = slog.LevelFatal
)

type LumberJackOption struct {
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	LocalTime  bool
	Compress   bool
}

func init() {
}
