/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package cmdline

import (
	"strings"

	"github.com/archnum/sdk.base/application"
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/tracer"
)

type (
	VersionPrinter func(app *application.Application)

	ConfigLoader func(app *application.Application, filepath string) error

	CmdLine struct {
		*Command
		app            *application.Application
		versionPrinter VersionPrinter
		configLoader   ConfigLoader
	}

	Option func(*CmdLine)
)

func WithVersionPrinter(printer VersionPrinter) Option {
	return func(cl *CmdLine) {
		cl.versionPrinter = printer
	}
}

func WithConfigLoader(loader ConfigLoader) Option {
	return func(cl *CmdLine) {
		cl.configLoader = loader
	}
}

func WithHandler(fn func() error) Option {
	return func(cl *CmdLine) {
		cl.Command.handler = fn
	}
}

func New(app *application.Application, opts ...Option) (*CmdLine, error) {
	if app == nil {
		return nil,
			failure.New("application not provided") ////////////////////////////////////////////////////////////////////
	}

	cmd := NewCommand(app.FullName(), app.FullName(), nil)
	cmd.app = app

	cl := &CmdLine{
		Command: cmd,
		app:     app,
	}

	for _, option := range opts {
		option(cl)
	}

	return cl, nil
}

func (cl *CmdLine) printVersion() {
	if cl.versionPrinter != nil {
		cl.versionPrinter(cl.app)
	} else {
		cl.app.PrintVersion()
	}
}

func (cl *CmdLine) searchConfigFile(args []string) (string, []string, error) {
	for n, arg := range args {
		if arg == "-config-file" || arg == "--config-file" {
			if len(args) < n+2 {
				return "", nil,
					failure.New("configuration file name not provided") ////////////////////////////////////////////////
			}

			return args[n+1], append(args[:n], args[n+2:]...), nil
		}

		if strings.HasPrefix(arg, "-config-file") || strings.HasPrefix(arg, "--config-file") {
			values := strings.Split(arg, "=")
			if len(values) != 2 || values[1] == "" {
				return "", nil,
					failure.New("configuration file name not provided") ////////////////////////////////////////////////
			}

			return values[1], append(args[:n], args[n+1:]...), nil
		}
	}

	return cl.app.ConfigFile(), args, nil
}

func (cl *CmdLine) loadConfig(args []string) ([]string, error) {
	filepath, rArgs, err := cl.searchConfigFile(args)
	if err != nil {
		return nil, err
	}

	if filepath == "" {
		return rArgs, nil
	}

	tracer.Log("CmdLine", kv.String("config_file", filepath)) //::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	if err := cl.configLoader(cl.app, filepath); err != nil {
		return nil, err
	}

	return rArgs, nil
}

func (cl *CmdLine) Run(args []string) error {
	tracer.Log("CmdLine") //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	if len(args) > 0 {
		arg := args[0]

		if arg == "version" || arg == "-version" || arg == "--version" {
			cl.printVersion()
			return failure.NoError
		}
	}

	if cl.configLoader != nil {
		rArgs, err := cl.loadConfig(args)
		if err != nil {
			return err
		}

		args = rArgs
	}

	return cl.Command.Run(args)
}

/*
####### END ############################################################################################################
*/
