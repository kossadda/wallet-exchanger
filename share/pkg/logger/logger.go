package logger

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/fatih/color"
)

// Constants for different environments.
const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// HandlerOptions contains options for configuring the logger handler.
type HandlerOptions struct {
	SlogOpts *slog.HandlerOptions
}

// Handler is a custom logger handler that adds color formatting and structured logging.
type Handler struct {
	opts HandlerOptions
	slog.Handler
	l     *log.Logger
	attrs []slog.Attr
}

// NewHandler creates and returns a new Handler with the specified output writer.
func (opts HandlerOptions) NewHandler(out io.Writer) *Handler {
	h := &Handler{
		Handler: slog.NewJSONHandler(out, opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}

// Handle processes a log record, formats it with color, and outputs it to the writer.
func (h *Handler) Handle(_ context.Context, r slog.Record) error {
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

	fields := make(map[string]interface{}, r.NumAttrs())

	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	var b []byte
	var err error

	if len(fields) > 0 {
		b, err = json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(
		timeStr,
		level,
		msg,
		color.WhiteString(string(b)),
	)

	return nil
}

// WithAttrs returns a new handler with the specified attributes added.
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{
		Handler: h.Handler,
		l:       h.l,
		attrs:   attrs,
	}
}

// WithGroup returns a new handler with the specified group name.
func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}

// SetupLogger initializes and returns a new logger with a default handler.
func SetupLogger() *slog.Logger {
	opts := HandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewHandler(os.Stdout)

	return slog.New(handler)
}

// SetupByEnv sets up the logger based on the provided environment string.
func SetupByEnv(env string) *slog.Logger {
	switch env {
	case envLocal:
		return SetupLogger()
	case envDev:
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return SetupLogger()
}

// Err creates an error attribute for logging.
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
