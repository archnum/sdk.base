/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package yaml

import "gopkg.in/yaml.v3"

type (
	YAML struct{}
)

func (p *YAML) Parse(bs []byte) (map[string]any, error) {
	var msa map[string]any

	if err := yaml.Unmarshal(bs, &msa); err != nil {
		return nil, err
	}

	return msa, nil
}

/*
####### END ############################################################################################################
*/
