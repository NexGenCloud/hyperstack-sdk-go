// Package flavor provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package flavor

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"github.com/nexgen/hyperstack-sdk-go/lib/time"

	"github.com/oapi-codegen/runtime"
)

// ErrorResponseModel defines model for ErrorResponseModel.
type ErrorResponseModel struct {
	ErrorReason *string `json:"error_reason,omitempty"`
	Message     *string `json:"message,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}

// FlavorFields defines model for Flavor Fields.
type FlavorFields struct {
	Cpu            *int       `json:"cpu,omitempty"`
	CreatedAt      *time.CustomTime `json:"created_at,omitempty"`
	Disk           *int       `json:"disk,omitempty"`
	Gpu            *string    `json:"gpu,omitempty"`
	GpuCount       *int       `json:"gpu_count,omitempty"`
	Id             *int       `json:"id,omitempty"`
	Name           *string    `json:"name,omitempty"`
	Ram            *int       `json:"ram,omitempty"`
	RegionName     *string    `json:"region_name,omitempty"`
	StockAvailable *bool      `json:"stock_available,omitempty"`
}

// FlavorItemGetResponse defines model for Flavor Item Get Response.
type FlavorItemGetResponse struct {
	Flavors    *[]FlavorFields `json:"flavors,omitempty"`
	Gpu        *string         `json:"gpu,omitempty"`
	RegionName *string         `json:"region_name,omitempty"`
}

// FlavorListResponse defines model for Flavor List Response.
type FlavorListResponse struct {
	Data    *[]FlavorItemGetResponse `json:"data,omitempty"`
	Message *string                  `json:"message,omitempty"`
	Status  *bool                    `json:"status,omitempty"`
}

// RetrieveFlavorsParams defines parameters for RetrieveFlavors.
type RetrieveFlavorsParams struct {
	Region *interface{} `form:"region,omitempty" json:"region,omitempty"`
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
	// RetrieveFlavors request
	RetrieveFlavors(ctx context.Context, params *RetrieveFlavorsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) RetrieveFlavors(ctx context.Context, params *RetrieveFlavorsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRetrieveFlavorsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewRetrieveFlavorsRequest generates requests for RetrieveFlavors
func NewRetrieveFlavorsRequest(server string, params *RetrieveFlavorsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/core/flavors")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Region != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "region", runtime.ParamLocationQuery, *params.Region); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
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
	// RetrieveFlavorsWithResponse request
	RetrieveFlavorsWithResponse(ctx context.Context, params *RetrieveFlavorsParams, reqEditors ...RequestEditorFn) (*RetrieveFlavorsResponse, error)
}

type RetrieveFlavorsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FlavorListResponse
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r RetrieveFlavorsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RetrieveFlavorsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// RetrieveFlavorsWithResponse request returning *RetrieveFlavorsResponse
func (c *ClientWithResponses) RetrieveFlavorsWithResponse(ctx context.Context, params *RetrieveFlavorsParams, reqEditors ...RequestEditorFn) (*RetrieveFlavorsResponse, error) {
	rsp, err := c.RetrieveFlavors(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRetrieveFlavorsResponse(rsp)
}

// ParseRetrieveFlavorsResponse parses an HTTP response from a RetrieveFlavorsWithResponse call
func ParseRetrieveFlavorsResponse(rsp *http.Response) (*RetrieveFlavorsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RetrieveFlavorsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FlavorListResponse
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

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}
