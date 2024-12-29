/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package buffer

import (
	"fmt"
	"strconv"
	"time"
)

type (
	Buffer struct {
		bs   []byte
		pool *Pool
	}
)

func (b *Buffer) Append(values ...any) {
	b.bs = fmt.Append(b.bs, values...)
}

func (b *Buffer) AppendBool(value bool) {
	b.bs = strconv.AppendBool(b.bs, value)
}

func (b *Buffer) AppendByte(value byte) {
	b.bs = append(b.bs, value)
}

func (b *Buffer) AppendFloat(value float64, fmt byte, prec, bitSize int) {
	b.bs = strconv.AppendFloat(b.bs, value, fmt, prec, bitSize)
}

func (b *Buffer) AppendInt(value int, base int) {
	b.bs = strconv.AppendInt(b.bs, int64(value), base)
}

func (b *Buffer) AppendInt64(value int64, base int) {
	b.bs = strconv.AppendInt(b.bs, value, base)
}

func (b *Buffer) AppendUint(value uint64, base int) {
	b.bs = strconv.AppendUint(b.bs, value, base)
}

func (b *Buffer) AppendString(value string) {
	b.bs = append(b.bs, value...)
}

func (b *Buffer) AppendALString(value string, v byte, size int) {
	if len(value) < size {
		b.AppendString(value)

		for range size - len(value) {
			b.AppendByte(v)
		}
	} else {
		b.AppendString(value[:size])
	}
}

func (b *Buffer) AppendARString(value string, v byte, size int) {
	if len(value) < size {
		for range size - len(value) {
			b.AppendByte(v)
		}

		b.AppendString(value)
	} else {
		b.AppendString(value[:size])
	}
}

func (b *Buffer) AppendQuotedString(value string) {
	b.bs = strconv.AppendQuote(b.bs, value)
}

func (b *Buffer) AppendTime(value time.Time, layout string) {
	b.bs = value.AppendFormat(b.bs, layout)
}

func (b *Buffer) Write(value []byte) (int, error) {
	b.bs = append(b.bs, value...)
	return len(value), nil
}

func (b *Buffer) WriteByte(value byte) error {
	b.AppendByte(value)
	return nil
}

func (b *Buffer) WriteString(value string) (int, error) {
	b.AppendString(value)
	return len(value), nil
}

func (b *Buffer) Len() int {
	return len(b.bs)
}

func (b *Buffer) Cap() int {
	return cap(b.bs)
}

func (b *Buffer) Bytes() []byte {
	return b.bs
}

func (b *Buffer) String() string {
	return string(b.bs)
}

func (b *Buffer) Reset() {
	b.bs = b.bs[:0]
}

func (b *Buffer) Free() {
	b.pool.put(b)
}

/*
####### END ############################################################################################################
*/
