// Package assigning_member_role provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package assigning_member_role

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

// AssignRbacRolePayload defines model for AssignRbacRolePayload.
type AssignRbacRolePayload struct {
	RoleId int `json:"role_id"`
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

// RbacRoleDetailResponseModel defines model for RbacRoleDetailResponseModel.
type RbacRoleDetailResponseModel struct {
	Message *string         `json:"message,omitempty"`
	Role    *RbacRoleFields `json:"role,omitempty"`
	Status  *bool           `json:"status,omitempty"`
}

// RbacRoleFields defines model for RbacRoleFields.
type RbacRoleFields struct {
	CreatedAt   *time.CustomTime              `json:"created_at,omitempty"`
	Description *string                 `json:"description,omitempty"`
	Id          *int                    `json:"id,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	Permissions *[]RolePermissionFields `json:"permissions,omitempty"`
	Policies    *[]RolePolicyFields     `json:"policies,omitempty"`
}

// RolePermissionFields defines model for RolePermissionFields.
type RolePermissionFields struct {
	Id         *int    `json:"id,omitempty"`
	Permission *string `json:"permission,omitempty"`
	Resource   *string `json:"resource,omitempty"`
}

// RolePolicyFields defines model for RolePolicyFields.
type RolePolicyFields struct {
	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// AssignRBACRolesJSONRequestBody defines body for AssignRBACRoles for application/json ContentType.
type AssignRBACRolesJSONRequestBody = AssignRbacRolePayload

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
	// AssignRBACRolesWithBody request with any body
	AssignRBACRolesWithBody(ctx context.Context, userId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AssignRBACRoles(ctx context.Context, userId int, body AssignRBACRolesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RemoveRoleFromAUser request
	RemoveRoleFromAUser(ctx context.Context, userId int, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) AssignRBACRolesWithBody(ctx context.Context, userId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAssignRBACRolesRequestWithBody(c.Server, userId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AssignRBACRoles(ctx context.Context, userId int, body AssignRBACRolesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAssignRBACRolesRequest(c.Server, userId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RemoveRoleFromAUser(ctx context.Context, userId int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRemoveRoleFromAUserRequest(c.Server, userId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewAssignRBACRolesRequest calls the generic AssignRBACRoles builder with application/json body
func NewAssignRBACRolesRequest(server string, userId int, body AssignRBACRolesJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAssignRBACRolesRequestWithBody(server, userId, "application/json", bodyReader)
}

// NewAssignRBACRolesRequestWithBody generates requests for AssignRBACRoles with any type of body
func NewAssignRBACRolesRequestWithBody(server string, userId int, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/auth/users/%s/assign-roles", pathParam0)
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

// NewRemoveRoleFromAUserRequest generates requests for RemoveRoleFromAUser
func NewRemoveRoleFromAUserRequest(server string, userId int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/auth/users/%s/roles", pathParam0)
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
	// AssignRBACRolesWithBodyWithResponse request with any body
	AssignRBACRolesWithBodyWithResponse(ctx context.Context, userId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AssignRBACRolesResponse, error)

	AssignRBACRolesWithResponse(ctx context.Context, userId int, body AssignRBACRolesJSONRequestBody, reqEditors ...RequestEditorFn) (*AssignRBACRolesResponse, error)

	// RemoveRoleFromAUserWithResponse request
	RemoveRoleFromAUserWithResponse(ctx context.Context, userId int, reqEditors ...RequestEditorFn) (*RemoveRoleFromAUserResponse, error)
}

type AssignRBACRolesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RbacRoleDetailResponseModel
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r AssignRBACRolesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AssignRBACRolesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RemoveRoleFromAUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CommonResponseModel
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r RemoveRoleFromAUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RemoveRoleFromAUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// AssignRBACRolesWithBodyWithResponse request with arbitrary body returning *AssignRBACRolesResponse
func (c *ClientWithResponses) AssignRBACRolesWithBodyWithResponse(ctx context.Context, userId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AssignRBACRolesResponse, error) {
	rsp, err := c.AssignRBACRolesWithBody(ctx, userId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAssignRBACRolesResponse(rsp)
}

func (c *ClientWithResponses) AssignRBACRolesWithResponse(ctx context.Context, userId int, body AssignRBACRolesJSONRequestBody, reqEditors ...RequestEditorFn) (*AssignRBACRolesResponse, error) {
	rsp, err := c.AssignRBACRoles(ctx, userId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAssignRBACRolesResponse(rsp)
}

// RemoveRoleFromAUserWithResponse request returning *RemoveRoleFromAUserResponse
func (c *ClientWithResponses) RemoveRoleFromAUserWithResponse(ctx context.Context, userId int, reqEditors ...RequestEditorFn) (*RemoveRoleFromAUserResponse, error) {
	rsp, err := c.RemoveRoleFromAUser(ctx, userId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRemoveRoleFromAUserResponse(rsp)
}

// ParseAssignRBACRolesResponse parses an HTTP response from a AssignRBACRolesWithResponse call
func ParseAssignRBACRolesResponse(rsp *http.Response) (*AssignRBACRolesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AssignRBACRolesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RbacRoleDetailResponseModel
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

// ParseRemoveRoleFromAUserResponse parses an HTTP response from a RemoveRoleFromAUserWithResponse call
func ParseRemoveRoleFromAUserResponse(rsp *http.Response) (*RemoveRoleFromAUserResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RemoveRoleFromAUserResponse{
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
