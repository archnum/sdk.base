/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package manager

import (
	"github.com/archnum/sdk.base/logger/level"
	"github.com/archnum/sdk.base/uuid"
)

type (
	Logger struct {
		id    uuid.UUID
		name  string
		level level.Level
	}
)

func (l *Logger) ID() string {
	return string(l.id)
}

func (l *Logger) Name() string {
	return l.name
}

func (l *Logger) Level() string {
	return l.level.String()
}

/*
####### END ############################################################################################################
*/
