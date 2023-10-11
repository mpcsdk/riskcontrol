package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	c   *resty.Client
	url string
}

func (c *Client) SetUrl(url string) *Client {
	c.url = url
	return c
}
func (c *Client) SetHeader(key, val string) *Client {
	c.c.SetHeader(key, val)
	return c
}

func (c *Client) PostStruct(body interface{}, data any) (*http.Response, error) {
	resp, err := c.c.R().SetBody(body).Post(c.url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return resp.RawResponse, nil
	}
	///
	err = json.Unmarshal(resp.Body(), &data)

	return resp.RawResponse, err
}
func (c *Client) Post(body interface{}) ([]byte, error) {
	resp, err := c.c.R().SetBody(body).Post(c.url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, nil
	}
	///

	return resp.Body(), err
}

func NewClient(url string) *Client {
	c := resty.New()
	c.SetHeader("Content-Type",
		"application/json;charset=utf-8",
	)
	s := &Client{
		c:   c,
		url: url,
	}
	return s
}
