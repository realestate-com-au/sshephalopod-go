package spmetadata

import (
	"github.com/RobotsAndPencils/go-saml"
)

func GetEntityDescriptor() (string, error) {
	sp := saml.ServiceProviderSettings{
		PublicCertPath:              "./saml_sp.crt",
		PrivateKeyPath:              "./saml_sp.key",
		IDPSSOURL:                   "http://idp/saml2",
		IDPSSODescriptorURL:         "http://idp/issuer",
		IDPPublicCertPath:           "./idp.crt",
		SPSignRequest:               true,
		AssertionConsumerServiceURL: "https://api-id.execute-api.us-west-2.amazonaws.com/stage/resource-path",
	}
	sp.Init()
	return sp.GetEntityDescriptor()
}
