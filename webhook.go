package mautic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var ErrInvalidWebhookID = errors.New("invalid webhook id")

type Webhook struct {
	IsPublished      bool      `json:"isPublished"`
	DateAdded        time.Time `json:"dateAdded"`
	DateModified     time.Time `json:"dateModified"`
	CreatedBy        int       `json:"createdBy"`
	CreatedByUser    string    `json:"createdByUser"`
	ModifiedBy       *string   `json:"modifiedBy"`
	ModifiedByUser   string    `json:"modifiedByUser"`
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	WebhookUrl       string    `json:"webhookUrl"`
	Secret           string    `json:"secret"`
	EventsOrderbyDir string    `json:"eventsOrderbyDir"`
	Category         struct {
		CreatedByUser  string  `json:"createdByUser"`
		ModifiedByUser string  `json:"modifiedByUser"`
		Id             int     `json:"id"`
		Title          string  `json:"title"`
		Alias          string  `json:"alias"`
		Description    *string `json:"description"`
		Color          *string `json:"color"`
		Bundle         string  `json:"bundle"`
	} `json:"category"`
	Triggers []string `json:"triggers"`
}

type getWebhookResponse struct {
	hook Webhook `json:"hook"`
}

// GetWebhook - API call to get a webhook by ID
// Returns a pointer to a Webhook struct and a error
func (c *Client) GetWebhook(ctx context.Context, id string) (*Webhook, error) {
	req, err := c.buildRequest(ctx, http.MethodGet, fmt.Sprintf("/hooks/%s", id), nil)

	if err != nil {
		return nil, err
	}

	var response getWebhookResponse

	err = c.sendRequest(req, &response)

	if err != nil {
		return nil, err
	}

	return &response.hook, nil
}

type listWebhooksResponse struct {
	total int                `json:"total"`
	hooks map[string]Webhook `json:"hooks"`
}

// ListWebhooks - API call to get the list of active webhooks
// Returns a slice of Webhook structs and an error
func (c *Client) ListWebhooks(ctx context.Context) ([]Webhook, error) {
	req, err := c.buildRequest(ctx, http.MethodGet, "/hooks", nil)

	if err != nil {
		return nil, err
	}

	var response listWebhooksResponse

	err = c.sendRequest(req, &response)

	if err != nil {
		return nil, err
	}

	var hooks []Webhook

	for _, hook := range response.hooks {
		hooks = append(hooks, hook)
	}

	return hooks, nil
}

type WebhookParams struct {
	ID               *int     `json:"id,omitempty"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	WebhookUrl       string   `json:"webhookUrl"`
	Secret           *string  `json:"secret"`
	EventsOrderbyDir string   `json:"eventsOrderbyDir"`
	Triggers         []string `json:"triggers"`
}

// CreateWebhook - API call to create a new webhook
// Returns a pointer to a Webhook struct and a error
func (c *Client) CreateWebhook(ctx context.Context, params WebhookParams) (*Webhook, error) {
	req, err := c.buildRequest(ctx, http.MethodPost, "/hooks/new", params)

	if err != nil {
		return nil, err
	}

	var response getWebhookResponse

	err = c.sendRequest(req, &response)

	if err != nil {
		return nil, err
	}

	return &response.hook, nil
}

// EditWebhook - API call to edit a webhook by ID
// If createIfNotExists is true, the webhook will be created if it does not exist
// Returns a pointer to a Webhook struct and a error
func (c *Client) EditWebhook(ctx context.Context, params WebhookParams, createIfNotExists bool) (*Webhook, error) {
	method := http.MethodPatch

	if createIfNotExists {
		method = http.MethodPut
	}

	if params.ID == nil {
		return nil, ErrInvalidWebhookID
	}

	req, err := c.buildRequest(ctx, method, fmt.Sprintf("/hooks/%s/update", params.ID), params)

	if err != nil {
		return nil, err
	}

	var response getWebhookResponse

	err = c.sendRequest(req, &response)

	if err != nil {
		return nil, err
	}

	return &response.hook, nil

}

// DeleteWebhook - API call to delete a webhook by ID
// Returns nil if successful, otherwise an error
func (c *Client) DeleteWebhook(ctx context.Context, id string) error {
	req, err := c.buildRequest(ctx, http.MethodDelete, fmt.Sprintf("/hooks/%s/delete", id), nil)

	if err != nil {
		return err
	}

	err = c.sendRequest(req, nil)

	return err
}

type WebhookTrigger struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type listWebhookTriggersResponse struct {
	triggers map[string]struct {
		Label       string `json:"label"`
		Description string `json:"description"`
	} `json:"triggers"`
}

// ListWebhookTriggers - API call to get the list of available webhook triggers
// Returns a slice of WebhookTrigger structs and an error
func (c *Client) ListWebhookTriggers(ctx context.Context) ([]WebhookTrigger, error) {
	req, err := c.buildRequest(ctx, http.MethodGet, "/hooks/triggers", nil)

	if err != nil {
		return nil, err
	}

	var apiResponse listWebhookTriggersResponse

	err = c.sendRequest(req, &apiResponse)

	if err != nil {
		return nil, err
	}

	var triggers []WebhookTrigger

	for k, v := range apiResponse.triggers {
		triggers = append(triggers, WebhookTrigger{
			Name:        k,
			Label:       v.Label,
			Description: v.Description,
		})
	}

	return triggers, nil
}
