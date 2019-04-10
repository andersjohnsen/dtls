package dtls

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"

	"github.com/pions/dtls/internal/crypto/ccm"
)

// State needed to handle encrypted input/output
type cryptoCCM struct {
	ccm cipher.AEAD
}

// nonce creates a nonce
func nonce(iv, data []byte) []byte {
	n := new(bytes.Buffer)
	n.Write(iv)
	n.Write(data)
	return n.Bytes()
}

// newCryptoCCM8 creates a cryptoCCM with the tag size set to 8
func newCryptoCCM8(key []byte, nonceSize int) (*cryptoCCM, error) {
	return newCryptoCCM(key, 8, nonceSize)
}

// newCryptoCCM creates a cryptoCCM, where M is tagSize and L is nonceSize
func newCryptoCCM(key []byte, tagSize, nonceSize int) (*cryptoCCM, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ccm, err := ccm.NewCCM(block, tagSize, nonceSize)
	if err != nil {
		return nil, err
	}

	return &cryptoCCM{
		ccm: ccm,
	}, nil
}

func (c *cryptoCCM) encrypt(pkt *recordLayer, raw []byte) ([]byte, error) {
	// https://github.com/bocajim/dtls/blob/master/crypto.go#L128
	return nil, nil

}

func (c *cryptoCCM) decrypt(in []byte) ([]byte, error) {
	// https://github.com/bocajim/dtls/blob/master/crypto.go#L160
	return nil, nil
}
