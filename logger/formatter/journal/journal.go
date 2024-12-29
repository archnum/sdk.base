/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package journal

import (
	"github.com/archnum/sdk.base/buffer"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger/level"
	"github.com/archnum/sdk.base/logger/record"
)

const (
	MaxLenLoggerID   = 8
	MaxLenLoggerName = 10
)

type (
	FormatterOptions struct {
		MaxLenLoggerID   int
		MaxLenLoggerName int
	}
)

func (opts *FormatterOptions) fix() {
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

	buf.AppendARString(rec.LoggerName, '.', impl.opts.MaxLenLoggerName)
	buf.AppendByte(':')
	buf.AppendALString(rec.LoggerID, '.', impl.opts.MaxLenLoggerID)
	buf.AppendByte(' ')
	buf.AppendString(rec.Message)

	if len(rec.KeyValues) > 0 {
		buf.AppendString(": ")
		kv.Logfmt(buf, rec.KeyValues)
	}
}

/*
####### END ############################################################################################################
*/
