/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package cmdline

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/archnum/sdk.base/application"
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/tracer"
)

type (
	Command struct {
		app         *application.Application
		name        string
		description string
		handler     func() error
		flagSet     *flag.FlagSet
		flags       []*flagVar
		subCmds     map[string]*Command
	}
)

func NewCommand(name, description string, handler func() error) *Command {
	return &Command{
		name:        name,
		description: description,
		handler:     handler,
		flagSet:     flag.NewFlagSet(name, flag.ContinueOnError),
		subCmds:     make(map[string]*Command),
	}
}

func (cmd *Command) Add(cmds ...*Command) {
	for _, c := range cmds {
		c.app = cmd.app
		cmd.subCmds[c.name] = c
	}
}

func (cmd *Command) setFlags() error {
	for _, flag := range cmd.flags {
		if err := flag.presetFlagValue(); err != nil {
			return err
		}

		flag.setFlag()
	}

	return nil
}

func (cmd *Command) help() {
	const size = 30
	fullname := cmd.app.FullName()

	flag := cmd.flagSet
	flag.SetOutput(os.Stdout)

	fmt.Println(strings.Repeat("=", size))

	if cmd.name == fullname {
		fmt.Println(fullname)
		fmt.Println("\nFLAGS:")
		flag.PrintDefaults()
		fmt.Println("  -help")
		fmt.Println("  -version")
	} else {
		fmt.Println(fullname, "...", cmd.name)
		fmt.Println("\nFLAGS:")
		flag.PrintDefaults()
		fmt.Println("  -help")
	}

	if len(cmd.subCmds) > 0 {
		fmt.Println("\nSUBCOMMANDS:")

		for name, sub := range cmd.subCmds {
			fmt.Printf("   %s\n", name)
			fmt.Printf("        %s\n", sub.description)
		}
	}

	fmt.Println(strings.Repeat("=", size))
}

func (cmd *Command) Run(args []string) error {
	tracer.Log("CmdLine", kv.String("cmd", cmd.name)) //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	if err := cmd.setFlags(); err != nil {
		return err
	}

	if len(args) > 0 {
		arg := args[0]

		if arg == "help" || arg == "-help" || arg == "--help" {
			cmd.help()
			return failure.NoError
		}

		flag := cmd.flagSet
		flag.SetOutput(io.Discard)

		if err := flag.Parse(args); err != nil {
			return err
		}

		args = flag.Args()
	}

	if len(args) == 0 {
		if cmd.handler == nil {
			if cmd.name == cmd.app.FullName() {
				return nil
			}

			return failure.New( ////////////////////////////////////////////////////////////////////////////////////////
				"no handler provided for this command",
				kv.String("name", cmd.name),
			)
		}

		return cmd.handler()
	}

	name, args := args[0], args[1:]

	sub, ok := cmd.subCmds[name]
	if ok {
		return sub.Run(args)
	}

	return failure.New("this command doesn't exist", kv.String("name", name)) //////////////////////////////////////////
}

/*
####### END ############################################################################################################
*/
