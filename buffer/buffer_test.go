/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package buffer

import (
	"slices"
	"testing"
	"time"
)

func Test(t *testing.T) {
	const size = 32

	buf := NewPool(size).Get()
	defer buf.Free()

	tests := []struct {
		desc string
		fn   func()
		want string
	}{
		{"Append", func() { buf.Append(0, "zero", false) }, "0zerofalse"},
		{"AppendBool(false)", func() { buf.AppendBool(false) }, "false"},
		{"AppendBool(true)", func() { buf.AppendBool(true) }, "true"},
		{"AppendByte", func() { buf.AppendByte('#') }, "#"},
		{"AppendFloat", func() { buf.AppendFloat(6.55957, 'E', -1, 64) }, "6.55957E+00"},
		{"AppendInt(>0)", func() { buf.AppendInt(23, 10) }, "23"},
		{"AppendInt(<0)", func() { buf.AppendInt(-7, 10) }, "-7"},
		{"AppendInt64(>0)", func() { buf.AppendInt64(23, 10) }, "23"},
		{"AppendInt64(<0)", func() { buf.AppendInt64(-7, 10) }, "-7"},
		{"AppendUint", func() { buf.AppendUint(16, 16) }, "10"},
		{"AppendString", func() { buf.AppendString("BZH") }, "BZH"},
		{"AppendALString(<)", func() { buf.AppendALString("12345", '-', 8) }, "12345---"},
		{"AppendALString(>)", func() { buf.AppendALString("12345", '-', 4) }, "1234"},
		{"AppendARString(<)", func() { buf.AppendARString("12345", '-', 8) }, "---12345"},
		{"AppendARString(>)", func() { buf.AppendARString("12345", '-', 4) }, "1234"},
		{"AppendQuotedString", func() { buf.AppendQuotedString("Hello World!") }, `"Hello World!"`},
		{
			"AppendTime",
			func() {
				buf.AppendTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), time.DateTime)
			},
			"2009-11-10 23:00:00",
		},
		{"Write", func() { _, _ = buf.Write([]byte{65, 66, 67}) }, "ABC"},
		{"WriteByte", func() { _ = buf.WriteByte(100) }, "d"},
		{"WriteString", func() { _, _ = buf.WriteString("The Lord of the Rings") }, "The Lord of the Rings"},
	}

	for _, tt := range tests {
		t.Run(
			tt.desc,
			func(t *testing.T) {
				tt.fn()

				if v := buf.String(); v != tt.want {
					t.Errorf("got %q, want %q", v, tt.want) //..........................................................
				}

				if slices.Compare(buf.Bytes(), []byte(tt.want)) != 0 {
					t.Errorf("got %q, want %q", buf.Bytes(), []byte(tt.want)) //........................................
				}

				if buf.Len() != len(tt.want) {
					t.Error("unexpected buffer length") //..............................................................
				}

				buf.Reset()

				if buf.Cap() != size {
					t.Error("buffer capacity has changed") //...........................................................
				}
			},
		)
	}
}

/*
####### END ############################################################################################################
*/
