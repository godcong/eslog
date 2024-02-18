package eslog

import (
	"io"
	"log/slog"

	"github.com/lmittmann/tint"
	"github.com/rs/zerolog"
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

type HandlerOptions struct {
	Type        string
	LoggerName  string
	Level       Leveler
	AddSource   bool
	ReplaceAttr func(groups []string, a Attr) Attr
	Converter   ZeroLogConverter
	TimeFormat  string
	NoColor     bool
}

func NewHandler(w any, opts HandlerOptions) Handler {
	switch opts.Type {
	case "zap":
		c, ok := w.(zapcore.Core)
		if !ok {
			panic("zap core is required")
		}
		return NewZapHandler(c, zapslog.HandlerOptions{
			LoggerName: opts.LoggerName,
			AddSource:  opts.AddSource,
		})
	case "zerolog":
		l, ok := w.(*zerolog.Logger)
		if !ok {
			panic("zerolog logger is required")
		}
		return NewZeroLogHandler(l, ZeroLogHandlerOptions{
			Level:     opts.Level,
			Converter: opts.Converter,
		})
	case "tint":
		w, ok := w.(io.Writer)
		if !ok {
			panic("tint logger is required")
		}
		return NewTintHandler(w, TintHandlerOptions{
			AddSource:   opts.AddSource,
			Level:       opts.Level,
			ReplaceAttr: opts.ReplaceAttr,
			TimeFormat:  opts.TimeFormat,
			NoColor:     opts.NoColor,
		})
	case "json":
		w, ok := w.(io.Writer)
		if !ok {
			panic("logger writer is required")
		}
		return slog.NewJSONHandler(w, &slog.HandlerOptions{
			Level:       opts.Level,
			ReplaceAttr: opts.ReplaceAttr,
			AddSource:   opts.AddSource,
		})
	case "text":
		w, ok := w.(io.Writer)
		if !ok {
			panic("logger writer is required")
		}
		return slog.NewTextHandler(w, &slog.HandlerOptions{
			Level:       opts.Level,
			ReplaceAttr: opts.ReplaceAttr,
			AddSource:   opts.AddSource,
		})
	default:
	}
	return slog.Default().Handler()
}

func NewZapHandler(c zapcore.Core, opts zapslog.HandlerOptions) *ZapHandler {
	return zapslog.NewHandler(c, &opts)
}

func NewZeroLogHandler(l *zerolog.Logger, opts ZeroLogHandlerOptions) *ZeroLogHandler {
	if opts.Logger == nil {
		opts.Logger = l
	}
	return opts.NewZerologHandler().(*ZeroLogHandler)
}

func NewTintHandler(w io.Writer, opts TintHandlerOptions) *TintHandler {
	return &TintHandler{
		Handler: tint.NewHandler(w, &opts),
	}
}

func NewLumberJack(opts LumberJackOption) *LumberJackLogger {
	return &LumberJackLogger{
		Filename:   opts.Filename,
		MaxSize:    opts.MaxSize,
		MaxAge:     opts.MaxAge,
		MaxBackups: opts.MaxBackups,
		LocalTime:  opts.LocalTime,
		Compress:   opts.Compress,
	}
}

func init() {
}
