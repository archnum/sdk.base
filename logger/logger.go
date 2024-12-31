/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package logger

import (
	"log"
	"maps"
	"sync"

	"github.com/archnum/sdk.base/buffer"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger/handler"
	"github.com/archnum/sdk.base/logger/level"
	"github.com/archnum/sdk.base/logger/manager"
	"github.com/archnum/sdk.base/logger/record"
	"github.com/archnum/sdk.base/uuid"
)

var (
	_bufPool = buffer.NewPool(256)
)

type (
	// fieldalignment
	Logger struct {
		handlers map[string]handler.Handler
		id       uuid.UUID
		name     string
		with     []kv.KeyValue
		level    level.Var
		mutex    sync.Mutex
	}
)

func New(id uuid.UUID, name string) *Logger {
	// TODO: vérifer que id est bien un UUID ?

	if name == "" {
		name = "main"
	}

	return &Logger{id: id, name: name, handlers: make(map[string]handler.Handler, 0)}
}

func (l *Logger) Register() {
	manager.RegisterLogger(l)
}

func (l *Logger) ID() uuid.UUID {
	return l.id
}

func (l *Logger) Name() string {
	return l.name
}

func (l *Logger) Level() level.Level {
	return l.level.Level()
}

func (l *Logger) SetLevel(level level.Level) {
	l.level.Set(level)
}

func (l *Logger) AddHandler(h handler.Handler) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.handlers[h.Name()] = h
}

func (l *Logger) RemoveHandler(name string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	delete(l.handlers, name)
}

func (l *Logger) New(id uuid.UUID, name string) *Logger {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	nl := &Logger{id: id, name: name, handlers: maps.Clone(l.handlers)}
	nl.SetLevel(l.Level())

	return nl
}

func (l *Logger) With(kvs ...kv.KeyValue) *Logger {
	nl := l.New(l.id, l.name)
	nl.with = append(l.with, kvs...)

	return nl
}

func formatAndLog(h handler.Handler, buf *buffer.Buffer, rec *record.Record) {
	defer func() {
		if data := recover(); data != nil {
			_ = 0 // TODO
		}
	}()

	h.Formatter().Format(buf, rec)

	if err := h.Log(rec.Level, buf.Bytes()); err != nil {
		log.Print(err) //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	}
}

func (l *Logger) FormatAndLog(level level.Level, msg string, kvs ...kv.KeyValue) {
	rec := record.New()
	defer rec.Free()

	rec.LoggerID = string(l.id)
	rec.LoggerName = l.name
	rec.Level = level
	rec.Message = msg

	rec.KeyValues = append(rec.KeyValues, l.with...)
	rec.KeyValues = append(rec.KeyValues, kvs...)

	l.mutex.Lock()
	defer l.mutex.Unlock()

	for _, h := range l.handlers {
		if level < h.Level() {
			continue
		}

		buf := _bufPool.Get()
		formatAndLog(h, buf, rec)
		buf.Free()
	}
}

func (l *Logger) Log(level level.Level, msg string, kvs ...kv.KeyValue) {
	if level < l.level.Level() {
		return
	}

	l.FormatAndLog(level, msg, kvs...)
}

func (l *Logger) Trace(msg string, kvs ...kv.KeyValue) {
	if level.Trace < l.level.Level() {
		return
	}

	l.FormatAndLog(level.Trace, msg, kvs...)
}

func (l *Logger) Debug(msg string, kvs ...kv.KeyValue) {
	if level.Debug < l.level.Level() {
		return
	}

	l.FormatAndLog(level.Debug, msg, kvs...)
}

func (l *Logger) Info(msg string, kvs ...kv.KeyValue) {
	if level.Info < l.level.Level() {
		return
	}

	l.FormatAndLog(level.Info, msg, kvs...)
}

func (l *Logger) Notice(msg string, kvs ...kv.KeyValue) {
	if level.Notice < l.level.Level() {
		return
	}

	l.FormatAndLog(level.Notice, msg, kvs...)
}

func (l *Logger) Warning(msg string, kvs ...kv.KeyValue) {
	if level.Warning < l.level.Level() {
		return
	}

	l.FormatAndLog(level.Warning, msg, kvs...)
}

func (l *Logger) Error(msg string, kvs ...kv.KeyValue) {
	if level.Error < l.level.Level() {
		return
	}

	l.FormatAndLog(level.Error, msg, kvs...)
}

func (l *Logger) LogArgs(level level.Level, msg string, args ...any) {
	if level < l.level.Level() {
		return
	}

	if len(args)%2 != 0 {
		l.FormatAndLog(level, msg, kv.String("ERROR", "odd number of arguments"))
		return
	}

	var kvs []kv.KeyValue

	for i := 0; i < len(args); i += 2 {
		key, ok := args[i].(string)
		if !ok {
			key = "?"
		}

		kvs = append(kvs, kv.Any(key, args[i+1]))
	}

	l.FormatAndLog(level, msg, kvs...)
}

/*
####### END ############################################################################################################
*/
