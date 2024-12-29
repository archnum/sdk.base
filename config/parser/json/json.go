/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package json

import "encoding/json"

type (
	JSON struct{}
)

func (p *JSON) Parse(bs []byte) (map[string]any, error) {
	var msa map[string]any

	if err := json.Unmarshal(bs, &msa); err != nil {
		return nil, err
	}

	return msa, nil
}

/*
####### END ############################################################################################################
*/
