/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package journal

import (
	_jnl "github.com/coreos/go-systemd/v22/journal"

	"github.com/archnum/sdk.base/logger/formatter"
	"github.com/archnum/sdk.base/logger/formatter/journal"
	"github.com/archnum/sdk.base/logger/level"
)

type (
	implHandler struct {
		name      string
		level     level.Var
		formatter formatter.Formatter
	}
)

func NewHandler(name string, level level.Level, opts *journal.FormatterOptions) *implHandler {
	impl := &implHandler{
		name:      name,
		formatter: journal.NewFormatter(opts),
	}

	impl.level.Set(level)

	return impl
}

func (impl *implHandler) Name() string {
	return "journal"
}

func (impl *implHandler) Level() level.Level {
	return impl.level.Level()
}

func (impl *implHandler) SetLevel(lvl level.Level) {
	impl.level.Set(lvl)
}

func (impl *implHandler) Enabled(level level.Level) bool {
	return true
}

func (impl *implHandler) Formatter() formatter.Formatter {
	return impl.formatter
}

func (impl *implHandler) Log(lvl level.Level, msg []byte) error {
	switch lvl {
	case level.Info:
		return _jnl.Send(string(msg), _jnl.PriInfo, nil)
	case level.Notice:
		return _jnl.Send(string(msg), _jnl.PriNotice, nil)
	case level.Warning:
		return _jnl.Send(string(msg), _jnl.PriWarning, nil)
	case level.Error:
		return _jnl.Send(string(msg), _jnl.PriErr, nil)
	default:
		return _jnl.Send(string(msg), _jnl.PriDebug, nil)
	}
}

/*
####### END ############################################################################################################
*/
