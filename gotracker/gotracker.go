/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package gotracker

import (
	"context"
	"errors"
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
		errs      error
		ctx       context.Context
		logger    *logger.Logger
		cancel    context.CancelFunc
		name      string
		waitGroup sync.WaitGroup
	}

	GoFunc func(ctx context.Context) error

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
	gt := &GoTracker{}

	for _, option := range opts {
		option(gt)
	}

	gt.ctx, gt.cancel = context.WithCancel(context.Background())

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

		gt.errs = errors.Join(gt.errs, fn(gt.ctx))
	}()
}

func (gt *GoTracker) Done() <-chan struct{} {
	return gt.ctx.Done()
}

func (gt *GoTracker) Stop() {
	gt.cancel()
}

func (gt *GoTracker) Wait() {
	gt.waitGroup.Wait()
}

func (gt *GoTracker) Err() error {
	return gt.errs
}

/*
####### END ############################################################################################################
*/
