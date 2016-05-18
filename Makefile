CONFIG_BUCKET=sshephalopod-config-bucket
KEYPAIR_BUCKET=sshephalopod-keypair-bucket
DOMAIN=sshephalopod-service-domain.com
IDP_METADATA=https://somewhere.okta.com/app/somejumbleofcharacters/sso/saml/metadata

all:
	@echo "use 'make build' to build sshephalopod components"
	@echo "use 'make deploy' to deploy sshephalopod"

build:
	make -C lambda build

deploy:
	make -C lambda deploy
	make -C cloudformation deploy
