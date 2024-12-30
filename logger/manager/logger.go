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
		ID    uuid.UUID
		Name  string
		Level level.Level
	}
)

/*
####### END ############################################################################################################
*/
