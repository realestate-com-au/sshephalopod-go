SUBDIRS=cakeypair
GOPATH=$(CURDIR)/..:$(CURDIR)/../vendor

test: $(SUBDIRS)
	cd $< && go test
