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
	switch strings.ToLower(level) {
	case "trace":
		return Trace
	case "debug":
		return Debug
	case "notice":
		return Notice
	case "warning":
		return Warning
	case "error":
		return Error
	default:
		return Info
	}
}

func (level Level) String() string {
	switch level {
	case Trace:
		return "trace"
	case Debug:
		return "debug"
	case Notice:
		return "notice"
	case Warning:
		return "warning"
	case Error:
		return "error"
	default:
		return "info"
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
