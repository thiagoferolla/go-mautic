package go_mautic

import (
	"context"
	"fmt"
	"time"
)

type Contact struct {
	Id             int        `json:"id"`
	DateAdded      time.Time  `json:"dateAdded"`
	CreatedBy      int        `json:"createdBy"`
	CreatedByUser  string     `json:"createdByUser"`
	DateModified   *time.Time `json:"dateModified"`
	ModifiedBy     int        `json:"modifiedBy"`
	ModifiedByUser string     `json:"modifiedByUser"`
	Owner          struct {
		Id        int    `json:"id"`
		Username  string `json:"username"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"owner"`
	Points         int        `json:"points"`
	LastActive     *time.Time `json:"lastActive"`
	DateIdentified *time.Time `json:"dateIdentified"`
	Color          string     `json:"color"`
	IPAddresses    map[string]struct {
		IpAddress string `json:"ipAddress"`
		IpDetails struct {
			City         string `json:"city"`
			Region       string `json:"region"`
			Country      string `json:"country"`
			Latitude     string `json:"latitude"`
			Longitude    string `json:"longitude"`
			Isp          string `json:"isp"`
			Organization string `json:"organization"`
			Timezone     string `json:"timezone"`
		} `json:"ipDetails"`
	} `json:"ipAddresses"`
	Fields  map[string]any `json:"fields"`
	UTMTags []struct {
		Id    int `json:"id"`
		Query struct {
			Page string `json:"page"`
			Cid  string `json:"cid"`
		} `json:"query"`
		Referer     string `json:"referer"`
		RemoteHost  string `json:"remoteHost"`
		UserAgent   string `json:"userAgent"`
		UtmCampaign string `json:"utmCampaign"`
		UtmContent  string `json:"utmContent"`
		UtmMedium   string `json:"utmMedium"`
		UtmSource   string `json:"utmSource"`
		UtmTerm     string `json:"utmTerm"`
	} `json:"utmtags"`
	Tags []struct {
		Tag string `json:"tag"`
	} `json:"tags"`
	DoNotContact []struct {
		Id        int    `json:"id"`
		Reason    int    `json:"reason"`
		Comments  string `json:"comments"`
		Channel   string `json:"channel"`
		ChannelId string `json:"channelId"`
	} `json:"doNotContact"`
}

type getContactResponse struct {
	Contact Contact `json:"contact"`
}

// GetContact - API call to get a contact by its contact ID
// Returns a pointer to a Contact struct and an error
func (c Client) GetContact(ctx context.Context, id string) (*Contact, error) {
	req, err := c.buildRequest(ctx, "GET", fmt.Sprintf("/api/contacts/%s", id), nil)

	if err != nil {
		return nil, err
	}

	var contact getContactResponse

	err = c.sendRequest(req, &contact)

	if err != nil {
		return nil, err
	}

	return &contact.Contact, nil
}

type listContactResponse struct {
	Total    int                `json:"total"`
	Contacts map[string]Contact `json:"contacts"`
}

// ListContacts - API call to get a list of contacts
// Returns a slice of Contact structs and an error
func (c Client) ListContacts(ctx context.Context) ([]Contact, error) {
	req, err := c.buildRequest(ctx, "GET", "/api/contacts", nil)

	if err != nil {
		return nil, err
	}

	var apiResponse listContactResponse

	err = c.sendRequest(req, &apiResponse)

	if err != nil {
		return nil, err
	}

	var contacts []Contact

	for _, contact := range apiResponse.Contacts {
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

type CreateContactRequest struct {
	Fields             map[string]any `json:"fields"`
	IPAddress          *string        `json:"ipAddress,omitempty"`
	LastActive         *string        `json:"lastActive,omitempty"`
	Owner              *string        `json:"owner,omitempty"`
	OverwriteWithBlank *bool          `json:"overwriteWithBlank,omitempty"`
}

func createContactRequestToPayload(pl CreateContactRequest) map[string]any {
	payload := map[string]any{}

	for k, v := range pl.Fields {
		payload[k] = v
	}

	if pl.IPAddress != nil {
		payload["ipAddress"] = pl.IPAddress
	}

	if pl.LastActive != nil {
		payload["lastActive"] = pl.LastActive
	}

	if pl.Owner != nil {
		payload["owner"] = pl.Owner
	}

	if pl.OverwriteWithBlank != nil {
		payload["overwriteWithBlank"] = pl.OverwriteWithBlank
	}

	return payload
}

// CreateContact - API call to create a new contact
// Returns a pointer to a Contact struct and an error
// If the contact already exists, the existing contact will be returned
func (c Client) CreateContact(ctx context.Context, pl CreateContactRequest) (*Contact, error) {
	payload := createContactRequestToPayload(pl)

	req, err := c.buildRequest(ctx, "POST", "/api/contacts/new", payload)

	if err != nil {
		return nil, err
	}

	var apiResponse getContactResponse

	err = c.sendRequest(req, &apiResponse)

	if err != nil {
		return nil, err
	}

	return &apiResponse.Contact, nil
}

// CreateBatchContact - API call to create multiple contacts at the same time
// Returns a slice of Contact structs and an error
// If any of the contacts already exist, the existing contacts will be returned
func (c Client) CreateBatchContact(ctx context.Context, pls []CreateContactRequest) ([]Contact, error) {
	var payload []map[string]any

	for _, pl := range pls {
		payload = append(payload, createContactRequestToPayload(pl))
	}

	req, err := c.buildRequest(ctx, "POST", "/api/contacts/new", payload)

	if err != nil {
		return nil, err
	}

	var apiResponse listContactResponse
	var contacts []Contact

	err = c.sendRequest(req, &apiResponse)

	if err != nil {
		return nil, err
	}

	for _, contact := range apiResponse.Contacts {
		contacts = append(contacts, contact)
	}

	return contacts, nil
}
