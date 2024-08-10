// Package vnc_url provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package vnc_url

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// ErrorResponseModel defines model for ErrorResponseModel.
type ErrorResponseModel struct {
	ErrorReason *string `json:"error_reason,omitempty"`
	Message     *string `json:"message,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}

// RequestConsole defines model for RequestConsole.
type RequestConsole struct {
	Message *string `json:"message,omitempty"`
	Status  *bool   `json:"status,omitempty"`
	Url     *string `json:"url,omitempty"`
}

// VNCURL defines model for VNCURL.
type VNCURL struct {
	Message *string       `json:"message,omitempty"`
	Status  *bool         `json:"status,omitempty"`
	VncUrl  *VNCURLFields `json:"vnc_url,omitempty"`
}

// VNCURLFields defines model for VNCURLFields.
type VNCURLFields struct {
	Id     *int    `json:"id,omitempty"`
	VncUrl *string `json:"vnc_url,omitempty"`
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
	// RequestConsole request
	RequestConsole(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetVNCConsoleLink request
	GetVNCConsoleLink(ctx context.Context, virtualMachineId int, jobId int, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) RequestConsole(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRequestConsoleRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetVNCConsoleLink(ctx context.Context, virtualMachineId int, jobId int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetVNCConsoleLinkRequest(c.Server, virtualMachineId, jobId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewRequestConsoleRequest generates requests for RequestConsole
func NewRequestConsoleRequest(server string, id int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/core/virtual-machines/%s/request-console", pathParam0)
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

// NewGetVNCConsoleLinkRequest generates requests for GetVNCConsoleLink
func NewGetVNCConsoleLinkRequest(server string, virtualMachineId int, jobId int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "virtual_machine_id", runtime.ParamLocationPath, virtualMachineId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "job_id", runtime.ParamLocationPath, jobId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/core/virtual-machines/%s/console/%s", pathParam0, pathParam1)
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
	// RequestConsoleWithResponse request
	RequestConsoleWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*RequestConsoleResponse, error)

	// GetVNCConsoleLinkWithResponse request
	GetVNCConsoleLinkWithResponse(ctx context.Context, virtualMachineId int, jobId int, reqEditors ...RequestEditorFn) (*GetVNCConsoleLinkResponse, error)
}

type RequestConsoleResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RequestConsole
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r RequestConsoleResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RequestConsoleResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetVNCConsoleLinkResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *VNCURL
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r GetVNCConsoleLinkResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetVNCConsoleLinkResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// RequestConsoleWithResponse request returning *RequestConsoleResponse
func (c *ClientWithResponses) RequestConsoleWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*RequestConsoleResponse, error) {
	rsp, err := c.RequestConsole(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRequestConsoleResponse(rsp)
}

// GetVNCConsoleLinkWithResponse request returning *GetVNCConsoleLinkResponse
func (c *ClientWithResponses) GetVNCConsoleLinkWithResponse(ctx context.Context, virtualMachineId int, jobId int, reqEditors ...RequestEditorFn) (*GetVNCConsoleLinkResponse, error) {
	rsp, err := c.GetVNCConsoleLink(ctx, virtualMachineId, jobId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetVNCConsoleLinkResponse(rsp)
}

// ParseRequestConsoleResponse parses an HTTP response from a RequestConsoleWithResponse call
func ParseRequestConsoleResponse(rsp *http.Response) (*RequestConsoleResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RequestConsoleResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RequestConsole
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

// ParseGetVNCConsoleLinkResponse parses an HTTP response from a GetVNCConsoleLinkWithResponse call
func ParseGetVNCConsoleLinkResponse(rsp *http.Response) (*GetVNCConsoleLinkResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetVNCConsoleLinkResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest VNCURL
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
