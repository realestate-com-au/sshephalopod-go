GOPATH=$(CURDIR)/../..:$(CURDIR)/../../vendor

SRC=cakeypair.go

default: build

install: build
	go install

build: cakeypair

cakeypair: $(SRC)
	go build -o $@ $<
