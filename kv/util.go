/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package kv

import (
	"unicode"
	"unicode/utf8"
)

func needsQuoting(s string) bool {
	if len(s) == 0 {
		return true
	}

	for i := 0; i < len(s); {
		if b := s[i]; b < utf8.RuneSelf {
			if b != '\\' && (b == ' ' || b == '=' || b == '"') {
				return true
			}

			i++
			continue
		}

		r, size := utf8.DecodeRuneInString(s[i:])

		if r == utf8.RuneError || unicode.IsSpace(r) || !unicode.IsPrint(r) {
			return true
		}

		i += size
	}

	return false
}

/*
####### END ############################################################################################################
*/
