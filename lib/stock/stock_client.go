// Package stock provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package stock

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ErrorResponseModel defines model for ErrorResponseModel.
type ErrorResponseModel struct {
	ErrorReason *string `json:"error_reason,omitempty"`
	Message     *string `json:"message,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}

// NewConfigurationsResponse defines model for NewConfigurationsResponse.
type NewConfigurationsResponse struct {
	N10x *int `json:"10x,omitempty"`
	N1x  *int `json:"1x,omitempty"`
	N2x  *int `json:"2x,omitempty"`
	N4x  *int `json:"4x,omitempty"`
	N8x  *int `json:"8x,omitempty"`
}

// NewModelResponse defines model for NewModelResponse.
type NewModelResponse struct {
	Available      *string                    `json:"available,omitempty"`
	Configurations *NewConfigurationsResponse `json:"configurations,omitempty"`
	Model          *string                    `json:"model,omitempty"`
	Planned100Days *string                    `json:"planned_100_days,omitempty"`
	Planned30Days  *string                    `json:"planned_30_days,omitempty"`
	Planned7Days   *string                    `json:"planned_7_days,omitempty"`
}

// NewStockResponse defines model for NewStockResponse.
type NewStockResponse struct {
	Models    *[]NewModelResponse `json:"models,omitempty"`
	Region    *string             `json:"region,omitempty"`
	StockType *string             `json:"stock-type,omitempty"`
}

// NewStockRetriveResponse defines model for NewStockRetriveResponse.
type NewStockRetriveResponse struct {
	Stocks *[]NewStockResponse `json:"stocks,omitempty"`
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
	// RetrieveGPUStocks request
	RetrieveGPUStocks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) RetrieveGPUStocks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRetrieveGPUStocksRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewRetrieveGPUStocksRequest generates requests for RetrieveGPUStocks
func NewRetrieveGPUStocksRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/core/stocks")
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
	// RetrieveGPUStocksWithResponse request
	RetrieveGPUStocksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RetrieveGPUStocksResponse, error)
}

type RetrieveGPUStocksResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *NewStockRetriveResponse
	JSON401      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r RetrieveGPUStocksResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RetrieveGPUStocksResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// RetrieveGPUStocksWithResponse request returning *RetrieveGPUStocksResponse
func (c *ClientWithResponses) RetrieveGPUStocksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RetrieveGPUStocksResponse, error) {
	rsp, err := c.RetrieveGPUStocks(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRetrieveGPUStocksResponse(rsp)
}

// ParseRetrieveGPUStocksResponse parses an HTTP response from a RetrieveGPUStocksWithResponse call
func ParseRetrieveGPUStocksResponse(rsp *http.Response) (*RetrieveGPUStocksResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RetrieveGPUStocksResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest NewStockRetriveResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}
