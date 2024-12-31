/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package level

import (
	"strings"
	"sync/atomic"
)

type (
	Level int
)

const (
	Trace   Level = -2
	Debug   Level = -1
	Info    Level = 0 // Zero Var
	Notice  Level = 1
	Warning Level = 2
	Error   Level = 3
)

func StringToLevel(level string) Level {
	switch strings.ToUpper(level) {
	case "TRACE":
		return Trace
	case "DEBUG":
		return Debug
	case "NOTICE":
		return Notice
	case "WARNING":
		return Warning
	case "ERROR":
		return Error
	default:
		return Info
	}
}

func (level Level) String() string {
	switch level {
	case Trace:
		return "TRACE"
	case Debug:
		return "DEBUG"
	case Notice:
		return "NOTICE"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	default:
		return "INFO"
	}
}

type (
	Var struct {
		level atomic.Int64
	}
)

func (v *Var) Level() Level {
	return Level(int(v.level.Load()))
}

func (v *Var) Set(level Level) {
	v.level.Store(int64(level))
}

/*
####### END ############################################################################################################
*/
