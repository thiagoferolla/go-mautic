package mautic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

var ErrInvalidFieldID = errors.New("invalid field id")

type FieldType string

const (
	CompanyField FieldType = "company"
	ContactField FieldType = "contact"
)

type Field struct {
	IsPublished         bool       `json:"isPublished"`
	DateAdded           time.Time  `json:"dateAdded"`
	CreatedBy           int        `json:"createdBy"`
	CreatedByUser       string     `json:"createdByUser"`
	DateModified        *time.Time `json:"dateModified"`
	ModifiedBy          *int       `json:"modifiedBy"`
	ModifiedByUser      *string    `json:"modifiedByUser"`
	Id                  int        `json:"id"`
	Label               string     `json:"label"`
	Alias               string     `json:"alias"`
	Type                string     `json:"type"`
	Group               *string    `json:"group"`
	Order               int        `json:"order"`
	Object              string     `json:"object"`
	DefaultValue        any        `json:"defaultValue"`
	IsRequired          bool       `json:"isRequired"`
	IsPubliclyUpdatable bool       `json:"isPubliclyUpdatable"`
	IsUniqueIdentifier  int        `json:"isUniqueIdentifier"`
	Properties          *struct {
		List []struct {
			Label string `json:"label"`
			Value string `json:"value"`
		} `json:"list"`
	} `json:"properties"`
}

type getFieldResponse struct {
	field Field `json:"field"`
}

// GetField - API call to get a field by ID
// the field type must be either CompanyField or ContactField
// Returns a pointer to a Field struct and a error
func (c *Client) GetField(ctx context.Context, fieldType FieldType, id int) (*Field, error) {
	req, err := c.buildRequest(ctx, http.MethodGet, fmt.Sprintf("/fields/%s/%d", fieldType, id), nil)

	if err != nil {
		return nil, err
	}

	var field getFieldResponse

	err = c.sendRequest(req, &field)

	if err != nil {
		return nil, err
	}

	return &field.field, nil
}

type ListFieldsParams struct {
	Search        string `url:"search"`
	Start         int    `url:"start"`
	Limit         int    `url:"limit"`
	OrderBy       string `url:"orderBy"`
	OrderByDir    string `url:"orderByDir"`
	PublishedOnly bool   `url:"publishedOnly"`
	Minimal       bool   `url:"minimal"`
}

type listFieldsResponse struct {
	total  int     `json:"total"`
	fields []Field `json:"fields"`
}

// ListFields - API call to get a list of fields
// the field type must be either CompanyField or ContactField
//
// You can pass a list of parameters to filter the results:
//
// Search - String or search command to filter entities by
// Start - The record number to start at (default 0)
// Limit - The maximum number of records to return (default 30)
// OrderBy - Column to sort by. Can use any column listed in the response.
// OrderByDir - Direction to sort by. Can be asc or desc (default asc)
// PublishedOnly - Only return currently published entities.
// Minimal - Return only array of entities without additional lists in it.
//
// Returns a slice of Field structs and a error
func (c *Client) ListFields(ctx context.Context, fieldType FieldType, params ListFieldsParams) ([]Field, error) {
	qs, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	req, err := c.buildRequest(ctx, http.MethodGet, fmt.Sprintf("/fields/%s?%s", fieldType, qs.Encode()), nil)

	if err != nil {
		return nil, err
	}

	var responses listFieldsResponse

	err = c.sendRequest(req, &responses)

	if err != nil {
		return nil, err
	}

	return responses.fields, nil
}

type CreateOrEditFieldParams struct {
	ID                  *int    `json:"-"`
	Label               string  `json:"label"`
	Alias               string  `json:"alias"`
	Description         *string `json:"description"`
	Type                string  `json:"type"`
	Group               string  `json:"group"`
	Order               int     `json:"order"`
	Object              string  `json:"object"`
	DefaultValue        string  `json:"defaultValue"`
	IsRequired          bool    `json:"isRequired"`
	IsPubliclyAvailable bool    `json:"isPubliclyAvailable"`
	IsUniqueIdentifier  bool    `json:"isUniqueIdentifier"`
	Properties          *struct {
		List []struct {
			Label string `json:"label"`
			Value string `json:"value"`
		} `json:"list"`
	} `json:"properties"`
}

// CreateField - API call to create a field
// the field type must be either CompanyField or ContactField
// Returns a pointer to a Field struct and a error
func (c *Client) CreateField(ctx context.Context, fieldType FieldType, params CreateOrEditFieldParams) (*Field, error) {
	req, err := c.buildRequest(ctx, http.MethodPost, fmt.Sprintf("/fields/%s/new", fieldType), params)

	if err != nil {
		return nil, err
	}

	var apiResponse getFieldResponse

	err = c.sendRequest(req, &apiResponse)

	if err != nil {
		return &apiResponse.field, err
	}

	return &apiResponse.field, nil
}

// EditField - API call to edit a field by ID
// the field type must be either CompanyField or ContactField
// If createIfNotExists is true, the field will be created if it does not exist
// Returns a pointer to a Field struct and a error
func (c *Client) EditField(ctx context.Context, fieldType FieldType, params CreateOrEditFieldParams, createIfNotExists bool) (*Field, error) {
	method := http.MethodPut

	if createIfNotExists {
		method = http.MethodPatch
	}

	if params.ID == nil {
		return nil, ErrInvalidFieldID
	}

	req, err := c.buildRequest(ctx, method, fmt.Sprintf("/fields/%s/%d/edit", fieldType, params.ID), params)

	if err != nil {
		return nil, err
	}

	var apiResponse getFieldResponse

	err = c.sendRequest(req, &apiResponse)

	if err != nil {
		return &apiResponse.field, err
	}

	return &apiResponse.field, nil
}

// DeleteField - API call to delete a field by ID
// the field type must be either CompanyField or ContactField
// Returns a error
func (c *Client) DeleteField(ctx context.Context, fieldType FieldType, id int) error {
	req, err := c.buildRequest(ctx, http.MethodDelete, fmt.Sprintf("/fields/%s/%d/delete", fieldType, id), nil)

	if err != nil {
		return err
	}

	var field getFieldResponse

	err = c.sendRequest(req, &field)

	if err != nil {
		return err
	}

	return nil
}
