/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package mapstruct

import (
	"time"

	"github.com/go-viper/mapstructure/v2"

	"github.com/archnum/sdk.base/failure"
)

type (
	DecoderConfig = mapstructure.DecoderConfig
	Option        func(*DecoderConfig)
)

func WithDefaults() Option {
	return func(dc *DecoderConfig) {
		dc.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		)
		dc.ErrorUnused = false
		dc.Squash = true
		dc.TagName = "ms"
		dc.WeaklyTypedInput = true
	}
}

func WithTag(name string) Option {
	return func(dc *DecoderConfig) {
		dc.TagName = name
	}
}

func NewDecoderConfig(out any, opts ...Option) *DecoderConfig {
	dc := &DecoderConfig{
		Result: out,
	}

	for _, option := range opts {
		option(dc)
	}

	return dc
}

func DecodeWithConfig(dc *DecoderConfig, in any) error {
	decoder, err := mapstructure.NewDecoder(dc)
	if err != nil {
		return failure.WithMessage(err, "failed to create a decoder") //////////////////////////////////////////////////
	}

	err = decoder.Decode(in)
	if err == nil {
		return nil
	}

	return failure.WithMessage(err, "failed to decode data") ///////////////////////////////////////////////////////////
}

func Decode(out, in any) error {
	return DecodeWithConfig(NewDecoderConfig(out, WithDefaults()), in)
}

/*
####### END ############################################################################################################
*/
