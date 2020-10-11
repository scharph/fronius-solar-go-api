package froniussolar

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// BaseURLV1 ...
	BaseURLV1 = "http://10.0.0.16/solar_api/v1"
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
		fmt.Println("Code:", res.StatusCode)
		body, err := ioutil.ReadAll(res.Body)
		fmt.Println("Data", string(body))

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

// GetCumulationInverterData ..
func (c *Client) GetCumulationInverterData(ctx context.Context, deviceid int) (*CumulationInverterData, error) {

	queryParams := url.Values{}

	queryParams.Add("Scope", "Device")
	queryParams.Add("DataCollection", "CumulationInverterData")
	queryParams.Add("DeviceId", fmt.Sprint(deviceid))
	if deviceid == 0 {
		return nil, errors.New("error: deviceid is 0")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/GetInverterRealtimeData.cgi", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = queryParams.Encode()

	req = req.WithContext(ctx)
	res := struct {
		Body *struct {
			Data *CumulationInverterData
		}
	}{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return res.Body.Data, nil
}

// CommonInverterData ..
func (c *Client) GetCommonInverterData(ctx context.Context, deviceid int) (*CommonInverterData, error) {

	queryParams := url.Values{}

	queryParams.Add("Scope", "Device")
	queryParams.Add("DataCollection", "CommonInverterData")
	queryParams.Add("DeviceId", fmt.Sprint(deviceid))
	if deviceid == 0 {
		return nil, errors.New("error: deviceid is 0")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/GetInverterRealtimeData.cgi", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = queryParams.Encode()

	req = req.WithContext(ctx)
	res := struct {
		Body *struct {
			Data *CommonInverterData
		}
	}{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return res.Body.Data, nil
}

// CommonInverterData ..
func (c *Client) GetMinMaxInverterData(ctx context.Context, deviceid int) (*MinMaxInverterData, error) {

	queryParams := url.Values{}

	queryParams.Add("Scope", "Device")
	queryParams.Add("DataCollection", "MinMaxInverterData")
	queryParams.Add("DeviceId", fmt.Sprint(deviceid))
	if deviceid == 0 {
		return nil, errors.New("error: deviceid is 0")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/GetInverterRealtimeData.cgi", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = queryParams.Encode()

	req = req.WithContext(ctx)
	res := struct {
		Body *struct {
			Data *MinMaxInverterData
		}
	}{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return res.Body.Data, nil
}
