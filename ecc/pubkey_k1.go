package ecc

import (
	"fmt"

	"github.com/aurawing/eos-go/btcsuite/btcd/btcec"
)

type innerK1PublicKey struct {
}

func newInnerK1PublicKey() innerPublicKey {
	return &innerK1PublicKey{}
}

func (p *innerK1PublicKey) key(content []byte) (*btcec.PublicKey, error) {
	key, err := btcec.ParsePubKey(content, btcec.S256())
	if err != nil {
		return nil, fmt.Errorf("parsePubKey: %s", err)
	}

	return key, nil
}

func (p *innerK1PublicKey) prefix() string {
	return PublicKeyPrefixCompat
}

func (p *innerK1PublicKey) keyMaterialSize() *int {
	return publicKeyDataSize
}

type innerK1PublicKeyYTA struct {
}

func newInnerK1PublicKeyYTA() innerPublicKey {
	return &innerK1PublicKeyYTA{}
}

func (p *innerK1PublicKeyYTA) key(content []byte) (*btcec.PublicKey, error) {
	key, err := btcec.ParsePubKey(content, btcec.S256())
	if err != nil {
		return nil, fmt.Errorf("parsePubKey: %s", err)
	}

	return key, nil
}

func (p *innerK1PublicKeyYTA) prefix() string {
	return PublicKeyPrefixCompatYTA
}

func (p *innerK1PublicKeyYTA) keyMaterialSize() *int {
	return publicKeyDataSize
}
