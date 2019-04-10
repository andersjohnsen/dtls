package dtls

import (
	"errors"
)

type cipherSuiteTLSPskWithAes128Ccm struct {
	ccm *cryptoCCM
}

func (c cipherSuiteTLSPskWithAes128Ccm) certificateType() clientCertificateType {
	// TODO: ?? we don't have a certificate type for PSK
}

func (c cipherSuiteTLSPskWithAes128Ccm) ID() cipherSuiteID {
	return 0xc0a4
}

func (c *cipherSuiteTLSPskWithAes128Ccm) init(masterSecret, clientRandom, serverRandom []byte, isClient bool) error {
	// TODO
	return err
}

func (c *cipherSuiteTLSPskWithAes128Ccm) encrypt(pkt *recordLayer, raw []byte) ([]byte, error) {
	if c.ccm == nil {
		return nil, errors.New("CipherSuite has not been initalized, unable to encrypt")
	}

	return c.ccm.encrypt(pkt, raw)
}

func (c *cipherSuiteTLSPskWithAes128Ccm) decrypt(raw []byte) ([]byte, error) {
	if c.ccm == nil {
		return nil, errors.New("CipherSuite has not been initalized, unable to decrypt ")
	}

	return c.ccm.decrypt(raw)
}
