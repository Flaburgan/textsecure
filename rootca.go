package textsecure

import (
	"crypto/x509"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

// rootPEM is the PEM formatted signing certificate of the Open Whisper Systems
// server to be used by the TLS client to verify its authenticity instead of
// relying on the system-wide set of root certificates.
var rootPEM = `
-----BEGIN CERTIFICATE-----
MIID4zCCAsugAwIBAgICEBgwDQYJKoZIhvcNAQELBQAwgY0xCzAJBgNVBAYTAlVT
MRMwEQYDVQQIDApDYWxpZm9ybmlhMRYwFAYDVQQHDA1TYW4gRnJhbmNpc2NvMR0w
GwYDVQQKDBRPcGVuIFdoaXNwZXIgU3lzdGVtczEdMBsGA1UECwwUT3BlbiBXaGlz
cGVyIFN5c3RlbXMxEzARBgNVBAMMClRleHRTZWN1cmUwHhcNMTkwMjE1MTczODE3
WhcNMjkwMzEyMTgyMDIwWjCBkDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlm
b3JuaWExHTAbBgNVBAoMFE9wZW4gV2hpc3BlciBTeXN0ZW1zMR0wGwYDVQQLDBRP
cGVuIFdoaXNwZXIgU3lzdGVtczEuMCwGA1UEAwwldGV4dHNlY3VyZS1zZXJ2aWNl
LndoaXNwZXJzeXN0ZW1zLm9yZzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC
ggEBAKzIEbXRRbfAosvPk4magHWzsHhwOzu7On7EA4xxqViHbN4ox4jl5Lh9mu6n
VW0eBvxc9zQKPG0ijgQJN/SV53jFwjqqtr4JYTsHzKs6bgHlYH6sW3XHxePj5JFK
SSXWY7lKNASVl5KkSmhaiYItEPExvSoPB9bNwupixZ5Ae0iIE/NYQA6yZXpQTY0d
BU0l1q0pQeXzLXqgJetThzSXr6j5soNO2KyRoMBNbI42fPUYvWRCOUfyUNI2fb3q
suZD+QQ7YKxl5hgDBU8oNCNN80sNWjhh5nFEOWGj5lxl1qYTkp3sWJJGYD6cuQDJ
1DrSKNbDUWnslIe+wvZfTx9+km0CAwEAAaNIMEYwRAYDVR0RBD0wO4IldGV4dHNl
Y3VyZS1zZXJ2aWNlLndoaXNwZXJzeXN0ZW1zLm9yZ4ISc2VydmljZS5zaWduYWwu
b3JnMA0GCSqGSIb3DQEBCwUAA4IBAQApay5HvPcMP+HE2vS3WOxL/ygG1o/q4zcO
/VYOfA7q2yiFN2FDF8lEcwEqcDMAz2+hGK/fXi2gaIYq6fp3fL9OtzIrXmUNCB2I
9PpuI4jj6xUtERecOXSaHE2C3TI3t7CIcvhbGU1OrJiDLbVFHE8RAetsJJyd2YWu
zBwd9U3oWS4ZNzjlwQLTOiJpoApSKmMlQ6OVfgdr6rRTI1ocw+q4/wDxcYEhiLoM
ljy42A/WrwXzyUMDkcAtZHTjkUAuSLivn434nLcYXalMUIW8sQNLksKTqVH26MKS
2t2HRVs4cwDfmtGzmWSLbgRBl/8Oquq5XLLNEUIM31NVcBUFpKhJ
-----END CERTIFICATE-----
`

var rootCA *x509.CertPool

func setupCA() {
	pem := []byte(rootPEM)
	if config.RootCA != "" && exists(config.RootCA) {
		b, err := ioutil.ReadFile(config.RootCA)
		if err != nil {
			log.Error(err)
			return
		}
		pem = b
	}

	rootCA = x509.NewCertPool()
	if !rootCA.AppendCertsFromPEM(pem) {
		log.Error("Cannot load PEM")
	}
}
