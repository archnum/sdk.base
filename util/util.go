/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package util

import (
	"fmt"
	"os"
	"strings"
)

func CleanString(s string) string {
	s = strings.ReplaceAll(s, "\n", "|")
	s = strings.ReplaceAll(s, "\t", " ")

	return s
}

func FileExist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}

	return b
}

func LookupEnv(key string, prefixes ...string) (string, bool) {
	var (
		value string
		found bool
	)

	for _, prefix := range prefixes {
		if tmp, ok := os.LookupEnv(strings.ToUpper(fmt.Sprintf("%s_%s", prefix, key))); ok {
			value = tmp
			found = true
		}
	}

	return value, found
}

/*
####### END ############################################################################################################
*/
