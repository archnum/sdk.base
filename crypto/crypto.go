/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"

	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
)

type (
	Crypto interface {
		SetKey(key string) error
		Encrypt(data []byte) ([]byte, error)
		EncryptString(s string) (string, error)
		Decrypt(data []byte) ([]byte, error)
		DecryptString(s string) (string, error)
	}

	implCrypto struct {
		key []byte
	}
)

func New() Crypto {
	return &implCrypto{
		key: []byte{
			0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF,
			0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF,
		},
	}
}

func (impl *implCrypto) SetKey(key string) error {
	hasher := sha256.New()

	if _, err := hasher.Write([]byte(key)); err != nil {
		return err
	}

	impl.key = hasher.Sum(nil)

	return nil
}

func (impl *implCrypto) Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(impl.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

func (impl *implCrypto) EncryptString(s string) (string, error) {
	data, err := impl.Encrypt([]byte(s))
	if err != nil {
		return "",
			failure.WithMessage(err, "failed to encrypt the string") ///////////////////////////////////////////////////
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

func (impl *implCrypto) Decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(impl.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()

	if nonceSize > len(data) {
		return nil,
			failure.New("not enough data") /////////////////////////////////////////////////////////////////////////////
	}

	nonce, cipherData := data[:nonceSize], data[nonceSize:]

	plainData, err := gcm.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return nil, err
	}

	return plainData, nil
}

func (impl *implCrypto) DecryptString(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "",
			failure.WithMessage( ///////////////////////////////////////////////////////////////////////////////////////
				err,
				"failed to decode this string",
				kv.String("string", s),
			)
	}

	data, err := impl.Decrypt(decoded)
	if err != nil {
		return "",
			failure.WithMessage( ///////////////////////////////////////////////////////////////////////////////////////
				err,
				"failed to decrypt this string",
				kv.String("string", s),
			)
	}

	return string(data), nil
}

/*
####### END ############################################################################################################
*/
