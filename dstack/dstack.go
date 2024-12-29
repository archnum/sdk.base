/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package dstack

import (
	"runtime/debug"
	"slices"
	"strings"

	"github.com/archnum/sdk.base/util"
)

func String(n int) string {
	stack := strings.Split(string(debug.Stack()), "\n")

	ss := []string{}

	for i := len(stack) - 1; i > 0; i-- {
		ss = append(ss, stack[i])

		if strings.HasPrefix(stack[i], "panic(") {
			ss = ss[0 : len(ss)-2]
			break
		}
	}

	stack = stack[:0]

	for _, s := range ss {
		if strings.Contains(s, ".go:") {
			s = strings.TrimSpace(s)

			if idx := strings.LastIndex(s, " "); idx < 0 {
				stack = append(stack, s)
			} else {
				stack = append(stack, s[:idx])
			}
		}
	}

	slices.Reverse(stack)

	n = util.If(len(stack) < n, len(stack), n)

	return strings.Join(stack[:n], " | ")
}

/*
####### END ############################################################################################################
*/
