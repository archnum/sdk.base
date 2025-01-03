/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package logger

import "github.com/archnum/sdk.base/kv"

func ArgsToKV(args ...any) []kv.KeyValue {
	if len(args)%2 != 0 {
		return []kv.KeyValue{kv.String("ERROR", "odd number of arguments")}
	}

	var kvs []kv.KeyValue

	for i := 0; i < len(args); i += 2 {
		key, ok := args[i].(string)
		if !ok {
			key = "?"
		}

		kvs = append(kvs, kv.Any(key, args[i+1]))
	}

	return kvs
}

/*
####### END ############################################################################################################
*/
