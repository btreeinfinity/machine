package egoscale

import (
	"net/http"
	"crypto/tls"
)

func NewClient(endpoint string, apiKey string, apiSecret string) *Client {
	cs := &Client {
		client: &http.Client {
			Transport: &http.Transport {
				Proxy: http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config { InsecureSkipVerify: false },
			},
		},
		endpoint: endpoint,
		apiKey: apiKey,
		apiSecret: apiSecret,
	}
	return cs
}
