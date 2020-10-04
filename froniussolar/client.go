package froniussolar

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	// BaseURLV1 ...
	BaseURLV1 = "https://scharph.free.beeceptor.com"
)

// NewClient ...
func NewClient() *Client {
	return &Client{
		BaseURL: BaseURLV1,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}
	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

// GetAPIVersionInfo ...
func (c *Client) GetAPIVersionInfo(ctx context.Context) (*Info, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/GetAPIVersion.cgi", strings.Replace(c.BaseURL, "v1", "", 1)), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	res := Info{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetInverterRealtimeData
func (c *Client) GetInverterRealtimeData(ctx context.Context, args InverterRealtimeDataArgs) (*Info, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/GetInverterRealtimeData.cgi", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	// res := Info{}

	// if err := c.sendRequest(req, &res); err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
