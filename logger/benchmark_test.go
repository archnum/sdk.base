/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package logger

import (
	"context"
	"io"
	"log/slog"
	"testing"

	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger/formatter/logfmt"
	"github.com/archnum/sdk.base/logger/handler/writer"
	"github.com/archnum/sdk.base/logger/level"
	"github.com/archnum/sdk.base/uuid"
)

func Benchmark_Log(b *testing.B) {
	l := New(uuid.Zero, "test")
	l.AddHandler(writer.NewHandler("", level.Info, logfmt.NewFormatter(nil), io.Discard))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.Info(
			"Ceci n'est qu'un message de test",
			kv.Bool("bool", false),
			kv.Int("int", 123),
			kv.String("string", "abcdefghijklmnopqrstuvwxyz"),
		)
	}
}

func Benchmark_Slog(b *testing.B) {
	l := slog.New(slog.NewTextHandler(io.Discard, nil))

	ctx := context.Background()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.LogAttrs(
			ctx,
			slog.LevelInfo,
			"Ceci n'est qu'un message de test",
			slog.Bool("bool", false),
			slog.Int("int", 123),
			slog.String("string", "abcdefghijklmnopqrstuvwxyz"),
		)
	}

}

/*
####### END ############################################################################################################
*/
