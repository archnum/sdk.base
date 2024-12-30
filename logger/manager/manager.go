/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package manager

import (
	"sync"

	"github.com/archnum/sdk.base/logger/level"
	"github.com/archnum/sdk.base/uuid"
)

var (
	_manager = newManager()
)

type (
	manager struct {
		loggers map[uuid.UUID]*Logger
		mutex   sync.Mutex
	}
)

func newManager() *manager {
	return &manager{
		loggers: make(map[uuid.UUID]*Logger),
	}
}

func (m *manager) registerLogger(l *Logger) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.loggers[l.ID] = l
}

func RegisterLogger(id uuid.UUID, name string, level level.Level) {
	l := &Logger{
		ID:    id,
		Name:  name,
		Level: level,
	}

	_manager.registerLogger(l)
}

func (m *manager) setLevel(id uuid.UUID, level level.Level) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if l, ok := m.loggers[id]; ok {
		l.Level = level
	}
}

func SetLevel(id uuid.UUID, level level.Level) {
	_manager.setLevel(id, level)
}

/*
####### END ############################################################################################################
*/
