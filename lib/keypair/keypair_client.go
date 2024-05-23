// Package keypair provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package keypair

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

// ErrorResponseModel defines model for ErrorResponseModel.
type ErrorResponseModel struct {
	ErrorReason *string `json:"error_reason,omitempty"`
	Message     *string `json:"message,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}

// ImportKeypairPayload defines model for Import Keypair Payload.
type ImportKeypairPayload struct {
	// EnvironmentName The name of the environment where the key pair is being created.
	EnvironmentName string `json:"environment_name"`

	// Name The name of the key pair that is being created.
	Name string `json:"name"`

	// PublicKey The public key that is being used to import an SSH key pair.
	PublicKey string `json:"public_key"`
}

// ImportedKeypairResponseAPIObject defines model for ImportedKeypairResponseAPIObject.
type ImportedKeypairResponseAPIObject struct {
	Keypair *KeypairFields `json:"keypair,omitempty"`
	Message *string        `json:"message,omitempty"`
	Status  *bool          `json:"status,omitempty"`
}

// KeypairFields defines model for Keypair Fields.
type KeypairFields struct {
	CreatedAt   *time.CustomTime `json:"created_at,omitempty"`
	Environment *string    `json:"environment,omitempty"`
	Fingerprint *string    `json:"fingerprint,omitempty"`
	Id          *int       `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	PublicKey   *string    `json:"public_key,omitempty"`
}

// Keypairs defines model for Keypairs.
type Keypairs struct {
	Keypairs *[]KeypairFields `json:"Keypairs,omitempty"`
	Message  *string          `json:"message,omitempty"`
	Status   *bool            `json:"status,omitempty"`
}

// ResponseModel defines model for ResponseModel.
type ResponseModel struct {
	Message *string `json:"message,omitempty"`
	Status  *bool   `json:"status,omitempty"`
}

// UpdateKeypairName defines model for Update Keypair Name.
type UpdateKeypairName struct {
	// Name The new key pair name.
	Name string `json:"name"`
}

// UpdatedKeypairNameResponseAPIObject defines model for UpdatedKeypairNameResponseAPIObject.
type UpdatedKeypairNameResponseAPIObject struct {
	Keypair *KeypairFields `json:"keypair,omitempty"`
	Message *string        `json:"message,omitempty"`
	Status  *bool          `json:"status,omitempty"`
}

// UpdateKeyPairNameJSONRequestBody defines body for UpdateKeyPairName for application/json ContentType.
type UpdateKeyPairNameJSONRequestBody = UpdateKeypairName

