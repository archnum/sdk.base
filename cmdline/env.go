/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package cmdline

import (
	"strconv"

	"github.com/archnum/sdk.base/application"
	"github.com/archnum/sdk.base/failure"
)

type (
	Loader[T any] func(app *application.Application, flag string) (T, error)
)

func Env[T any](fn func(app *application.Application, env string) (T, error), envVars ...string) []Loader[T] {
	var loaders []Loader[T]

	for _, env := range envVars {
		loaders = append(
			loaders,
			func(app *application.Application, _ string) (T, error) {
				return fn(app, env)
			},
		)
	}

	return loaders
}

func EnvBool(envVars ...string) []Loader[bool] {
	return Env(
		func(app *application.Application, env string) (bool, error) {
			if str, ok := app.LookupEnv(env); ok {
				return strconv.ParseBool(str)
			}

			return false, failure.NoError
		},
		envVars...,
	)
}

func EnvString(envVars ...string) []Loader[string] {
	return Env(
		func(app *application.Application, env string) (string, error) {
			if value, ok := app.LookupEnv(env); ok {
				return value, nil
			}

			return "", failure.NoError
		},
		envVars...,
	)
}

/*
####### END ############################################################################################################
*/
