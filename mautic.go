package mautic

import (
	"errors"
	"net/http"
)

type Client struct {
	config ClientConfig

	httpClient *http.Client
}

var ErrNoConfigProvided = errors.New("no config provided")

func New(config *ClientConfig) (*Client, error) {
	if config == nil {
		return nil, ErrNoConfigProvided
	}

	return &Client{*config, &http.Client{}}, nil
}
