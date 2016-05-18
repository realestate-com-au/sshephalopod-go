ZIPFILE=sshephalopod
REV=$(shell git show-ref --head -s ^HEAD)

all: deploy

deploy:
	./deploy.sh $(CONFIG_BUCKET) $(ZIPFILE)-$(REV).zip \
		$(IDP_METADATA) $(DOMAIN) $(KEYPAIR_BUCKET) $(KEYPAIR_NAME)