// ImportKeyPairJSONRequestBody defines body for ImportKeyPair for application/json ContentType.
type ImportKeyPairJSONRequestBody = ImportKeypairPayload

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
	// DeleteKeyPair request
	DeleteKeyPair(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateKeyPairNameWithBody request with any body
	UpdateKeyPairNameWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateKeyPairName(ctx context.Context, id int, body UpdateKeyPairNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListKeyPairs request
	ListKeyPairs(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ImportKeyPairWithBody request with any body
	ImportKeyPairWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	ImportKeyPair(ctx context.Context, body ImportKeyPairJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) DeleteKeyPair(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteKeyPairRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateKeyPairNameWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateKeyPairNameRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateKeyPairName(ctx context.Context, id int, body UpdateKeyPairNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateKeyPairNameRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListKeyPairs(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListKeyPairsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ImportKeyPairWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewImportKeyPairRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ImportKeyPair(ctx context.Context, body ImportKeyPairJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewImportKeyPairRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewDeleteKeyPairRequest generates requests for DeleteKeyPair
func NewDeleteKeyPairRequest(server string, id int) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/core/keypair/%s", pathParam0)
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

// NewUpdateKeyPairNameRequest calls the generic UpdateKeyPairName builder with application/json body
func NewUpdateKeyPairNameRequest(server string, id int, body UpdateKeyPairNameJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateKeyPairNameRequestWithBody(server, id, "application/json", bodyReader)
}

// NewUpdateKeyPairNameRequestWithBody generates requests for UpdateKeyPairName with any type of body
func NewUpdateKeyPairNameRequestWithBody(server string, id int, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/core/keypair/%s", pathParam0)
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

// NewListKeyPairsRequest generates requests for ListKeyPairs
func NewListKeyPairsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/core/keypairs")
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

// NewImportKeyPairRequest calls the generic ImportKeyPair builder with application/json body
func NewImportKeyPairRequest(server string, body ImportKeyPairJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewImportKeyPairRequestWithBody(server, "application/json", bodyReader)
}

// NewImportKeyPairRequestWithBody generates requests for ImportKeyPair with any type of body
func NewImportKeyPairRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/core/keypairs")
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
	// DeleteKeyPairWithResponse request
	DeleteKeyPairWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteKeyPairResponse, error)

	// UpdateKeyPairNameWithBodyWithResponse request with any body
	UpdateKeyPairNameWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateKeyPairNameResponse, error)

	UpdateKeyPairNameWithResponse(ctx context.Context, id int, body UpdateKeyPairNameJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateKeyPairNameResponse, error)

	// ListKeyPairsWithResponse request
	ListKeyPairsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListKeyPairsResponse, error)

	// ImportKeyPairWithBodyWithResponse request with any body
	ImportKeyPairWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ImportKeyPairResponse, error)

	ImportKeyPairWithResponse(ctx context.Context, body ImportKeyPairJSONRequestBody, reqEditors ...RequestEditorFn) (*ImportKeyPairResponse, error)
}

type DeleteKeyPairResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ResponseModel
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r DeleteKeyPairResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteKeyPairResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateKeyPairNameResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *UpdatedKeypairNameResponseAPIObject
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r UpdateKeyPairNameResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateKeyPairNameResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListKeyPairsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Keypairs
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r ListKeyPairsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListKeyPairsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ImportKeyPairResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ImportedKeypairResponseAPIObject
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
	JSON409      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r ImportKeyPairResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ImportKeyPairResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// DeleteKeyPairWithResponse request returning *DeleteKeyPairResponse
func (c *ClientWithResponses) DeleteKeyPairWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteKeyPairResponse, error) {
	rsp, err := c.DeleteKeyPair(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteKeyPairResponse(rsp)
}

// UpdateKeyPairNameWithBodyWithResponse request with arbitrary body returning *UpdateKeyPairNameResponse
func (c *ClientWithResponses) UpdateKeyPairNameWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateKeyPairNameResponse, error) {
	rsp, err := c.UpdateKeyPairNameWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateKeyPairNameResponse(rsp)
}

func (c *ClientWithResponses) UpdateKeyPairNameWithResponse(ctx context.Context, id int, body UpdateKeyPairNameJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateKeyPairNameResponse, error) {
	rsp, err := c.UpdateKeyPairName(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateKeyPairNameResponse(rsp)
}

// ListKeyPairsWithResponse request returning *ListKeyPairsResponse
func (c *ClientWithResponses) ListKeyPairsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListKeyPairsResponse, error) {
	rsp, err := c.ListKeyPairs(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListKeyPairsResponse(rsp)
}

// ImportKeyPairWithBodyWithResponse request with arbitrary body returning *ImportKeyPairResponse
func (c *ClientWithResponses) ImportKeyPairWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ImportKeyPairResponse, error) {
	rsp, err := c.ImportKeyPairWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseImportKeyPairResponse(rsp)
}

func (c *ClientWithResponses) ImportKeyPairWithResponse(ctx context.Context, body ImportKeyPairJSONRequestBody, reqEditors ...RequestEditorFn) (*ImportKeyPairResponse, error) {
	rsp, err := c.ImportKeyPair(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseImportKeyPairResponse(rsp)
}

// ParseDeleteKeyPairResponse parses an HTTP response from a DeleteKeyPairWithResponse call
func ParseDeleteKeyPairResponse(rsp *http.Response) (*DeleteKeyPairResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteKeyPairResponse{
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

// ParseUpdateKeyPairNameResponse parses an HTTP response from a UpdateKeyPairNameWithResponse call
func ParseUpdateKeyPairNameResponse(rsp *http.Response) (*UpdateKeyPairNameResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateKeyPairNameResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest UpdatedKeypairNameResponseAPIObject
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

// ParseListKeyPairsResponse parses an HTTP response from a ListKeyPairsWithResponse call
func ParseListKeyPairsResponse(rsp *http.Response) (*ListKeyPairsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListKeyPairsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Keypairs
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

// ParseImportKeyPairResponse parses an HTTP response from a ImportKeyPairWithResponse call
func ParseImportKeyPairResponse(rsp *http.Response) (*ImportKeyPairResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ImportKeyPairResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ImportedKeypairResponseAPIObject
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

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	}

	return response, nil
}
