package cakeypair

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"golang.org/x/crypto/ssh"
)

// Generate an SSH keypair for the CA, if the keypair doesn't already exist
// (as determined by looking in the named S3 bucket for it)
