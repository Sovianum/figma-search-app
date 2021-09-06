package apitest

import (
	"io"
	"net/http"
)

func (s *Suite) NewClient() *Client {
	httpClient := s.server.Client()
	return &Client{
		baseUrl:    s.server.URL,
		httpClient: httpClient,
	}
}

type Client struct {
	baseUrl    string
	httpClient *http.Client
}

type Response struct {
	Status int
	Body   []byte
}

func (cl *Client) Post(url string, body io.Reader) *Response {
	fullPath := cl.baseUrl + url

	resp, err := cl.httpClient.Post(fullPath, "application/json", body)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return &Response{
		Status: resp.StatusCode,
		Body:   b,
	}
}
