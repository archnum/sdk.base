/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package cmdline

import (
	"github.com/archnum/sdk.base/application"
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
)

type (
	flagVar struct {
		presetFlagValue func() error
		setFlag         func()
	}
)

func presetFlagValue[T any](app *application.Application, name string, p *T, loaders []Loader[T]) error {
	for _, loader := range loaders {
		if value, err := loader(app, name); err != nil {
			if err != failure.NoError {
				return failure.WithMessage( ////////////////////////////////////////////////////////////////////////////
					err,
					"flag preset error",
					kv.String("name", name),
				)
			}
		} else {
			*p = value
		}
	}

	return nil
}

func (cmd *Command) BoolVar(p *bool, name string, usage string, loaders ...Loader[bool]) {
	cmd.flags = append(
		cmd.flags,
		&flagVar{
			presetFlagValue: func() error {
				return presetFlagValue(cmd.app, name, p, loaders)
			},
			setFlag: func() {
				cmd.flagSet.BoolVar(p, name, *p, usage)
			},
		},
	)
}

func (cmd *Command) StringVar(p *string, name string, usage string, loaders ...Loader[string]) {
	cmd.flags = append(
		cmd.flags,
		&flagVar{
			presetFlagValue: func() error {
				return presetFlagValue(cmd.app, name, p, loaders)
			},
			setFlag: func() {
				cmd.flagSet.StringVar(p, name, *p, usage)
			},
		},
	)
}

/*
####### END ############################################################################################################
*/
