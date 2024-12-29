/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package gotracker

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger"
	"github.com/archnum/sdk.base/util"
)

var (
	_counter atomic.Uint64
)

type (
	GoTracker struct {
		name      string
		logger    *logger.Logger
		waitGroup sync.WaitGroup
		stopCh    chan struct{}
	}

	GoFunc func(chan struct{})

	Option func(*GoTracker)
)

func WithName(name string) Option {
	return func(gt *GoTracker) {
		gt.name = name
	}
}

func WithLogger(logger *logger.Logger) Option {
	return func(gt *GoTracker) {
		gt.logger = logger
	}
}

func New(opts ...Option) *GoTracker {
	gt := &GoTracker{
		stopCh: make(chan struct{}),
	}

	for _, option := range opts {
		option(gt)
	}

	return gt
}

func (gt *GoTracker) getName(name string) string {
	return strings.Join([]string{fmt.Sprintf("%03d", _counter.Add(1)), gt.name, name}, ".")
}

func (gt *GoTracker) Go(name string, fn GoFunc) {
	name = gt.getName(name)

	gt.waitGroup.Add(1)

	go func() { //@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		defer gt.waitGroup.Done()

		if gt.logger != nil {
			gt.logger.Trace("GoTracker started", kv.String("name", name)) //::::::::::::::::::::::::::::::::::::::::::::

			defer func() {
				gt.logger.Trace("GoTracker stopped", kv.String("name", name)) //::::::::::::::::::::::::::::::::::::::::
			}()
		}

		defer func() {
			if data := recover(); data != nil {
				if gt.logger != nil {
					gt.logger.Error( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
						"GoTracker error recovered",
						kv.String("name", gt.name),
						kv.Any("data", data),
						kv.String("stack", util.Stack(5)),
					)
				}
			}
		}()

		fn(gt.stopCh)
	}()
}

func (gt *GoTracker) Stop() {
	close(gt.stopCh)
}

func (gt *GoTracker) Wait() {
	gt.waitGroup.Wait()
}

/*
####### END ############################################################################################################
*/
