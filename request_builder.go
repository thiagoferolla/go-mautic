package go_mautic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) buildRequest(ctx context.Context, method, url string, request any) (*http.Request, error) {
	furl := c.getFullUrl(url)

	if request == nil {
		return http.NewRequestWithContext(ctx, method, furl, nil)
	}

	var rbytes []byte

	rbytes, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(rbytes)

	return http.NewRequestWithContext(ctx, method, furl, buf)
}

func decodeStringResponse(body io.Reader, v *string) error {
	b, err := io.ReadAll(body)

	if err != nil {
		return err
	}

	*v = string(b)

	return nil
}

func decodeResponse(body io.Reader, v any) error {
	if body == nil {
		return nil
	} else if result, ok := v.(*string); ok {
		return decodeStringResponse(body, result)
	}

	return json.NewDecoder(body).Decode(v)
}

func (c Client) getFullUrl(endpoint string) string {
	return fmt.Sprintf("%s%s", c.config.baseUrl, endpoint)
}

func (c Client) sendRequest(req *http.Request, v any) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")

	req.SetBasicAuth(c.config.user, c.config.password)

	res, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	return decodeResponse(res.Body, v)
}
