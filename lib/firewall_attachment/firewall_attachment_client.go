// Package firewall_attachment provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package firewall_attachment

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// AttachFirewallWithVM defines model for AttachFirewallWithVM.
type AttachFirewallWithVM struct {
	Vms []int `json:"vms"`
}

// ErrorResponseModel defines model for ErrorResponseModel.
type ErrorResponseModel struct {
	ErrorReason *string `json:"error_reason,omitempty"`
	Message     *string `json:"message,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}

// ResponseModel defines model for ResponseModel.
type ResponseModel struct {
	Message *string `json:"message,omitempty"`
	Status  *bool   `json:"status,omitempty"`
}

// AttachFirewallsToVmsJSONRequestBody defines body for AttachFirewallsToVms for application/json ContentType.
type AttachFirewallsToVmsJSONRequestBody = AttachFirewallWithVM

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
	// AttachFirewallsToVmsWithBody request with any body
	AttachFirewallsToVmsWithBody(ctx context.Context, firewallId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AttachFirewallsToVms(ctx context.Context, firewallId int, body AttachFirewallsToVmsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) AttachFirewallsToVmsWithBody(ctx context.Context, firewallId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAttachFirewallsToVmsRequestWithBody(c.Server, firewallId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AttachFirewallsToVms(ctx context.Context, firewallId int, body AttachFirewallsToVmsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAttachFirewallsToVmsRequest(c.Server, firewallId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewAttachFirewallsToVmsRequest calls the generic AttachFirewallsToVms builder with application/json body
func NewAttachFirewallsToVmsRequest(server string, firewallId int, body AttachFirewallsToVmsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAttachFirewallsToVmsRequestWithBody(server, firewallId, "application/json", bodyReader)
}

// NewAttachFirewallsToVmsRequestWithBody generates requests for AttachFirewallsToVms with any type of body
func NewAttachFirewallsToVmsRequestWithBody(server string, firewallId int, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "firewall_id", runtime.ParamLocationPath, firewallId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/core/firewalls/%s/update-attachments", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

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
	// AttachFirewallsToVmsWithBodyWithResponse request with any body
	AttachFirewallsToVmsWithBodyWithResponse(ctx context.Context, firewallId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AttachFirewallsToVmsResponse, error)

	AttachFirewallsToVmsWithResponse(ctx context.Context, firewallId int, body AttachFirewallsToVmsJSONRequestBody, reqEditors ...RequestEditorFn) (*AttachFirewallsToVmsResponse, error)
}

type AttachFirewallsToVmsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ResponseModel
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON403      *ErrorResponseModel
	JSON404      *ErrorResponseModel
	JSON409      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r AttachFirewallsToVmsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AttachFirewallsToVmsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// AttachFirewallsToVmsWithBodyWithResponse request with arbitrary body returning *AttachFirewallsToVmsResponse
func (c *ClientWithResponses) AttachFirewallsToVmsWithBodyWithResponse(ctx context.Context, firewallId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AttachFirewallsToVmsResponse, error) {
	rsp, err := c.AttachFirewallsToVmsWithBody(ctx, firewallId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAttachFirewallsToVmsResponse(rsp)
}

func (c *ClientWithResponses) AttachFirewallsToVmsWithResponse(ctx context.Context, firewallId int, body AttachFirewallsToVmsJSONRequestBody, reqEditors ...RequestEditorFn) (*AttachFirewallsToVmsResponse, error) {
	rsp, err := c.AttachFirewallsToVms(ctx, firewallId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAttachFirewallsToVmsResponse(rsp)
}

// ParseAttachFirewallsToVmsResponse parses an HTTP response from a AttachFirewallsToVmsWithResponse call
func ParseAttachFirewallsToVmsResponse(rsp *http.Response) (*AttachFirewallsToVmsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AttachFirewallsToVmsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ResponseModel
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

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	}

	return response, nil
}
