/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package logger

import (
	"log"

	"github.com/archnum/sdk.base/logger/level"
)

type (
	adapter struct {
		logger *Logger
		level  level.Level
	}
)

func (a *adapter) Write(msg []byte) (int, error) {
	a.logger.Log(a.level, string(msg))
	return len(msg), nil
}

func (l *Logger) NewStdLogger(level level.Level, prefix string, flag int) *log.Logger {
	a := &adapter{
		logger: l,
		level:  level,
	}

	return log.New(a, prefix, flag)
}

/*
####### END ############################################################################################################
*/
