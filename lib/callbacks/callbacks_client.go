// Package callbacks provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package callbacks

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

// AttachCallbackPayload defines model for Attach Callback Payload.
type AttachCallbackPayload struct {
	Url string `json:"url"`
}

// AttachCallbackResponse defines model for Attach Callback Response.
type AttachCallbackResponse struct {
	Message *string `json:"message,omitempty"`
	Status  *bool   `json:"status,omitempty"`
	Url     *string `json:"url,omitempty"`
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

// AttachCallbackToVirtualMachineJSONRequestBody defines body for AttachCallbackToVirtualMachine for application/json ContentType.
type AttachCallbackToVirtualMachineJSONRequestBody = AttachCallbackPayload

// UpdateVirtualMachineCallbackJSONRequestBody defines body for UpdateVirtualMachineCallback for application/json ContentType.
type UpdateVirtualMachineCallbackJSONRequestBody = AttachCallbackPayload

// AttachCallbackToVolumeJSONRequestBody defines body for AttachCallbackToVolume for application/json ContentType.
type AttachCallbackToVolumeJSONRequestBody = AttachCallbackPayload

// UpdateVolumeCallbackJSONRequestBody defines body for UpdateVolumeCallback for application/json ContentType.
type UpdateVolumeCallbackJSONRequestBody = AttachCallbackPayload

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
	// AttachCallbackToVirtualMachineWithBody request with any body
	AttachCallbackToVirtualMachineWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AttachCallbackToVirtualMachine(ctx context.Context, id int, body AttachCallbackToVirtualMachineJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteVirtualMachineCallback request
	DeleteVirtualMachineCallback(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateVirtualMachineCallbackWithBody request with any body
	UpdateVirtualMachineCallbackWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateVirtualMachineCallback(ctx context.Context, id int, body UpdateVirtualMachineCallbackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AttachCallbackToVolumeWithBody request with any body
	AttachCallbackToVolumeWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AttachCallbackToVolume(ctx context.Context, id int, body AttachCallbackToVolumeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteVolumeCallback request
	DeleteVolumeCallback(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateVolumeCallbackWithBody request with any body
	UpdateVolumeCallbackWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateVolumeCallback(ctx context.Context, id int, body UpdateVolumeCallbackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) AttachCallbackToVirtualMachineWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAttachCallbackToVirtualMachineRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AttachCallbackToVirtualMachine(ctx context.Context, id int, body AttachCallbackToVirtualMachineJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAttachCallbackToVirtualMachineRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteVirtualMachineCallback(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteVirtualMachineCallbackRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateVirtualMachineCallbackWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateVirtualMachineCallbackRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateVirtualMachineCallback(ctx context.Context, id int, body UpdateVirtualMachineCallbackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateVirtualMachineCallbackRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AttachCallbackToVolumeWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAttachCallbackToVolumeRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AttachCallbackToVolume(ctx context.Context, id int, body AttachCallbackToVolumeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAttachCallbackToVolumeRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteVolumeCallback(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteVolumeCallbackRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateVolumeCallbackWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateVolumeCallbackRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateVolumeCallback(ctx context.Context, id int, body UpdateVolumeCallbackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateVolumeCallbackRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewAttachCallbackToVirtualMachineRequest calls the generic AttachCallbackToVirtualMachine builder with application/json body
func NewAttachCallbackToVirtualMachineRequest(server string, id int, body AttachCallbackToVirtualMachineJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAttachCallbackToVirtualMachineRequestWithBody(server, id, "application/json", bodyReader)
}

// NewAttachCallbackToVirtualMachineRequestWithBody generates requests for AttachCallbackToVirtualMachine with any type of body
func NewAttachCallbackToVirtualMachineRequestWithBody(server string, id int, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/core/virtual-machines/%s/attach-callback", pathParam0)
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

// NewDeleteVirtualMachineCallbackRequest generates requests for DeleteVirtualMachineCallback
func NewDeleteVirtualMachineCallbackRequest(server string, id int) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/core/virtual-machines/%s/delete-callback", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateVirtualMachineCallbackRequest calls the generic UpdateVirtualMachineCallback builder with application/json body
func NewUpdateVirtualMachineCallbackRequest(server string, id int, body UpdateVirtualMachineCallbackJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateVirtualMachineCallbackRequestWithBody(server, id, "application/json", bodyReader)
}

// NewUpdateVirtualMachineCallbackRequestWithBody generates requests for UpdateVirtualMachineCallback with any type of body
func NewUpdateVirtualMachineCallbackRequestWithBody(server string, id int, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/core/virtual-machines/%s/update-callback", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAttachCallbackToVolumeRequest calls the generic AttachCallbackToVolume builder with application/json body
func NewAttachCallbackToVolumeRequest(server string, id int, body AttachCallbackToVolumeJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAttachCallbackToVolumeRequestWithBody(server, id, "application/json", bodyReader)
}

// NewAttachCallbackToVolumeRequestWithBody generates requests for AttachCallbackToVolume with any type of body
func NewAttachCallbackToVolumeRequestWithBody(server string, id int, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/core/volumes/%s/attach-callback", pathParam0)
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

// NewDeleteVolumeCallbackRequest generates requests for DeleteVolumeCallback
func NewDeleteVolumeCallbackRequest(server string, id int) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/core/volumes/%s/delete-callback", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateVolumeCallbackRequest calls the generic UpdateVolumeCallback builder with application/json body
func NewUpdateVolumeCallbackRequest(server string, id int, body UpdateVolumeCallbackJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateVolumeCallbackRequestWithBody(server, id, "application/json", bodyReader)
}

// NewUpdateVolumeCallbackRequestWithBody generates requests for UpdateVolumeCallback with any type of body
func NewUpdateVolumeCallbackRequestWithBody(server string, id int, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/core/volumes/%s/update-callback", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
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
	// AttachCallbackToVirtualMachineWithBodyWithResponse request with any body
	AttachCallbackToVirtualMachineWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AttachCallbackToVirtualMachineResponse, error)

	AttachCallbackToVirtualMachineWithResponse(ctx context.Context, id int, body AttachCallbackToVirtualMachineJSONRequestBody, reqEditors ...RequestEditorFn) (*AttachCallbackToVirtualMachineResponse, error)

	// DeleteVirtualMachineCallbackWithResponse request
	DeleteVirtualMachineCallbackWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteVirtualMachineCallbackResponse, error)

	// UpdateVirtualMachineCallbackWithBodyWithResponse request with any body
	UpdateVirtualMachineCallbackWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateVirtualMachineCallbackResponse, error)

	UpdateVirtualMachineCallbackWithResponse(ctx context.Context, id int, body UpdateVirtualMachineCallbackJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateVirtualMachineCallbackResponse, error)

	// AttachCallbackToVolumeWithBodyWithResponse request with any body
	AttachCallbackToVolumeWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AttachCallbackToVolumeResponse, error)

	AttachCallbackToVolumeWithResponse(ctx context.Context, id int, body AttachCallbackToVolumeJSONRequestBody, reqEditors ...RequestEditorFn) (*AttachCallbackToVolumeResponse, error)

	// DeleteVolumeCallbackWithResponse request
	DeleteVolumeCallbackWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteVolumeCallbackResponse, error)

	// UpdateVolumeCallbackWithBodyWithResponse request with any body
	UpdateVolumeCallbackWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateVolumeCallbackResponse, error)

	UpdateVolumeCallbackWithResponse(ctx context.Context, id int, body UpdateVolumeCallbackJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateVolumeCallbackResponse, error)
}

type AttachCallbackToVirtualMachineResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AttachCallbackResponse
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r AttachCallbackToVirtualMachineResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AttachCallbackToVirtualMachineResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteVirtualMachineCallbackResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ResponseModel
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r DeleteVirtualMachineCallbackResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteVirtualMachineCallbackResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateVirtualMachineCallbackResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AttachCallbackResponse
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r UpdateVirtualMachineCallbackResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateVirtualMachineCallbackResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AttachCallbackToVolumeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AttachCallbackResponse
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r AttachCallbackToVolumeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AttachCallbackToVolumeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteVolumeCallbackResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ResponseModel
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r DeleteVolumeCallbackResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteVolumeCallbackResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateVolumeCallbackResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AttachCallbackResponse
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r UpdateVolumeCallbackResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateVolumeCallbackResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// AttachCallbackToVirtualMachineWithBodyWithResponse request with arbitrary body returning *AttachCallbackToVirtualMachineResponse
func (c *ClientWithResponses) AttachCallbackToVirtualMachineWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AttachCallbackToVirtualMachineResponse, error) {
	rsp, err := c.AttachCallbackToVirtualMachineWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAttachCallbackToVirtualMachineResponse(rsp)
}

func (c *ClientWithResponses) AttachCallbackToVirtualMachineWithResponse(ctx context.Context, id int, body AttachCallbackToVirtualMachineJSONRequestBody, reqEditors ...RequestEditorFn) (*AttachCallbackToVirtualMachineResponse, error) {
	rsp, err := c.AttachCallbackToVirtualMachine(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAttachCallbackToVirtualMachineResponse(rsp)
}

// DeleteVirtualMachineCallbackWithResponse request returning *DeleteVirtualMachineCallbackResponse
func (c *ClientWithResponses) DeleteVirtualMachineCallbackWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteVirtualMachineCallbackResponse, error) {
	rsp, err := c.DeleteVirtualMachineCallback(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteVirtualMachineCallbackResponse(rsp)
}

// UpdateVirtualMachineCallbackWithBodyWithResponse request with arbitrary body returning *UpdateVirtualMachineCallbackResponse
func (c *ClientWithResponses) UpdateVirtualMachineCallbackWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateVirtualMachineCallbackResponse, error) {
	rsp, err := c.UpdateVirtualMachineCallbackWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateVirtualMachineCallbackResponse(rsp)
}

func (c *ClientWithResponses) UpdateVirtualMachineCallbackWithResponse(ctx context.Context, id int, body UpdateVirtualMachineCallbackJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateVirtualMachineCallbackResponse, error) {
	rsp, err := c.UpdateVirtualMachineCallback(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateVirtualMachineCallbackResponse(rsp)
}

// AttachCallbackToVolumeWithBodyWithResponse request with arbitrary body returning *AttachCallbackToVolumeResponse
func (c *ClientWithResponses) AttachCallbackToVolumeWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AttachCallbackToVolumeResponse, error) {
	rsp, err := c.AttachCallbackToVolumeWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAttachCallbackToVolumeResponse(rsp)
}

func (c *ClientWithResponses) AttachCallbackToVolumeWithResponse(ctx context.Context, id int, body AttachCallbackToVolumeJSONRequestBody, reqEditors ...RequestEditorFn) (*AttachCallbackToVolumeResponse, error) {
	rsp, err := c.AttachCallbackToVolume(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAttachCallbackToVolumeResponse(rsp)
}

// DeleteVolumeCallbackWithResponse request returning *DeleteVolumeCallbackResponse
func (c *ClientWithResponses) DeleteVolumeCallbackWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteVolumeCallbackResponse, error) {
	rsp, err := c.DeleteVolumeCallback(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteVolumeCallbackResponse(rsp)
}

// UpdateVolumeCallbackWithBodyWithResponse request with arbitrary body returning *UpdateVolumeCallbackResponse
func (c *ClientWithResponses) UpdateVolumeCallbackWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateVolumeCallbackResponse, error) {
	rsp, err := c.UpdateVolumeCallbackWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateVolumeCallbackResponse(rsp)
}

func (c *ClientWithResponses) UpdateVolumeCallbackWithResponse(ctx context.Context, id int, body UpdateVolumeCallbackJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateVolumeCallbackResponse, error) {
	rsp, err := c.UpdateVolumeCallback(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateVolumeCallbackResponse(rsp)
}

// ParseAttachCallbackToVirtualMachineResponse parses an HTTP response from a AttachCallbackToVirtualMachineWithResponse call
func ParseAttachCallbackToVirtualMachineResponse(rsp *http.Response) (*AttachCallbackToVirtualMachineResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AttachCallbackToVirtualMachineResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AttachCallbackResponse
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

// ParseDeleteVirtualMachineCallbackResponse parses an HTTP response from a DeleteVirtualMachineCallbackWithResponse call
func ParseDeleteVirtualMachineCallbackResponse(rsp *http.Response) (*DeleteVirtualMachineCallbackResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteVirtualMachineCallbackResponse{
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

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseUpdateVirtualMachineCallbackResponse parses an HTTP response from a UpdateVirtualMachineCallbackWithResponse call
func ParseUpdateVirtualMachineCallbackResponse(rsp *http.Response) (*UpdateVirtualMachineCallbackResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateVirtualMachineCallbackResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AttachCallbackResponse
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

// ParseAttachCallbackToVolumeResponse parses an HTTP response from a AttachCallbackToVolumeWithResponse call
func ParseAttachCallbackToVolumeResponse(rsp *http.Response) (*AttachCallbackToVolumeResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AttachCallbackToVolumeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AttachCallbackResponse
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

// ParseDeleteVolumeCallbackResponse parses an HTTP response from a DeleteVolumeCallbackWithResponse call
func ParseDeleteVolumeCallbackResponse(rsp *http.Response) (*DeleteVolumeCallbackResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteVolumeCallbackResponse{
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

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseUpdateVolumeCallbackResponse parses an HTTP response from a UpdateVolumeCallbackWithResponse call
func ParseUpdateVolumeCallbackResponse(rsp *http.Response) (*UpdateVolumeCallbackResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateVolumeCallbackResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AttachCallbackResponse
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
