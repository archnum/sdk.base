/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package formatter

import (
	"github.com/archnum/sdk.base/buffer"
	"github.com/archnum/sdk.base/logger/record"
)

type (
	Formatter interface {
		Format(buf *buffer.Buffer, rec *record.Record)
	}
)

/*
####### END ############################################################################################################
*/
