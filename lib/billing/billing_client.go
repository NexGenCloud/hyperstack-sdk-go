// Package billing provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package billing

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"github.com/nexgen/hyperstack-sdk-go/lib/time"
)

// BillingMetricesFields defines model for Billing metrices fields.
type BillingMetricesFields struct {
	Active         *bool      `json:"active,omitempty"`
	BillPerMinute  *float32   `json:"bill_per_minute,omitempty"`
	CreateTime     *time.CustomTime `json:"create_time,omitempty"`
	Name           *string    `json:"name,omitempty"`
	OrganizationId *int       `json:"organization_id,omitempty"`
	ResourceId     *int       `json:"resource_id,omitempty"`
	ResourceType   *string    `json:"resource_type,omitempty"`
	TerminateTime  *time.CustomTime `json:"terminate_time,omitempty"`
	TotalBill      *float32   `json:"total_bill,omitempty"`
	TotalUpTime    *float32   `json:"total_up_time,omitempty"`
}

// BillingMetricesResponse defines model for Billing metrices response.
type BillingMetricesResponse struct {
	Data    *[]BillingMetricesFields `json:"data,omitempty"`
	Message *string                  `json:"message,omitempty"`
	Status  *bool                    `json:"status,omitempty"`
}

// ErrorResponseModel defines model for ErrorResponseModel.
type ErrorResponseModel struct {
	ErrorReason *string `json:"error_reason,omitempty"`
	Message     *string `json:"message,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}

// LastDayCostFields defines model for Last day cost fields.
type LastDayCostFields struct {
	ClustersCost  *float32 `json:"clusters_cost,omitempty"`
	InstancesCost *float32 `json:"instances_cost,omitempty"`
	TotalCost     *float32 `json:"total_cost,omitempty"`
	VolumesCost   *float32 `json:"volumes_cost,omitempty"`
}

// LastDayCostResponse defines model for Last day cost response.
type LastDayCostResponse struct {
	Data    *LastDayCostFields `json:"data,omitempty"`
	Message *string            `json:"message,omitempty"`
	Status  *bool              `json:"status,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetLastDayCost request
	GetLastDayCost(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetUsage request
	GetUsage(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetLastDayCost(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLastDayCostRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetUsage(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUsageRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetLastDayCostRequest generates requests for GetLastDayCost
func NewGetLastDayCostRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/billing/billing/last-day-cost")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetUsageRequest generates requests for GetUsage
func NewGetUsageRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/billing/billing/usage")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetLastDayCostWithResponse request
	GetLastDayCostWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetLastDayCostResponse, error)

	// GetUsageWithResponse request
	GetUsageWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUsageResponse, error)
}

type GetLastDayCostResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *LastDayCostResponse
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON403      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r GetLastDayCostResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLastDayCostResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUsageResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *BillingMetricesResponse
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON403      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r GetUsageResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUsageResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetLastDayCostWithResponse request returning *GetLastDayCostResponse
func (c *ClientWithResponses) GetLastDayCostWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetLastDayCostResponse, error) {
	rsp, err := c.GetLastDayCost(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLastDayCostResponse(rsp)
}

// GetUsageWithResponse request returning *GetUsageResponse
func (c *ClientWithResponses) GetUsageWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUsageResponse, error) {
	rsp, err := c.GetUsage(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUsageResponse(rsp)
}

// ParseGetLastDayCostResponse parses an HTTP response from a GetLastDayCostWithResponse call
func ParseGetLastDayCostResponse(rsp *http.Response) (*GetLastDayCostResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLastDayCostResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest LastDayCostResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseGetUsageResponse parses an HTTP response from a GetUsageWithResponse call
func ParseGetUsageResponse(rsp *http.Response) (*GetUsageResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUsageResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest BillingMetricesResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}
