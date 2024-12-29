/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package logfmt

import (
	"github.com/archnum/sdk.base/buffer"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger/level"
	"github.com/archnum/sdk.base/logger/record"
)

const (
	TimestampLayout = "2006-01-02 15:04:05.000"

	MaxLenLoggerID   = 8
	MaxLenLoggerName = 10
)

type (
	FormatterOptions struct {
		TimestampLayout  string
		WithColors       bool
		MaxLenLoggerID   int
		MaxLenLoggerName int
	}
)

func (opts *FormatterOptions) fix() {
	if opts.TimestampLayout == "" {
		opts.TimestampLayout = TimestampLayout
	}

	if opts.MaxLenLoggerID == 0 {
		opts.MaxLenLoggerID = MaxLenLoggerID
	}

	if opts.MaxLenLoggerName == 0 {
		opts.MaxLenLoggerName = MaxLenLoggerName
	}
}

type (
	implFormatter struct {
		opts FormatterOptions
	}
)

func NewFormatter(opts *FormatterOptions) *implFormatter {
	if opts == nil {
		opts = &FormatterOptions{}
	}

	clone := *opts
	clone.fix()

	return &implFormatter{opts: clone}
}

func (impl *implFormatter) Format(buf *buffer.Buffer, rec *record.Record) {
	buf.AppendTime(rec.Timestamp, impl.opts.TimestampLayout)
	buf.AppendByte(' ')

	if impl.opts.WithColors {
		switch rec.Level {
		case level.Trace:
			buf.AppendString("\033[46m\033[30m")
		case level.Debug:
			buf.AppendString("\033[47m\033[30m")
		case level.Notice:
			buf.AppendString("\033[42m\033[30m")
		case level.Warning:
			buf.AppendString("\033[43m\033[30m")
		case level.Error:
			buf.AppendString("\033[41m\033[30m")
		}
	} else {
		switch rec.Level {
		case level.Trace:
			buf.AppendString("TRA ")
		case level.Debug:
			buf.AppendString("DEB ")
		case level.Notice:
			buf.AppendString("NOT ")
		case level.Warning:
			buf.AppendString("WAR ")
		case level.Error:
			buf.AppendString("ERR ")
		default:
			buf.AppendString("INF ")
		}
	}

	buf.AppendARString(rec.LoggerName, '.', impl.opts.MaxLenLoggerName)
	buf.AppendByte(':')
	buf.AppendALString(rec.LoggerID, '.', impl.opts.MaxLenLoggerID)

	if impl.opts.WithColors {
		buf.AppendString("\033[0m ")
	} else {
		buf.AppendByte(' ')
	}

	buf.AppendString(rec.Message)

	if len(rec.KeyValues) > 0 {
		buf.AppendString(": ")
		kv.Logfmt(buf, rec.KeyValues)
	}

	buf.AppendByte('\n')
}

/*
####### END ############################################################################################################
*/
