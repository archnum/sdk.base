/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package config

import (
	"path/filepath"
	"slices"
	"strings"

	"github.com/archnum/sdk.base/config/parser/json"
	"github.com/archnum/sdk.base/config/parser/yaml"
	"github.com/archnum/sdk.base/config/reader/file"
	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/mapstruct"
	"github.com/archnum/sdk.base/mergemap"
)

type (
	Reader interface {
		Read() ([]byte, error)
	}

	Parser interface {
		Parse(bs []byte) (map[string]any, error)
	}

	Loader interface {
		Load() (map[string]any, error)
	}

	Config struct {
		data map[string]any
	}
)

func New() *Config {
	return &Config{
		data: make(map[string]any),
	}
}

func (cfg *Config) Merge(msa map[string]any) error {
	if msa == nil {
		return nil
	}

	return mergemap.Merge(cfg.data, msa)
}

func (cfg *Config) ReadAndParse(r Reader, p Parser) error {
	if r == nil {
		return failure.New("reader cannot be nil") /////////////////////////////////////////////////////////////////////
	}

	if p == nil {
		return failure.New("parser cannot be nil") /////////////////////////////////////////////////////////////////////
	}

	bs, err := r.Read()
	if err != nil {
		return failure.WithMessage(err, "failed to read") //////////////////////////////////////////////////////////////
	}

	msa, err := p.Parse(bs)
	if err != nil {
		return failure.WithMessage(err, "failed to parse") /////////////////////////////////////////////////////////////
	}

	return cfg.Merge(msa)
}

func (cfg *Config) Load(l Loader) error {
	if l == nil {
		return failure.New("loader cannot be nil") /////////////////////////////////////////////////////////////////////
	}

	msa, err := l.Load()
	if err != nil {
		return failure.WithMessage(err, "failed to load") //////////////////////////////////////////////////////////////
	}

	return cfg.Merge(msa)
}

func (cfg *Config) Get() map[string]any {
	return cfg.data
}

func (cfg *Config) DecodeWithConfig(dc *mapstruct.DecoderConfig) error {
	return mapstruct.DecodeWithConfig(dc, cfg.data)
}

func (cfg *Config) Decode(to any) error {
	return mapstruct.Decode(to, cfg.data)
}

type (
	FileParser interface {
		Parser
		Exts() []string
	}
)

func (cfg *Config) DecodeFile(to any, path string, fps ...FileParser) error {
	var parser Parser

	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".json":
		parser = &json.JSON{}
	case ".yml", ".yaml":
		parser = &yaml.YAML{}
	default:
		for _, fp := range fps {
			if slices.Contains(fp.Exts(), ext) {
				parser = fp
				break
			}
		}
	}

	if parser == nil {
		return failure.New("file extension unknown", kv.String("ext", ext)) ////////////////////////////////////////////
	}

	if err := cfg.ReadAndParse(&file.File{Path: path}, parser); err != nil {
		return err
	}

	return cfg.Decode(to)
}

/*
####### END ############################################################################################################
*/
