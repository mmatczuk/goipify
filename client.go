package ipify

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	URL string
}

func (c *Client) CheckIP() (string, error) {
	resp, err := http.Get(c.URL)
	if err != nil {
		return "", fmt.Errorf("failed to access end: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("service error")
	}

	data, err := ioutil.ReadAll(io.LimitReader(resp.Body, 32))
	if err != nil {
		return "", fmt.Errorf("io error: %s", err)
	}

	return string(data), nil
}
