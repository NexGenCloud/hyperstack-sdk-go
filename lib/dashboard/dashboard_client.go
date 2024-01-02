// Package dashboard provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package dashboard

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ContainerOverviewFields defines model for Container Overview Fields.
type ContainerOverviewFields struct {
	CostPerHour *float32 `json:"cost_per_hour,omitempty"`
	Count       *int     `json:"count,omitempty"`
	Gpus        *int     `json:"gpus,omitempty"`
	Ram         *int     `json:"ram,omitempty"`
	Vcpus       *int     `json:"vcpus,omitempty"`
}

// DashboardInfoResponse defines model for Dashboard Info Response.
type DashboardInfoResponse struct {
	Message  *string       `json:"message,omitempty"`
	Overview *OverviewInfo `json:"overview,omitempty"`
	Status   *bool         `json:"status,omitempty"`
}

// ErrorResponseModel defines model for ErrorResponseModel.
type ErrorResponseModel struct {
	ErrorReason *string `json:"error_reason,omitempty"`
	Message     *string `json:"message,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}

// InstanceOverviewFields defines model for Instance Overview Fields.
type InstanceOverviewFields struct {
	CostPerHour *float32 `json:"cost_per_hour,omitempty"`
	Count       *int     `json:"count,omitempty"`
	Gpus        *int     `json:"gpus,omitempty"`
	Ram         *int     `json:"ram,omitempty"`
	Vcpus       *int     `json:"vcpus,omitempty"`
}

// OverviewInfo defines model for Overview Info.
type OverviewInfo struct {
	Container *ContainerOverviewFields `json:"container,omitempty"`
	Instance  *InstanceOverviewFields  `json:"instance,omitempty"`
	Volume    *VolumeOverviewFields    `json:"volume,omitempty"`
}

// VolumeOverviewFields defines model for Volume Overview Fields.
type VolumeOverviewFields struct {
	CostPerHour *float32 `json:"cost_per_hour,omitempty"`
	Count       *int     `json:"count,omitempty"`
	Using       *int     `json:"using,omitempty"`
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
	// GetInstancesContainersAndVolumesOverview request
	GetInstancesContainersAndVolumesOverview(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetInstancesContainersAndVolumesOverview(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetInstancesContainersAndVolumesOverviewRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetInstancesContainersAndVolumesOverviewRequest generates requests for GetInstancesContainersAndVolumesOverview
func NewGetInstancesContainersAndVolumesOverviewRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/core/dashboard")
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
	// GetInstancesContainersAndVolumesOverviewWithResponse request
	GetInstancesContainersAndVolumesOverviewWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetInstancesContainersAndVolumesOverviewResponse, error)
}

type GetInstancesContainersAndVolumesOverviewResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *DashboardInfoResponse
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r GetInstancesContainersAndVolumesOverviewResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetInstancesContainersAndVolumesOverviewResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetInstancesContainersAndVolumesOverviewWithResponse request returning *GetInstancesContainersAndVolumesOverviewResponse
func (c *ClientWithResponses) GetInstancesContainersAndVolumesOverviewWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetInstancesContainersAndVolumesOverviewResponse, error) {
	rsp, err := c.GetInstancesContainersAndVolumesOverview(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetInstancesContainersAndVolumesOverviewResponse(rsp)
}

// ParseGetInstancesContainersAndVolumesOverviewResponse parses an HTTP response from a GetInstancesContainersAndVolumesOverviewWithResponse call
func ParseGetInstancesContainersAndVolumesOverviewResponse(rsp *http.Response) (*GetInstancesContainersAndVolumesOverviewResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetInstancesContainersAndVolumesOverviewResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest DashboardInfoResponse
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

	}

	return response, nil
}
