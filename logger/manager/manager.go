/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package manager

import (
	"github.com/archnum/sdk.base/logger/level"
	"github.com/archnum/sdk.base/uuid"
)

var (
	_manager = newManager()
)

type (
	Logger interface {
		ID() uuid.UUID
		Name() string
		Level() level.Level
		SetLevel(level level.Level)
	}

	manager struct {
		callback func(l Logger)
	}
)

func newManager() *manager {
	return &manager{}
}

func (m *manager) registerCallback(cb func(Logger)) {
	m.callback = cb
}

func RegisterCallback(cb func(Logger)) {
	_manager.registerCallback(cb)
}

func (m *manager) registerLogger(l Logger) {
	m.callback(l)
}

func RegisterLogger(l Logger) {
	_manager.registerLogger(l)
}

/*
####### END ############################################################################################################
*/
