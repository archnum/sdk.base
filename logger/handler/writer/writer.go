/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package writer

import (
	"io"

	"github.com/archnum/sdk.base/logger/formatter"
	"github.com/archnum/sdk.base/logger/level"
)

type (
	implHandler struct {
		name      string
		level     level.Var
		formatter formatter.Formatter
		writer    io.Writer
	}
)

func NewHandler(name string, level level.Level, f formatter.Formatter, w io.Writer) *implHandler {
	impl := &implHandler{
		name:      name,
		formatter: f,
		writer:    w,
	}

	impl.level.Set(level)

	return impl
}

func (impl *implHandler) Name() string {
	return impl.name
}

func (impl *implHandler) Level() level.Level {
	return impl.level.Level()
}

func (impl *implHandler) SetLevel(lvl level.Level) {
	impl.level.Set(lvl)
}

func (impl *implHandler) Formatter() formatter.Formatter {
	return impl.formatter
}

func (impl *implHandler) Log(_ level.Level, msg []byte) error {
	_, err := impl.writer.Write(msg)
	return err
}

/*
####### END ############################################################################################################
*/
