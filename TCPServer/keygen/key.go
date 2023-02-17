package keygen

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
)

func NewKey() (publicKey string, privateKey string, err error) {
	pvtKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return publicKey, privateKey, err
	}

	x509Encoded, err := x509.MarshalECPrivateKey(pvtKey)
	if err != nil {
		return publicKey, privateKey, err
	}

	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, err := x509.MarshalPKIXPublicKey(&pvtKey.PublicKey)
	if err != nil {
		return publicKey, privateKey, err
	}

	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	privateKey = string(pemEncoded)
	publicKey = string(pemEncodedPub)

	return publicKey, privateKey, err
}
