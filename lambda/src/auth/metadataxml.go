package auth

import (
	"encoding/xml"
	"errors"
)

type EntityDescriptor struct {
	IDPSSODescriptor IDPSSODescriptor `xml:"IDPSSODescriptor"`
}

type IDPSSODescriptor struct {
	SingleSignOnServiceList []SingleSignOnService `xml:"SingleSignOnService"`
}

type SingleSignOnService struct {
	Binding  string `xml:"Binding,attr"`
	Location string `xml:"Location,attr"`
}

func GetSSOUrl(xmlpayload []byte) (string, error) {
	var q EntityDescriptor
	err := xml.Unmarshal(xmlpayload, &q)
	if err != nil {
		return "", err
	}
	for _, sso := range q.IDPSSODescriptor.SingleSignOnServiceList {
		if sso.Binding == "urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" {
			return sso.Location, nil
		}
	}
	return "", errors.New("Can't find SSO url in the XML")
}
