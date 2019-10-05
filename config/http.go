package config

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

func (cfg Config) HTTPSignatureAuth() runtime.ClientAuthInfoWriter {
	f := func(req runtime.ClientRequest, reg strfmt.Registry) error {

		req.SetHeaderParam("v-c-merchant-id", cfg.MerchantId)
		req.SetHeaderParam("Date", time.Now().Format(time.RFC1123))
		req.SetHeaderParam("Host")

		skipDigest := strings.ToUpper(req.GetMethod()) == http.MethodGet

		if !skipDigest {
			digest, err := generateDigest(req.GetBody())
			if err != nil {
				return err
			}
			req.SetHeaderParam("Digest", digest)
		}

		signature, err := cfg.generateHTTPSignatureHeader(req, skipDigest)
		if err != nil {
			return err
		}
		req.SetHeaderParam("Signature", signature)

		return nil
	}
	return AuthenticateRequestFunc(f)
}

func generateDigest(body []byte) (string, error) {
	hasher := sha256.New()
	if _, err := hasher.Write(body); err != nil {
		return "", fmt.Errorf("unable to hash data: %v", err)
	}
	b64 := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	digest := strings.Replace(b64, "\n", "", -1)
	return "SHA-256=" + digest, nil
}

func (cfg Config) generateHTTPSignatureHeader(req runtime.ClientRequest, skipDigest bool) (string, error) {

	headerNames := []string{"Host", "Date", "(request-target)", "Digest", "v-c-merchant-id"}
	if skipDigest {
		headerNames = append(headerNames[:3], headerNames[4:]...)
	}

	signatureValue, err := cfg.generateHTTPSignatureValue(req, headerNames...)
	if err != nil {
		return "", fmt.Errorf("unable to generate HTTP signature header: %v", err)
	}

	keyid := fmt.Sprintf(`keyid="%s"`, cfg.MerchantKeyId)
	algorithm := fmt.Sprintf(`algorithm="%s"`, "HmacSHA256")
	headers := fmt.Sprintf(`headers="%s"`, strings.ToLower(strings.Join(headerNames, " ")))
	signature := fmt.Sprintf(`signature="%s"`, signatureValue)

	return strings.Join([]string{keyid, algorithm, headers, signature}, ", "), nil
}

func (cfg Config) generateHTTPSignatureValue(req runtime.ClientRequest, headerNames ...string) (string, error) {

	headers := req.GetHeaderParams()

	var keyValuePairs [][]string
	for _, header := range headerNames {

		key := strings.ToLower(header)
		switch key {

		case "host", "date", "digest", "v-c-merchant-id":
			keyValuePairs = append(keyValuePairs, []string{key, headers.Get(header)})

		case "(request-target)":
			u := url.URL{Path: req.GetPath(), RawQuery: req.GetQueryParams().Encode()}
			requestTarget := fmt.Sprintf("%s %s", strings.ToLower(req.GetMethod()), u.String())
			keyValuePairs = append(keyValuePairs, []string{key, requestTarget})
		}
	}

	var lines []string
	for _, kvp := range keyValuePairs {
		lines = append(lines, strings.Join(kvp, ": "))
	}

	toSign := strings.Join(lines, "\n")
	hash := hmac.New(sha256.New, []byte(cfg.MerchantSecretKey))

	if _, err := hash.Write([]byte(toSign)); err != nil {
		return "", fmt.Errorf("failed to sign data: %v", err)
	}

	b64 := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return strings.Replace(b64, "\n", "", -1), nil
}
