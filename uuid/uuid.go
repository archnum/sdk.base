/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"regexp"

	"github.com/archnum/sdk.base/failure"
)

const (
	Zero = UUID("00000000-0000-0000-0000-000000000000")
)

type (
	UUID string
)

var (
	_uuidRE = regexp.MustCompile(`^[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{12}$`)
)

func Validate(id UUID) bool {
	return _uuidRE.MatchString(string(id))
}

func ConvertString(value string) (UUID, bool) {
	id := UUID(value)
	return id, Validate(id)
}

func generate(reader io.Reader) (UUID, error) {
	var (
		bs [16]byte
		bd [36]byte
	)

	if _, err := io.ReadFull(reader, bs[:]); err != nil {
		return "",
			failure.WithMessage(err, "failed to generate a UUID") //////////////////////////////////////////////////////
	}

	hex.Encode(bd[:8], bs[:4])
	hex.Encode(bd[9:13], bs[4:6])
	hex.Encode(bd[14:18], bs[6:8])
	hex.Encode(bd[19:23], bs[8:10])
	hex.Encode(bd[24:], bs[10:])

	bd[8] = '-'
	bd[13] = '-'
	bd[18] = '-'
	bd[23] = '-'

	return UUID(bd[:]), nil
}

func New() (UUID, error) {
	return generate(rand.Reader)
}

func String() (string, error) {
	id, err := New()
	return string(id), err
}

/*
####### END ############################################################################################################
*/
