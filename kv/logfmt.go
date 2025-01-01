/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package kv

import (
	"fmt"
	"time"

	"github.com/archnum/sdk.base/buffer"
)

const (
	DefaultTimeLayout = time.DateTime
)

var (
	_bufPool = buffer.NewPool(256)
)

func LogfmtAppendString(buf *buffer.Buffer, s string) {
	if needsQuoting(s) {
		buf.AppendQuotedString(s)
	} else {
		buf.AppendString(s)
	}
}

func Logfmt(buf *buffer.Buffer, kvs []KeyValue) {
	for i, kv := range kvs {
		if i > 0 {
			buf.AppendByte(' ')
		}

		// key
		LogfmtAppendString(buf, kv.Key)

		// =
		buf.AppendByte('=')

		// value
		value := kv.Value

		switch value.Kind() {
		case KindAny:
			buf.Append(value.any)

		case KindBool:
			buf.AppendBool(value.Bool())

		case KindDuration:
			buf.AppendString(value.Duration().String())

		case KindFloat:
			buf.AppendFloat(value.Float(), 'g', -1, 64)

		case KindInt:
			buf.AppendInt(int(value.num), 10)

		case KindInt64:
			buf.AppendInt64(int64(value.num), 10)

		case KindString:
			LogfmtAppendString(buf, value.String())

		case KindTime:
			LogfmtAppendString(buf, value.Time().Format(DefaultTimeLayout))

		case KindUint:
			buf.AppendUint(value.num, 10)

		default:
			panic(fmt.Sprintf("bad kind: %d", value.Kind())) ///////////////////////////////////////////////////////////
		}
	}
}

func LogfmtMessage(msg string, kvs []KeyValue) string {
	buf := _bufPool.Get()
	defer buf.Free()

	buf.AppendString(msg)

	if len(kvs) > 0 {
		buf.AppendString(": ")
		Logfmt(buf, kvs)
	}

	return buf.String()
}

/*
####### END ############################################################################################################
*/
