package go_mautic

import "net/http"

type Client struct {
	config ClientConfig

	httpClient *http.Client
}

func New(config ClientConfig) Client {
	return Client{config, &http.Client{}}
}
