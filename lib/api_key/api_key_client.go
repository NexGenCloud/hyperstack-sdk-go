// Package api_key provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package api_key

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"github.com/NexGenCloud/hyperstack-sdk-go/lib/time"

	"github.com/oapi-codegen/runtime"
)

// ApiKeyFields defines model for ApiKeyFields.
type ApiKeyFields struct {
	CreatedAt   *time.CustomTime `json:"created_at,omitempty"`
	Description *string    `json:"description,omitempty"`
	Id          *int       `json:"id,omitempty"`
	Key         *string    `json:"key,omitempty"`
	Name        *string    `json:"name,omitempty"`
}

// CommonResponseModel defines model for CommonResponseModel.
type CommonResponseModel struct {
	Message *string `json:"message,omitempty"`
	Status  *bool   `json:"status,omitempty"`
}

// ErrorResponseModel defines model for ErrorResponseModel.
type ErrorResponseModel struct {
	ErrorReason *string `json:"error_reason,omitempty"`
	Message     *string `json:"message,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}

// GenerateUpdateApiKeyPayload defines model for GenerateUpdateApiKeyPayload.
type GenerateUpdateApiKeyPayload struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
}

// GenerateUpdateApiKeyResponseModel defines model for GenerateUpdateApiKeyResponseModel.
type GenerateUpdateApiKeyResponseModel struct {
	ApiKey  *ApiKeyFields `json:"api_key,omitempty"`
	Message *string       `json:"message,omitempty"`
	Status  *bool         `json:"status,omitempty"`
}

// GetApiKeysResponseModel defines model for GetApiKeysResponseModel.
type GetApiKeysResponseModel struct {
	ApiKeys *[]ApiKeyFields `json:"api_keys,omitempty"`
	Message *string         `json:"message,omitempty"`
	Status  *bool           `json:"status,omitempty"`
}

// GenerateApiKeyJSONRequestBody defines body for GenerateApiKey for application/json ContentType.
type GenerateApiKeyJSONRequestBody = GenerateUpdateApiKeyPayload

// UpdateApiKeyJSONRequestBody defines body for UpdateApiKey for application/json ContentType.
type UpdateApiKeyJSONRequestBody = GenerateUpdateApiKeyPayload

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
	// RetrieveApiKeys request
	RetrieveApiKeys(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GenerateApiKeyWithBody request with any body
	GenerateApiKeyWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	GenerateApiKey(ctx context.Context, body GenerateApiKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiKey request
	DeleteApiKey(ctx context.Context, apiKeyId int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateApiKeyWithBody request with any body
	UpdateApiKeyWithBody(ctx context.Context, apiKeyId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateApiKey(ctx context.Context, apiKeyId int, body UpdateApiKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) RetrieveApiKeys(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRetrieveApiKeysRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GenerateApiKeyWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGenerateApiKeyRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GenerateApiKey(ctx context.Context, body GenerateApiKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGenerateApiKeyRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteApiKey(ctx context.Context, apiKeyId int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiKeyRequest(c.Server, apiKeyId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateApiKeyWithBody(ctx context.Context, apiKeyId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateApiKeyRequestWithBody(c.Server, apiKeyId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateApiKey(ctx context.Context, apiKeyId int, body UpdateApiKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateApiKeyRequest(c.Server, apiKeyId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewRetrieveApiKeysRequest generates requests for RetrieveApiKeys
func NewRetrieveApiKeysRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api-key")
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

// NewGenerateApiKeyRequest calls the generic GenerateApiKey builder with application/json body
func NewGenerateApiKeyRequest(server string, body GenerateApiKeyJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGenerateApiKeyRequestWithBody(server, "application/json", bodyReader)
}

// NewGenerateApiKeyRequestWithBody generates requests for GenerateApiKey with any type of body
func NewGenerateApiKeyRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api-key/generate")
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

// NewDeleteApiKeyRequest generates requests for DeleteApiKey
func NewDeleteApiKeyRequest(server string, apiKeyId int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "api_key_id", runtime.ParamLocationPath, apiKeyId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api-key/%s", pathParam0)
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

// NewUpdateApiKeyRequest calls the generic UpdateApiKey builder with application/json body
func NewUpdateApiKeyRequest(server string, apiKeyId int, body UpdateApiKeyJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateApiKeyRequestWithBody(server, apiKeyId, "application/json", bodyReader)
}

// NewUpdateApiKeyRequestWithBody generates requests for UpdateApiKey with any type of body
func NewUpdateApiKeyRequestWithBody(server string, apiKeyId int, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "api_key_id", runtime.ParamLocationPath, apiKeyId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api-key/%s", pathParam0)
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
	// RetrieveApiKeysWithResponse request
	RetrieveApiKeysWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RetrieveApiKeysResponse, error)

	// GenerateApiKeyWithBodyWithResponse request with any body
	GenerateApiKeyWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GenerateApiKeyResponse, error)

	GenerateApiKeyWithResponse(ctx context.Context, body GenerateApiKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*GenerateApiKeyResponse, error)

	// DeleteApiKeyWithResponse request
	DeleteApiKeyWithResponse(ctx context.Context, apiKeyId int, reqEditors ...RequestEditorFn) (*DeleteApiKeyResponse, error)

	// UpdateApiKeyWithBodyWithResponse request with any body
	UpdateApiKeyWithBodyWithResponse(ctx context.Context, apiKeyId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateApiKeyResponse, error)

	UpdateApiKeyWithResponse(ctx context.Context, apiKeyId int, body UpdateApiKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateApiKeyResponse, error)
}

type RetrieveApiKeysResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetApiKeysResponseModel
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r RetrieveApiKeysResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RetrieveApiKeysResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GenerateApiKeyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GenerateUpdateApiKeyResponseModel
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON403      *ErrorResponseModel
	JSON409      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r GenerateApiKeyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GenerateApiKeyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteApiKeyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CommonResponseModel
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r DeleteApiKeyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteApiKeyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateApiKeyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GenerateUpdateApiKeyResponseModel
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r UpdateApiKeyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateApiKeyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// RetrieveApiKeysWithResponse request returning *RetrieveApiKeysResponse
func (c *ClientWithResponses) RetrieveApiKeysWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RetrieveApiKeysResponse, error) {
	rsp, err := c.RetrieveApiKeys(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRetrieveApiKeysResponse(rsp)
}

// GenerateApiKeyWithBodyWithResponse request with arbitrary body returning *GenerateApiKeyResponse
func (c *ClientWithResponses) GenerateApiKeyWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GenerateApiKeyResponse, error) {
	rsp, err := c.GenerateApiKeyWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGenerateApiKeyResponse(rsp)
}

func (c *ClientWithResponses) GenerateApiKeyWithResponse(ctx context.Context, body GenerateApiKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*GenerateApiKeyResponse, error) {
	rsp, err := c.GenerateApiKey(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGenerateApiKeyResponse(rsp)
}

// DeleteApiKeyWithResponse request returning *DeleteApiKeyResponse
func (c *ClientWithResponses) DeleteApiKeyWithResponse(ctx context.Context, apiKeyId int, reqEditors ...RequestEditorFn) (*DeleteApiKeyResponse, error) {
	rsp, err := c.DeleteApiKey(ctx, apiKeyId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiKeyResponse(rsp)
}

// UpdateApiKeyWithBodyWithResponse request with arbitrary body returning *UpdateApiKeyResponse
func (c *ClientWithResponses) UpdateApiKeyWithBodyWithResponse(ctx context.Context, apiKeyId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateApiKeyResponse, error) {
	rsp, err := c.UpdateApiKeyWithBody(ctx, apiKeyId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateApiKeyResponse(rsp)
}

func (c *ClientWithResponses) UpdateApiKeyWithResponse(ctx context.Context, apiKeyId int, body UpdateApiKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateApiKeyResponse, error) {
	rsp, err := c.UpdateApiKey(ctx, apiKeyId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateApiKeyResponse(rsp)
}

// ParseRetrieveApiKeysResponse parses an HTTP response from a RetrieveApiKeysWithResponse call
func ParseRetrieveApiKeysResponse(rsp *http.Response) (*RetrieveApiKeysResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RetrieveApiKeysResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetApiKeysResponseModel
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

// ParseGenerateApiKeyResponse parses an HTTP response from a GenerateApiKeyWithResponse call
func ParseGenerateApiKeyResponse(rsp *http.Response) (*GenerateApiKeyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GenerateApiKeyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GenerateUpdateApiKeyResponseModel
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

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	}

	return response, nil
}

// ParseDeleteApiKeyResponse parses an HTTP response from a DeleteApiKeyWithResponse call
func ParseDeleteApiKeyResponse(rsp *http.Response) (*DeleteApiKeyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiKeyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CommonResponseModel
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

// ParseUpdateApiKeyResponse parses an HTTP response from a UpdateApiKeyWithResponse call
func ParseUpdateApiKeyResponse(rsp *http.Response) (*UpdateApiKeyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateApiKeyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GenerateUpdateApiKeyResponseModel
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
