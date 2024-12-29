/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package main

import (
	"time"

	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger"
	"github.com/archnum/sdk.base/logger/handler"
	"github.com/archnum/sdk.base/logger/level"
	"github.com/archnum/sdk.base/uuid"
)

func main() {
	l := logger.New(uuid.Zero, "main").With(kv.String("env", "devel"))
	l.SetLevel(level.Trace)
	l.AddHandler(handler.Console)

	l.Trace("trace", kv.Any("nil", nil))
	l.Debug("debug", kv.Float("float", 3456789.141))
	l.Info("info", kv.Duration("duration", 8937*time.Second))
	l.Notice("notice", kv.Time("now", time.Now()))
	l.Warning("warning", kv.String("string", "Hello World!"))
	l.Error("error", kv.Bool("bool", false))
}

/*
####### END ############################################################################################################
*/
