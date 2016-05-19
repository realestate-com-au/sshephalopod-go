package main

import (
	"cakeypair"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	bucket := "sshephalopod-keys-tokyo"
	keybase := "fridayclub.realestate.com.au-sshephalopod-ca"

	svc := s3.New(session.New())

	thinger := cakeypair.S3Kv{Bucket: bucket, Kv: svc}

	pub, priv, err := cakeypair.Cakeypair(thinger, keybase)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Pubkey: " + pub + ", PrivKey: " + priv)
}
