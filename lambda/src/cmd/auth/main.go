package main

import (
  "auth"
  "flag"
  "fmt"
  "os"
)

func Usage() {
  fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
  flag.PrintDefaults()
  os.Exit(1)
}

func main() {
  // url := "https://rea.okta.com/app/exk15c6oowgUqfFqk1d8/sso/saml/metadata"
  url := flag.String("sp", "", "service provider url i.e. https://rea.okta.com/app/zzk87d6hhtrVshNew8k1/sso/saml/metadata")
  apiId := flag.String("api-id", "", "API GW id")
  resourcePath := flag.String("resource-path", "", "API GW resource path")
  stage := flag.String("stage", "prod", "API GW stage")
  region := flag.String("region", "us-west-2", "API GW region")

  flag.Parse()
  if (*url == "") {
    Usage()
  }

  assertionURl := fmt.Sprintf("https://%s.execute-api.%s.amazonaws.com/%s/%s", *apiId, *region, *stage, *resourcePath)

	metadata, err := auth.GetIdpMetadata(*url)
	if err != nil {
		panic(err)
	}

	ssoUrl, err := auth.GetSSOUrl(metadata)
	if err != nil {
		panic(err)
	}

	authnRequest, err := auth.GetAuthorizeRequest(ssoUrl, assertionURl)
	fmt.Println(authnRequest)
}
