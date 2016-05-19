package cakeypair

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	// 	"golang.org/x/crypto/ssh"
)

// Generate an SSH keypair for the CA, if the keypair doesn't already exist
// (as determined by looking in the named S3 bucket for it)

type CAKeyPair struct {
	PubKey  string
	PrivKey string
}

type KeypairGetterInput struct {
	Service *s3.S3
	Bucket  string
	Key     string
}

// if an object store contains a named key, return it
func KeypairGetter(ref *KeypairGetterInput) (*CAKeyPair, error) {

	getparams := s3.GetObjectInput{Bucket: &ref.Bucket, Key: &ref.Key}
	getresp, err := ref.Service.GetObject(&getparams)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(getresp)

	keyPair := CAKeyPair{PubKey: "something", PrivKey: "else"}
	return &keyPair, nil
}

func Cakeypair() {
	svc := s3.New(session.New())

	params := KeypairGetterInput{Service: svc, Bucket: "sshephalopod-keys-tokyo", Key: "fridayclub.realestate.com.au-sshephalopod-ca"}
	resp, err := KeypairGetter(&params)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
