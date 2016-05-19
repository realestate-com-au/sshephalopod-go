package cakeypair

import (
	"golang.org/x/crypto/ssh"
)

// Generate an SSH keypair for the CA, if the keypair doesn't already exist
// (as determined by looking in the named S3 bucket for it)

type Kv interface {
	Get(string) string
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

func Cakeypair(thinger Kv, keybase string) (string, string, error) {
	pair := CAKeyPair{thinger.Get(keybase + ".pub"), thinger.Get(keybase)}
	err := pair.validate()

	if err != nil {
		return "", "", err
	}

	return pair.PubKey, pair.PrivKey, nil
}
