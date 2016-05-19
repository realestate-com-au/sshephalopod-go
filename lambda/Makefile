#!/usr/bin/env make -f

GOPATH=$(CURDIR):$(CURDIR)/vendor

ZIPFILE=sshephalopod
BINARIES=bin/cakeypair bin/spmetadata bin/auth bin/signing
ZIP_ARGS="-9rq"
REV=$(shell git show-ref --head -s ^HEAD)

CONFIG_BUCKET?=sshephalopod-code-tokyo
DOMAIN?=fridayclub.realestate.com.au
KEYPAIR_BUCKET?=sshephalopod-keys-tokyo
IDP_METADATA?=https://rea.okta.com/app/exk13t140gl0MQ5sX1d8/sso/saml/metadata

CONFIG=config.json
KEYPAIR_NAME=$(DOMAIN)-sshephalopod-ca

BUILD=./bin/gb build
TEST=./bin/gb test

SRCS=$(wildcard src/**/*.go) $(wildcard vendor/src/**/*.go)

.PHONY: devel deploy vendor build $(BINARIES)

all:
	@echo "make build: create the local binaries"
	@echo "make upload: publish zipfile to S3 in DEV"
	@echo "make deploy: update the stacks in DEV"

zip: $(SOURCES) node_modules
	rm -f $(ZIPFILE).zip
	zip $(ZIP_ARGS) $(ZIPFILE).zip $(SOURCES) node_modules

upload: zip
	aws s3 cp $(ZIPFILE).zip s3://$(CONFIG_BUCKET)/$(ZIPFILE)-$(REV).zip
	aws s3 cp $(CONFIG) s3://$(KEYPAIR_BUCKET)/$(CONFIG)

deploy: upload onlydeploy

onlydeploy:
	./deploy.sh $(CONFIG_BUCKET) $(ZIPFILE)-$(REV).zip \
		$(IDP_METADATA) $(DOMAIN) $(KEYPAIR_BUCKET) $(KEYPAIR_NAME)

gofmt: 
	find src/ -type f -iname *.go | xargs gofmt -w -l

bin/gb:
	go build -o bin/gb github.com/constabulary/gb/cmd/gb

bin/gb-vendor:
	go build -o bin/gb-vendor github.com/constabulary/gb/cmd/gb-vendor

vendor:
	./vendor_libs

devel: src/ vendor gofmt ./bin/cakeypair ./bin/spmetadata ./bin/auth ./bin/signing

build: $(BINARIES) test

./bin/cakeypair: $(CAKEYPAIR_SRCS)
	$(BUILD) cmd/cakeypair

./bin/spmetadata: $(SRCS)
	$(BUILD) cmd/spmetadata

./bin/auth: $(SRCS)
	$(BUILD) cmd/auth

./bin/signing: $(SRCS)
	$(BUILD) cmd/signing

test:
	$(TEST)

doc:
	godoc -http :6060
