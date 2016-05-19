package cakeypair

import (
	"io"
	"testing"

	"crypto/rsa"
)

type testKvStore map[string]string

func TestAddsKeyWhenAbsent(t *testing.T) {
}

func (k testKvStore) GetObject(obj getObjectInput) (foo objectOutput, err error) {
	foo.Body = bytes.NewBuffer([]byte{k[path]})
	err = nil
	return foo, err
}

func (k testKvStore) PutObject(path string, r io.Reader) error {
	v, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	k[path] = string(v)
	return nil
}

type getObjectInput struct {
	Bucket string
	Key    string
}

type putObjectInput struct {
	Bucket string
	Key    string
	Body   string
}

type objectOutput struct {
	Body string
}

// Implementation
type kvStore interface {
	GetObject(obj getObjectInput) (objectOutput, error)
	PutObject(obj putObjectInput) (objectOutput, error)
}

func From(bucket kvStore, privatePath string, publicPath string) (rsa.PrivateKey, error) {
	private, err := bucket.GetObject(privatePath)

}
