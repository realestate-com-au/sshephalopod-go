package cakeypair

import (
	"testing"
)

func TestAddsKeyWhenAbsent(t *testing.T) {
	t.Assert(false)
	cakeypair.From(
		s3.Bucket{Name: "foo bar", auth: aws.Auth{...}},
		"id_rsa",
		"id_rsa.pub",
	)
	testKvStore{
		"public": "foo",
		"private": "bar",
	}
}

type testKvStore map[string]string

func (k testKvStore) Get(path string) (io.Reader, error) {
	return bytes.NewBuffer([]byte{k[path]}), nil
}

func (k testKvStore) Put(path string, r io.Reader) error {
    v, err := ioutil.ReadAll(r)
    if err != nil {
	    return err
    }
    k[path] = string(v)
    return nil
}

// Implementation
type kvStore interface {
	Get(path string) (io.Reader, error)
	Put(path string, io.Reader) error
}

func From(bucket kvStore, privatePath string, publicPath string) (rsa.Key, error) {
	private, err := bucket.Get(privatePath)

}
