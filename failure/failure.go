/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package failure

import (
	"github.com/archnum/sdk.base/buffer"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/util"
)

var (
	_bufPool = buffer.NewPool(256)
)

func format(buf *buffer.Buffer, msg string, kvs []kv.KeyValue) {
	buf.AppendString(msg)

	if len(kvs) > 0 {
		buf.AppendString(": ")
		kv.Logfmt(buf, kvs)
	}
}

type (
	base struct {
		message string
	}
)

func New(msg string, kvs ...kv.KeyValue) error {
	buf := _bufPool.Get()
	defer buf.Free()

	format(buf, msg, kvs)

	return &base{message: buf.String()}
}

func (b *base) Error() string {
	return b.message
}

type (
	withMessage struct {
		cause   error
		message string
	}
)

func WithMessage(err error, msg string, kvs ...kv.KeyValue) error {
	if err == nil {
		return nil
	}

	buf := _bufPool.Get()
	defer buf.Free()

	format(buf, msg, kvs)

	buf.AppendString(" >> ")
	buf.AppendString(util.CleanString(err.Error()))

	return &withMessage{
		cause:   err,
		message: buf.String(),
	}
}

func (wm *withMessage) Error() string {
	return wm.message
}

func (wm *withMessage) Unwrap() error {
	return wm.cause
}

func Wrap(cause, err error) error {
	return WithMessage(cause, err.Error())
}

/*
####### END ############################################################################################################
*/
