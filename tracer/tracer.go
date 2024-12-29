/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package tracer

import (
	"os"

	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger"
	"github.com/archnum/sdk.base/logger/handler"
	"github.com/archnum/sdk.base/logger/level"
	"github.com/archnum/sdk.base/uuid"
)

var (
	_logger     = logger.New(uuid.Zero, "TRACER")
	_, _enabled = os.LookupEnv("__TRACER")
)

func init() {
	_logger.SetLevel(level.Trace)
	_logger.AddHandler(handler.Console)
}

func Log(msg string, kvs ...kv.KeyValue) {
	if _enabled {
		_logger.Log(level.Trace, msg, kvs...) //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	}
}

/*
####### END ############################################################################################################
*/
