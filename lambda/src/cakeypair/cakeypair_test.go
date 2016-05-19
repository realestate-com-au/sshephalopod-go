package cakeypair

import (
	"io/ioutil"
	"os"
	"testing"
)

type testKvStore map[string]string

func TestGetsValidKeys(t *testing.T) {
	myKv := testKvStore{"foo": "bar", "foo.pub": "banana"}

	pub, priv, err := Cakeypair(myKv, "foo")

	if err == nil {
		t.Fatal("Expected an error where pubkey/privkey are invalid")
		return
	}

	f, err := os.Open("testdata/testkey")
	if err != nil {
		t.Fatal("Cannot open testdata/testkey")
	}
	foo, _ := ioutil.ReadAll(f)
	f.Close()
	myKv["valid"] = string(foo)

	f, err = os.Open("testdata/testkey.pub")
	if err != nil {
		t.Fatal("Cannot open testdata/testkey")
	}
	foo, _ = ioutil.ReadAll(f)
	f.Close()
	myKv["valid.pub"] = string(foo)

	pub, priv, err = Cakeypair(myKv, "valid")

	if err != nil {
		t.Error(err)
		return
	}

	if pub != myKv["valid.pub"] {
		t.Errorf("Expected a valid public key")
	}

	if priv != myKv["valid"] {
		t.Errorf("Expected a valid private key")
	}

}

func (k testKvStore) Get(key string) string {
	return k[key]
}
