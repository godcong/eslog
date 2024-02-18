package eslog

import (
	"io"

	"github.com/lmittmann/tint"
	"github.com/rs/zerolog"
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

type HandlerOptions struct {
	Level        Leveler
	ReplaceAttrs map[string]func(groups []string, a Attr) Attr
}

func NewZapHandler(c zapcore.Core, opts zapslog.HandlerOptions) *ZapHandler {
	return zapslog.NewHandler(c, &opts)
}

func NewZeroLogHandler(l zerolog.Logger, opts ZeroLogHandlerOptions) *ZeroLogHandler {
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
