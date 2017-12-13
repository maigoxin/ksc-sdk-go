package ksc

import (
	"fmt"
	"net/http"
)

type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
}

func Sign(request *http.Request, credentials ...Credentials) *http.Request {
	return Sign4(request, credentials...)
}

// Sign4 signs a request with Signed Signature Version 4.
func Sign4(request *http.Request, credentials ...Credentials) *http.Request {
	keys := chooseKeys(credentials)

	prepareRequestV4(request)
	meta := new(metadata)

	// Task 1
	hashedCanonReq := hashedCanonicalRequestV4(request, meta)

	// Task 2
	stringToSign := stringToSignV4(request, hashedCanonReq, meta)

	// Task 3
	signingKey := signingKeyV4(keys.SecretAccessKey, meta.date, meta.region, meta.service)
	signature := signatureV4(signingKey, stringToSign)

	request.Header.Set("Authorization", buildAuthHeaderV4(signature, meta, keys))

	return request
}

type metadata struct {
	algorithm       string
	credentialScope string
	signedHeaders   string
	date            string
	region          string
	service         string
}

const (
	envAccessKey       = "AWS_ACCESS_KEY"
	envAccessKeyID     = "AWS_ACCESS_KEY_ID"
	envSecretKey       = "AWS_SECRET_KEY"
	envSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
	envSecurityToken   = "AWS_SECURITY_TOKEN"
)
