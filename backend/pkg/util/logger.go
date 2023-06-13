package util

import (
	"github.com/fatih/color"
	"github.com/m-mizutani/clog"
	"golang.org/x/exp/slog"
)

var Logger = slog.New(clog.New(
	clog.WithColor(true),
	clog.WithSource(true),
	clog.WithLevel(slog.LevelDebug), // TODO: set info in production
	clog.WithTimeFmt("2006-01-02 15:04:05"),
	clog.WithColorMap(&clog.ColorMap{
		Level: map[slog.Level]*color.Color{
			slog.LevelDebug: color.New(color.FgGreen, color.Bold),
			slog.LevelInfo:  color.New(color.FgCyan, color.Bold),
			slog.LevelWarn:  color.New(color.FgYellow, color.Bold),
			slog.LevelError: color.New(color.FgRed, color.Bold),
		},
		LevelDefault: color.New(color.FgBlue, color.Bold),
		Time:         color.New(color.FgWhite),
		Message:      color.New(color.FgHiWhite),

		AttrKey:   color.New(color.FgBlue),
		AttrValue: color.New(color.FgHiWhite),
	}),
))
