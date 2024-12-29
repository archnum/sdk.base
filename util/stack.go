/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package util

import (
	"runtime/debug"
	"slices"
	"strings"
)

func Stack(n int) string {
	stack := strings.Split(string(debug.Stack()), "\n")

	tmp := []string{}

	for i := len(stack) - 1; i > 0; i-- {
		tmp = append(tmp, stack[i])

		if strings.HasPrefix(stack[i], "panic(") {
			tmp = tmp[0 : len(tmp)-2]
			break
		}
	}

	stack = stack[:0]

	for _, s := range tmp {
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

	n = If(len(stack) < n, len(stack), n)

	return strings.Join(stack[:n], " | ")
}

/*
####### END ############################################################################################################
*/
