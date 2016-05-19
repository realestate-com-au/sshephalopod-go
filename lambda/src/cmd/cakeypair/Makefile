GOPATH=$(CURDIR)/../../..:$(CURDIR)/../../../vendor

SRCS=main.go
DEST=cakeypair

default: build

install: build
	go install

build: $(DEST)

$(DEST): $(SRCS)
	go build -o $@ $^
