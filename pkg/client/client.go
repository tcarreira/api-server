package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIClient struct {
	client   *http.Client
	endpoint string
}

type Config struct {
	Endpoint string
}

type ClientError struct {
	error
}

var ErrorNotFound = &ClientError{fmt.Errorf("not found")}

func NewAPIClient(conf Config) (*APIClient, error) {
	if conf.Endpoint == "" {
		return nil, fmt.Errorf("endpoint is required")
	}
	cli := APIClient{
		client:   http.DefaultClient,
		endpoint: conf.Endpoint,
	}
	return &cli, nil
}

func newRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "api-client-go:v0")
	return req, nil
}

func (c *APIClient) DoGET(path string) ([]byte, error) {
	req, err := newRequest("GET", c.endpoint+path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrorNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, body)
	}

	return body, err
}

func (c *APIClient) DoPOST(path string, data interface{}) ([]byte, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := newRequest("POST", c.endpoint+path, bytes.NewBuffer(dataBytes))
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrorNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, body)
	}

	return body, err
}

func (c *APIClient) DoPUT(path string, data interface{}) ([]byte, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := newRequest("PUT", c.endpoint+path, bytes.NewBuffer(dataBytes))
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrorNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, body)
	}

	return body, err
}

func (c *APIClient) DoDELETE(path string) error {
	req, err := newRequest("DELETE", c.endpoint+path, nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return ErrorNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, body)
	}

	return nil
}
