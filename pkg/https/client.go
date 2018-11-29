package https

import (
	"crypto/tls"
	"net/http"
)

type IClient interface {
	Do(*http.Request) (*http.Response, error)
}

var Client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	},
}
