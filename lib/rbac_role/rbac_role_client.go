// Package rbac_role provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package rbac_role

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

// RBACRole defines model for RBAC Role.
type RBACRole struct {
	Message *string         `json:"message,omitempty"`
	Role    *RBACRoleFields `json:"role,omitempty"`
	Status  *bool           `json:"status,omitempty"`
}

// RBACRoleFields defines model for RBAC Role Fields.
type RBACRoleFields struct {
	CreatedAt   *time.CustomTime              `json:"created_at,omitempty"`
	Description *string                 `json:"description,omitempty"`
	Id          *int                    `json:"id,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	Permissions *[]RolePermissionFields `json:"permissions,omitempty"`
	Policies    *[]RolePolicyFields     `json:"policies,omitempty"`
}

// RBACRolePayload defines model for RBAC Role Payload.
type RBACRolePayload struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Permissions *[]int `json:"permissions,omitempty"`
	Policies    *[]int `json:"policies,omitempty"`
}

// RBACRoles defines model for RBAC Roles.
type RBACRoles struct {
	Message *string           `json:"message,omitempty"`
	Roles   *[]RBACRoleFields `json:"roles,omitempty"`
	Status  *bool             `json:"status,omitempty"`
}

// RBACRoleTMPGet defines model for RBACRoleTMPGet.
type RBACRoleTMPGet struct {
	Message *string         `json:"message,omitempty"`
	Roles   *RBACRoleFields `json:"roles,omitempty"`
	Status  *bool           `json:"status,omitempty"`
}

// ResponseModel defines model for ResponseModel.
type ResponseModel struct {
	Message *string `json:"message,omitempty"`
	Status  *bool   `json:"status,omitempty"`
}

// RolePermissionFields defines model for Role Permission Fields.
type RolePermissionFields struct {
	Id         *int    `json:"id,omitempty"`
	Permission *string `json:"permission,omitempty"`
	Resource   *string `json:"resource,omitempty"`
}

// RolePolicyFields defines model for Role Policy Fields.
type RolePolicyFields struct {
	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// CreateRBACRoleJSONRequestBody defines body for CreateRBACRole for application/json ContentType.
type CreateRBACRoleJSONRequestBody = RBACRolePayload

// UpdateARBACRoleJSONRequestBody defines body for UpdateARBACRole for application/json ContentType.
type UpdateARBACRoleJSONRequestBody = RBACRolePayload

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
	// ListRBACRoles request
	ListRBACRoles(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateRBACRoleWithBody request with any body
	CreateRBACRoleWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateRBACRole(ctx context.Context, body CreateRBACRoleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteARBACRole request
	DeleteARBACRole(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetARBACRoleDetail request
	GetARBACRoleDetail(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateARBACRoleWithBody request with any body
	UpdateARBACRoleWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateARBACRole(ctx context.Context, id int, body UpdateARBACRoleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ListRBACRoles(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListRBACRolesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateRBACRoleWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateRBACRoleRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateRBACRole(ctx context.Context, body CreateRBACRoleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateRBACRoleRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteARBACRole(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteARBACRoleRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetARBACRoleDetail(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetARBACRoleDetailRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateARBACRoleWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateARBACRoleRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateARBACRole(ctx context.Context, id int, body UpdateARBACRoleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateARBACRoleRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListRBACRolesRequest generates requests for ListRBACRoles
func NewListRBACRolesRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/auth/roles")
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

// NewCreateRBACRoleRequest calls the generic CreateRBACRole builder with application/json body
func NewCreateRBACRoleRequest(server string, body CreateRBACRoleJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateRBACRoleRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateRBACRoleRequestWithBody generates requests for CreateRBACRole with any type of body
func NewCreateRBACRoleRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/auth/roles")
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

// NewDeleteARBACRoleRequest generates requests for DeleteARBACRole
func NewDeleteARBACRoleRequest(server string, id int) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/auth/roles/%s", pathParam0)
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

// NewGetARBACRoleDetailRequest generates requests for GetARBACRoleDetail
func NewGetARBACRoleDetailRequest(server string, id int) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/auth/roles/%s", pathParam0)
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

// NewUpdateARBACRoleRequest calls the generic UpdateARBACRole builder with application/json body
func NewUpdateARBACRoleRequest(server string, id int, body UpdateARBACRoleJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateARBACRoleRequestWithBody(server, id, "application/json", bodyReader)
}

// NewUpdateARBACRoleRequestWithBody generates requests for UpdateARBACRole with any type of body
func NewUpdateARBACRoleRequestWithBody(server string, id int, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/auth/roles/%s", pathParam0)
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
	// ListRBACRolesWithResponse request
	ListRBACRolesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListRBACRolesResponse, error)

	// CreateRBACRoleWithBodyWithResponse request with any body
	CreateRBACRoleWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateRBACRoleResponse, error)

	CreateRBACRoleWithResponse(ctx context.Context, body CreateRBACRoleJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateRBACRoleResponse, error)

	// DeleteARBACRoleWithResponse request
	DeleteARBACRoleWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteARBACRoleResponse, error)

	// GetARBACRoleDetailWithResponse request
	GetARBACRoleDetailWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetARBACRoleDetailResponse, error)

	// UpdateARBACRoleWithBodyWithResponse request with any body
	UpdateARBACRoleWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateARBACRoleResponse, error)

	UpdateARBACRoleWithResponse(ctx context.Context, id int, body UpdateARBACRoleJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateARBACRoleResponse, error)
}

type ListRBACRolesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RBACRoles
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r ListRBACRolesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListRBACRolesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateRBACRoleResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *RBACRole
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON409      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r CreateRBACRoleResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateRBACRoleResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteARBACRoleResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ResponseModel
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r DeleteARBACRoleResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteARBACRoleResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetARBACRoleDetailResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RBACRoleTMPGet
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r GetARBACRoleDetailResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetARBACRoleDetailResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateARBACRoleResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RBACRole
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r UpdateARBACRoleResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateARBACRoleResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListRBACRolesWithResponse request returning *ListRBACRolesResponse
func (c *ClientWithResponses) ListRBACRolesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListRBACRolesResponse, error) {
	rsp, err := c.ListRBACRoles(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListRBACRolesResponse(rsp)
}

// CreateRBACRoleWithBodyWithResponse request with arbitrary body returning *CreateRBACRoleResponse
func (c *ClientWithResponses) CreateRBACRoleWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateRBACRoleResponse, error) {
	rsp, err := c.CreateRBACRoleWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateRBACRoleResponse(rsp)
}

func (c *ClientWithResponses) CreateRBACRoleWithResponse(ctx context.Context, body CreateRBACRoleJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateRBACRoleResponse, error) {
	rsp, err := c.CreateRBACRole(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateRBACRoleResponse(rsp)
}

// DeleteARBACRoleWithResponse request returning *DeleteARBACRoleResponse
func (c *ClientWithResponses) DeleteARBACRoleWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteARBACRoleResponse, error) {
	rsp, err := c.DeleteARBACRole(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteARBACRoleResponse(rsp)
}

// GetARBACRoleDetailWithResponse request returning *GetARBACRoleDetailResponse
func (c *ClientWithResponses) GetARBACRoleDetailWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetARBACRoleDetailResponse, error) {
	rsp, err := c.GetARBACRoleDetail(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetARBACRoleDetailResponse(rsp)
}

// UpdateARBACRoleWithBodyWithResponse request with arbitrary body returning *UpdateARBACRoleResponse
func (c *ClientWithResponses) UpdateARBACRoleWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateARBACRoleResponse, error) {
	rsp, err := c.UpdateARBACRoleWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateARBACRoleResponse(rsp)
}

func (c *ClientWithResponses) UpdateARBACRoleWithResponse(ctx context.Context, id int, body UpdateARBACRoleJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateARBACRoleResponse, error) {
	rsp, err := c.UpdateARBACRole(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateARBACRoleResponse(rsp)
}

// ParseListRBACRolesResponse parses an HTTP response from a ListRBACRolesWithResponse call
func ParseListRBACRolesResponse(rsp *http.Response) (*ListRBACRolesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListRBACRolesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RBACRoles
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

// ParseCreateRBACRoleResponse parses an HTTP response from a CreateRBACRoleWithResponse call
func ParseCreateRBACRoleResponse(rsp *http.Response) (*CreateRBACRoleResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateRBACRoleResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest RBACRole
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

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

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest ErrorResponseModel
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	}

	return response, nil
}

// ParseDeleteARBACRoleResponse parses an HTTP response from a DeleteARBACRoleWithResponse call
func ParseDeleteARBACRoleResponse(rsp *http.Response) (*DeleteARBACRoleResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteARBACRoleResponse{
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

// ParseGetARBACRoleDetailResponse parses an HTTP response from a GetARBACRoleDetailWithResponse call
func ParseGetARBACRoleDetailResponse(rsp *http.Response) (*GetARBACRoleDetailResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetARBACRoleDetailResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RBACRoleTMPGet
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

// ParseUpdateARBACRoleResponse parses an HTTP response from a UpdateARBACRoleWithResponse call
func ParseUpdateARBACRoleResponse(rsp *http.Response) (*UpdateARBACRoleResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateARBACRoleResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RBACRole
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
