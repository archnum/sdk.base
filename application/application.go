/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package application

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/util"
	"github.com/archnum/sdk.base/uuid"
)

const (
	_defaultVersion = "0.0.0"
	_defaultDesc    = "?"
)

type (
	Application struct {
		id          uuid.UUID
		name        string
		ecosystem   string
		version     string
		builtAt     time.Time
		shortDesc   string
		longDesc    string
		startedAt   time.Time
		configFile  string
		environment string
	}

	Option func(*Application) error
)

func WithEcosystem(name string) Option {
	return func(app *Application) error {
		app.ecosystem = name
		return nil
	}
}

func WithVersion(version string) Option {
	return func(app *Application) error {
		app.version = version
		return nil
	}
}

func WithBuiltAt(builtAt string) Option {
	return func(app *Application) error {
		seconds, err := strconv.ParseInt(builtAt, 0, 64)
		if err != nil {
			return err
		}

		app.builtAt = time.Unix(seconds, 0)

		return nil
	}
}

func WithShortDesc(desc string) Option {
	return func(app *Application) error {
		app.shortDesc = desc
		return nil
	}
}

func WithLongDesc(desc string) Option {
	return func(app *Application) error {
		app.longDesc = desc
		return nil
	}
}

func WithConfigFile(path string) Option {
	return func(app *Application) error {
		app.configFile = path
		return nil
	}
}

func configure(app *Application) {
	if value, ok := util.LookupEnv("ecosystem", app.name); ok {
		app.ecosystem = value
	}

	if value, ok := app.LookupEnv("config_file"); ok {
		app.configFile = value
	}

	if value, ok := app.LookupEnv("environment"); ok {
		app.environment = value
	}
}

func New(name string, opts ...Option) (*Application, error) {
	id, err := uuid.New()
	if err != nil {
		return nil, err
	}

	app := &Application{
		id:        id,
		name:      name,
		version:   _defaultVersion,
		builtAt:   time.Time{},
		shortDesc: _defaultDesc,
		longDesc:  _defaultDesc,
		startedAt: time.Now(),
	}

	for _, option := range opts {
		if err := option(app); err != nil {
			return nil,
				failure.WithMessage( ///////////////////////////////////////////////////////////////////////////////////
					err,
					"failed to create this application",
					kv.String("name", name),
				)
		}
	}

	configure(app)

	return app, nil
}

func (app *Application) ID() uuid.UUID {
	return app.id
}

func (app *Application) Name() string {
	return app.name
}

func (app *Application) Ecosystem() string {
	return app.ecosystem
}

func (app *Application) FullName() string {
	if app.ecosystem == "" {
		return app.name
	}

	return app.ecosystem + "." + app.name
}

func (app *Application) Version() string {
	return app.version
}

func (app *Application) BuiltAt() time.Time {
	return app.builtAt
}

func (app *Application) ShortDesc() string {
	return app.shortDesc
}

func (app *Application) LongDesc() string {
	return app.longDesc
}

func (app *Application) StartedAt() time.Time {
	return app.startedAt
}

func (app *Application) ConfigFile() string {
	return app.configFile
}

func (app *Application) Environment() string {
	return app.environment
}

func (app *Application) LookupEnv(key string) (string, bool) {
	prefixes := []string{app.name}

	if app.ecosystem != "" {
		prefixes = append(prefixes, app.ecosystem)
	}

	return util.LookupEnv(key, prefixes...)
}

func (app *Application) PrintVersion() {
	fullname := app.FullName()

	msg := fmt.Sprintf(
		"%s v%s %s %s",
		fullname,
		app.Version(),
		app.BuiltAt().Format(time.DateTime),
		runtime.Version(),
	)

	fmt.Printf(
		"\n%s\n%s (c) 2024 Archivage Numérique %s\n\n",
		msg,
		strings.Repeat("=", len(fullname)),
		strings.Repeat("=", len(msg)-len(fullname)-30),
	)
}

func (app *Application) Exit(err error) {
	exitCode := 0

	if err != nil {
		if err == failure.NoError {
			exitCode = -1
		} else {
			fmt.Fprintf( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
				os.Stderr,
				"ERROR - application %s - %s\n",
				app.FullName(),
				err.Error(),
			)

			exitCode = -2
		}
	}

	os.Exit(exitCode)
}

/*
####### END ############################################################################################################
*/
