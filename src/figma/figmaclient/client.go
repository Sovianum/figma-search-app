package figmaclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
)

var figmaURL = parseURLOrPanic("https://api.figma.com")

type Client interface {
	GetFile(ctx context.Context, id FileID) (*File, error)
}

func NewClient(httpClient *http.Client, token string) Client {
	return &client{
		token:      token,
		httpClient: httpClient,
	}
}

type client struct {
	token      string
	httpClient *http.Client
}

var _ Client = (*client)(nil)

func (c *client) GetFile(ctx context.Context, id FileID) (*File, error) {
	req, err := c.newGetRequest(ctx, fmt.Sprintf("v1/files/%s", id))
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result File
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) newGetRequest(ctx context.Context, urlSuffix string) (*http.Request, error) {
	requestURL := *figmaURL
	requestURL.Path = path.Join(requestURL.Path, urlSuffix)

	req, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("X-Figma-Token", c.token)

	return req, nil
}
