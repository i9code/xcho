package xcho

import (
	`crypto/tls`
	`github.com/go-resty/resty/v2`
)

// Client 客户端封装
type Client struct {
	*resty.Client
}

// Request 请求封装
type Request struct {
	*resty.Request
}

func NewHttpClient(config *HttpClientConfig) (client *Client, err error) {
	restyClient := resty.New()
	if "" != config.Proxy.Host {
		restyClient.SetProxy(config.Proxy.Addr())
	}
	if 0 != config.Timeout {
		restyClient.SetTimeout(config.Timeout)
	}
	if config.Payload.Get {
		restyClient.SetAllowGetMethodPayload(true)
	}
	if config.Certificate.Skip {
		// nolint:gosec
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	} else {
		if "" != config.Certificate.Root {
			restyClient.SetRootCertificate(config.Certificate.Root)
		}
		if 0 != len(config.Certificate.Clients) {
			certificates := make([]tls.Certificate, 0, len(config.Certificate.Clients))
			for _, c := range config.Certificate.Clients {
				certificate, err := tls.LoadX509KeyPair(c.Public, c.Private)
				if nil != err {
					continue
				}
				certificates = append(certificates, certificate)
			}
			restyClient.SetCertificates(certificates...)
		}
	}
	if 0 != len(config.Headers) {
		restyClient.SetHeaders(config.Headers)
	}
	if 0 != len(config.Queries) {
		restyClient.SetQueryParams(config.Queries)
	}
	if 0 != len(config.Forms) {
		restyClient.SetFormData(config.Forms)
	}
	if 0 != len(config.Cookies) {
		restyClient.SetCookies(config.Cookies)
	}
	if "" != config.Auth.Type {
		switch config.Auth.Type {
		case AuthTypeBasic:
			restyClient.SetBasicAuth(config.Auth.Username, config.Auth.Password)
		case AuthTypeToken:
			restyClient.SetAuthToken(config.Auth.Token)
			if "" != config.Auth.Scheme {
				restyClient.SetAuthScheme(config.Auth.Scheme)
			}
		}
	}
	client = &Client{Client: restyClient}

	return
}

func newRequest(client *Client) *Request {
	return &Request{Request: client.R()}
}
