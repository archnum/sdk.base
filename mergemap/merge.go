/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package mergemap

import (
	"reflect"

	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
)

const (
	_maxDepth = 10
)

type (
	config struct {
		maxDepth int
	}

	Option func(*config)
)

func WithMaxDepth(maxDepth int) Option {
	return func(c *config) {
		c.maxDepth = maxDepth
	}
}

func toMap(in any) (map[string]any, bool) {
	value := reflect.ValueOf(in)
	if value.Kind() == reflect.Map {
		out := make(map[string]any)

		for _, k := range value.MapKeys() {
			out[k.String()] = value.MapIndex(k).Interface()
		}

		return out, true
	}

	return nil, false
}

func merge(c *config, dst, src map[string]any, depth int) (map[string]any, error) {
	if c.maxDepth > 0 && depth > c.maxDepth {
		return dst,
			failure.New( ///////////////////////////////////////////////////////////////////////////////////////////////
				"the merger is too deep",
				kv.Int("depth", depth),
				kv.Int("max", c.maxDepth),
			)
	}

	var err error

	for key, srcValue := range src {
		if dstValue, ok := dst[key]; ok {
			srcMap, srcOk := toMap(srcValue)
			dstMap, dstOk := toMap(dstValue)

			if srcOk && dstOk {
				srcValue, err = merge(c, dstMap, srcMap, depth+1)
				if err != nil {
					return dst, err
				}
			}
		}

		dst[key] = srcValue
	}

	return dst, err
}

func Merge(dst, src map[string]any, opts ...Option) error {
	cfg := &config{
		maxDepth: _maxDepth,
	}

	for _, option := range opts {
		option(cfg)
	}

	_, err := merge(cfg, dst, src, 0)

	return err
}

/*
####### END ############################################################################################################
*/
