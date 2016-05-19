package cakeypair

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"golang.org/x/crypto/ssh"
)

const KEY_SIZE = 2048

// Generate an SSH keypair for the CA, if the keypair doesn't already exist
// (as determined by looking in the named S3 bucket for it)

type Kv interface {
	Get(string) string
	Put(string, string) error
}

type CAKeyPair struct {
	PubKey  string
	PrivKey string
}

func (c CAKeyPair) validate() error {
	_, _, _, _, err := ssh.ParseAuthorizedKey([]byte(c.PubKey))
	if err != nil {
		return err
	}

	_, err = ssh.ParsePrivateKey([]byte(c.PrivKey))
	if err != nil {
		return err
	}

	return nil
}

func NewCAKeyPair() (CAKeyPair, error) {
	private, err := rsa.GenerateKey(rand.Reader, KEY_SIZE)
	if err != nil {
		return CAKeyPair{}, err
	}

	privKeyDER := x509.MarshalPKCS1PrivateKey(private)
	privKeyBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privKeyDER,
	}

	pubkey, err := ssh.NewPublicKey(private.Public())
	if err != nil {
		return CAKeyPair{}, err
	}

	pubKeyAuth := ssh.MarshalAuthorizedKey(pubkey)

	c := CAKeyPair{
		PrivKey: string(pem.EncodeToMemory(&privKeyBlock)),
		PubKey:  string(pubKeyAuth),
	}

	return c, nil
}

func Cakeypair(thinger Kv, keybase string) (string, string, error) {
	pair := CAKeyPair{
		thinger.Get(keybase + ".pub"),
		thinger.Get(keybase),
	}
	err := pair.validate()

	if err == nil {
		return pair.PubKey, pair.PrivKey, nil
	}

	pair, err = NewCAKeyPair()

	if err != nil {
		return "", "", err
	}

	return pair.PubKey, pair.PrivKey, nil
}
