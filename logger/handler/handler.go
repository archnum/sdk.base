/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package handler

import (
	"os"

	"github.com/archnum/sdk.base/logger/formatter"
	"github.com/archnum/sdk.base/logger/formatter/logfmt"
	"github.com/archnum/sdk.base/logger/handler/journal"
	"github.com/archnum/sdk.base/logger/handler/writer"
	"github.com/archnum/sdk.base/logger/level"
)

type (
	Handler interface {
		Name() string
		Level() level.Level
		SetLevel(lvl level.Level)
		Formatter() formatter.Formatter
		Log(lvl level.Level, msg []byte) error
	}
)

var (
	Console = writer.NewHandler(
		"console",
		level.Trace,
		logfmt.NewFormatter(&logfmt.FormatterOptions{WithColors: true}),
		os.Stdout,
	)

	Journal = journal.NewHandler("journal", level.Trace, nil)
)

/*
####### END ############################################################################################################
*/
