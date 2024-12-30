/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package record

import (
	"sync"
	"time"

	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger/level"
)

const (
	_maxKV = 20
)

type (
	// fieldalignment
	Record struct {
		Timestamp  time.Time
		LoggerID   string
		LoggerName string
		Message    string
		KeyValues  []kv.KeyValue
		Level      level.Level
	}
)

var (
	_poolRecord = sync.Pool{
		New: func() any {
			return &Record{
				KeyValues: make([]kv.KeyValue, 0, _maxKV),
			}
		},
	}
)

func New() *Record {
	rec := _poolRecord.Get().(*Record)
	rec.Timestamp = time.Now()

	return rec
}

func (rec *Record) Free() {
	rec.KeyValues = rec.KeyValues[:0]
	_poolRecord.Put(rec)
}

/*
####### END ############################################################################################################
*/
