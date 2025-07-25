// Package user provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package user

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
)

// AddUserInfoSuccessResponseModel defines model for AddUserInfoSuccessResponseModel.
type AddUserInfoSuccessResponseModel struct {
	Data    *UsersInfoFields `json:"data,omitempty"`
	Message *string          `json:"message,omitempty"`
	Status  *bool            `json:"status,omitempty"`
}

// ErrorResponseModel defines model for ErrorResponseModel.
type ErrorResponseModel struct {
	ErrorReason *string `json:"error_reason,omitempty"`
	Message     *string `json:"message,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}

// Userinfopostpayload defines model for Userinfopostpayload.
type Userinfopostpayload struct {
	BillingAddress1 *string `json:"billing_address1,omitempty"`
	BillingAddress2 *string `json:"billing_address2,omitempty"`
	Business        bool    `json:"business"`
	CompanyName     *string `json:"company_name,omitempty"`
	Country         string  `json:"country"`
	Email           *string `json:"email,omitempty"`
	Name            *string `json:"name,omitempty"`
	Phone           *string `json:"phone,omitempty"`
	State           *string `json:"state,omitempty"`
	VatNumber       *string `json:"vat_number,omitempty"`
	ZipCode         string  `json:"zip_code"`
}

// UsersInfoFields defines model for UsersInfoFields.
type UsersInfoFields struct {
	BillingAddress1 *string    `json:"billing_address1,omitempty"`
	BillingAddress2 *string    `json:"billing_address2,omitempty"`
	Business        *bool      `json:"business,omitempty"`
	CompanyName     *string    `json:"company_name,omitempty"`
	Country         *string    `json:"country,omitempty"`
	CreatedAt       *time.CustomTime `json:"created_at,omitempty"`
	Email           *string    `json:"email,omitempty"`
	Id              *int       `json:"id,omitempty"`
	Name            *string    `json:"name,omitempty"`
	OrganizationId  *int       `json:"organization_id,omitempty"`
	Phone           *string    `json:"phone,omitempty"`
	State           *string    `json:"state,omitempty"`
	StripeUserId    *string    `json:"stripe_user_id,omitempty"`
	VatNumber       *string    `json:"vat_number,omitempty"`
	ZipCode         *string    `json:"zip_code,omitempty"`
}

// UsersInfoListResponse defines model for UsersInfoListResponse.
type UsersInfoListResponse struct {
	Message   *string          `json:"message,omitempty"`
	Status    *bool            `json:"status,omitempty"`
	UsersInfo *UsersInfoFields `json:"users_info,omitempty"`
}

// PostInsertBillingInfoJSONRequestBody defines body for PostInsertBillingInfo for application/json ContentType.
type PostInsertBillingInfoJSONRequestBody = Userinfopostpayload

// PutUpdateBillingInfoJSONRequestBody defines body for PutUpdateBillingInfo for application/json ContentType.
type PutUpdateBillingInfoJSONRequestBody = Userinfopostpayload

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
	// GetRetrieveBillingInfo request
	GetRetrieveBillingInfo(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostInsertBillingInfoWithBody request with any body
	PostInsertBillingInfoWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostInsertBillingInfo(ctx context.Context, body PostInsertBillingInfoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutUpdateBillingInfoWithBody request with any body
	PutUpdateBillingInfoWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutUpdateBillingInfo(ctx context.Context, body PutUpdateBillingInfoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetRetrieveBillingInfo(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRetrieveBillingInfoRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostInsertBillingInfoWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostInsertBillingInfoRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostInsertBillingInfo(ctx context.Context, body PostInsertBillingInfoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostInsertBillingInfoRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutUpdateBillingInfoWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutUpdateBillingInfoRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutUpdateBillingInfo(ctx context.Context, body PutUpdateBillingInfoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutUpdateBillingInfoRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetRetrieveBillingInfoRequest generates requests for GetRetrieveBillingInfo
func NewGetRetrieveBillingInfoRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/billing/user/info")
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

// NewPostInsertBillingInfoRequest calls the generic PostInsertBillingInfo builder with application/json body
func NewPostInsertBillingInfoRequest(server string, body PostInsertBillingInfoJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostInsertBillingInfoRequestWithBody(server, "application/json", bodyReader)
}

// NewPostInsertBillingInfoRequestWithBody generates requests for PostInsertBillingInfo with any type of body
func NewPostInsertBillingInfoRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/billing/user/info")
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

// NewPutUpdateBillingInfoRequest calls the generic PutUpdateBillingInfo builder with application/json body
func NewPutUpdateBillingInfoRequest(server string, body PutUpdateBillingInfoJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutUpdateBillingInfoRequestWithBody(server, "application/json", bodyReader)
}

// NewPutUpdateBillingInfoRequestWithBody generates requests for PutUpdateBillingInfo with any type of body
func NewPutUpdateBillingInfoRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/billing/user/info")
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
	// GetRetrieveBillingInfoWithResponse request
	GetRetrieveBillingInfoWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRetrieveBillingInfoResponse, error)

	// PostInsertBillingInfoWithBodyWithResponse request with any body
	PostInsertBillingInfoWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostInsertBillingInfoResponse, error)

	PostInsertBillingInfoWithResponse(ctx context.Context, body PostInsertBillingInfoJSONRequestBody, reqEditors ...RequestEditorFn) (*PostInsertBillingInfoResponse, error)

	// PutUpdateBillingInfoWithBodyWithResponse request with any body
	PutUpdateBillingInfoWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutUpdateBillingInfoResponse, error)

	PutUpdateBillingInfoWithResponse(ctx context.Context, body PutUpdateBillingInfoJSONRequestBody, reqEditors ...RequestEditorFn) (*PutUpdateBillingInfoResponse, error)
}

type GetRetrieveBillingInfoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *UsersInfoListResponse
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON403      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r GetRetrieveBillingInfoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetRetrieveBillingInfoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostInsertBillingInfoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AddUserInfoSuccessResponseModel
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON403      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r PostInsertBillingInfoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostInsertBillingInfoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PutUpdateBillingInfoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AddUserInfoSuccessResponseModel
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON403      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r PutUpdateBillingInfoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PutUpdateBillingInfoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetRetrieveBillingInfoWithResponse request returning *GetRetrieveBillingInfoResponse
func (c *ClientWithResponses) GetRetrieveBillingInfoWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRetrieveBillingInfoResponse, error) {
	rsp, err := c.GetRetrieveBillingInfo(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRetrieveBillingInfoResponse(rsp)
}

// PostInsertBillingInfoWithBodyWithResponse request with arbitrary body returning *PostInsertBillingInfoResponse
func (c *ClientWithResponses) PostInsertBillingInfoWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostInsertBillingInfoResponse, error) {
	rsp, err := c.PostInsertBillingInfoWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostInsertBillingInfoResponse(rsp)
}

func (c *ClientWithResponses) PostInsertBillingInfoWithResponse(ctx context.Context, body PostInsertBillingInfoJSONRequestBody, reqEditors ...RequestEditorFn) (*PostInsertBillingInfoResponse, error) {
	rsp, err := c.PostInsertBillingInfo(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostInsertBillingInfoResponse(rsp)
}

// PutUpdateBillingInfoWithBodyWithResponse request with arbitrary body returning *PutUpdateBillingInfoResponse
func (c *ClientWithResponses) PutUpdateBillingInfoWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutUpdateBillingInfoResponse, error) {
	rsp, err := c.PutUpdateBillingInfoWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutUpdateBillingInfoResponse(rsp)
}

func (c *ClientWithResponses) PutUpdateBillingInfoWithResponse(ctx context.Context, body PutUpdateBillingInfoJSONRequestBody, reqEditors ...RequestEditorFn) (*PutUpdateBillingInfoResponse, error) {
	rsp, err := c.PutUpdateBillingInfo(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutUpdateBillingInfoResponse(rsp)
}

// ParseGetRetrieveBillingInfoResponse parses an HTTP response from a GetRetrieveBillingInfoWithResponse call
func ParseGetRetrieveBillingInfoResponse(rsp *http.Response) (*GetRetrieveBillingInfoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetRetrieveBillingInfoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest UsersInfoListResponse
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

// ParsePostInsertBillingInfoResponse parses an HTTP response from a PostInsertBillingInfoWithResponse call
func ParsePostInsertBillingInfoResponse(rsp *http.Response) (*PostInsertBillingInfoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostInsertBillingInfoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AddUserInfoSuccessResponseModel
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

// ParsePutUpdateBillingInfoResponse parses an HTTP response from a PutUpdateBillingInfoWithResponse call
func ParsePutUpdateBillingInfoResponse(rsp *http.Response) (*PutUpdateBillingInfoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutUpdateBillingInfoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AddUserInfoSuccessResponseModel
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
