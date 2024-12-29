/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package failure

import (
	"github.com/archnum/sdk.base/buffer"
	"github.com/archnum/sdk.base/kv"
)

var (
	_bufPool = buffer.NewPool(256)
)

type (
	Failure struct {
		Cause   error
		Message string
	}
)

func format(buf *buffer.Buffer, msg string, kvs []kv.KeyValue) {
	buf.AppendString(msg)

	if len(kvs) > 0 {
		buf.AppendString(": ")
		kv.Logfmt(buf, kvs)
	}
}

func New(msg string, kvs ...kv.KeyValue) *Failure {
	buf := _bufPool.Get()
	defer buf.Free()

	format(buf, msg, kvs)

	return &Failure{Message: buf.String()}
}

func WithMessage(cause error, msg string, kvs ...kv.KeyValue) *Failure {
	if cause == nil {
		return New(msg, kvs...)
	}

	buf := _bufPool.Get()
	defer buf.Free()

	format(buf, msg, kvs)

	buf.AppendString(" >> ")
	buf.AppendString(cause.Error())

	return &Failure{
		Cause:   cause,
		Message: buf.String(),
	}
}

func Wrap(cause, err error) error {
	if cause == nil {
		return err
	}

	return &Failure{
		Cause:   cause,
		Message: err.Error(),
	}
}

func (f *Failure) Error() string {
	return f.Message
}

func (f *Failure) Unwrap() error {
	return f.Cause
}

/*
####### END ############################################################################################################
*/
