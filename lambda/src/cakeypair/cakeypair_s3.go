package cakeypair

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Kv struct {
	Kv     *s3.S3
	Logger *log.Logger
	Bucket string
}

func (s S3Kv) log(err error) {
	if err != nil && s.Logger != nil {
		s.Logger.Print(err)
	}
}

func (s S3Kv) Get(path string) string {
	getparams := s3.GetObjectInput{Bucket: &s.Bucket, Key: &path}
	getresp, err := s.Kv.GetObject(&getparams)
	s.log(err)
	foo, _ := ioutil.ReadAll(getresp.Body)
	return string(foo)
}

func (s S3Kv) Put(path string, body string) error {
	putparams := s3.PutObjectInput{Bucket: &s.Bucket, Key: &path, Body: bytes.NewReader([]byte(body))}
	_, err := s.Kv.PutObject(&putparams)
	s.log(err)
	return err
}
