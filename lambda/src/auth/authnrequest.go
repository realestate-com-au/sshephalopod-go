package auth

import (
	"github.com/RobotsAndPencils/go-saml"
)

func GetAuthorizeRequest(ssoUrl string, assertionUrl string) (string, error) {
  sp := saml.ServiceProviderSettings{
		PublicCertPath:              "./saml_sp.crt",
		PrivateKeyPath:              "./saml_sp.key",
		IDPSSOURL:                   ssoUrl,
		IDPSSODescriptorURL:         ssoUrl,
		IDPPublicCertPath:           "./idp.crt",
		SPSignRequest:               false,
		AssertionConsumerServiceURL: assertionUrl,
	}
	sp.Init()

  authnRequest := sp.GetAuthnRequest()
  authnRequestXml, err := authnRequest.String()
  if err != nil {
    return "", err
  }

	return authnRequestXml, nil
}
