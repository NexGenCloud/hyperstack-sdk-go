// Package organization provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package organization

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"github.com/nexgen/hyperstack-sdk-go/lib/time"
)

// ErrorResponseModel defines model for ErrorResponseModel.
type ErrorResponseModel struct {
	ErrorReason *string `json:"error_reason,omitempty"`
	Message     *string `json:"message,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}

// OrganizationInfoModel defines model for Organization Info Model.
type OrganizationInfoModel struct {
	CreatedAt *time.CustomTime               `json:"created_at,omitempty"`
	Id        int                      `json:"id"`
	Name      string                   `json:"name"`
	Users     *[]OrganizationUserModel `json:"users,omitempty"`
}

// OrganizationModel defines model for Organization Model.
type OrganizationModel struct {
	Message      *string                `json:"message,omitempty"`
	Organization *OrganizationInfoModel `json:"organization,omitempty"`
	Status       *bool                  `json:"status,omitempty"`
}

// OrganizationUserModel defines model for Organization User Model.
type OrganizationUserModel struct {
	Email     *string                         `json:"email,omitempty"`
	Id        *int                            `json:"id,omitempty"`
	JoinedAt  *time.CustomTime                      `json:"joined_at,omitempty"`
	Name      *string                         `json:"name,omitempty"`
	RbacRoles *[]RBACRoleFieldForOrganization `json:"rbac_roles,omitempty"`
	Role      *string                         `json:"role,omitempty"`
	Sub       *string                         `json:"sub,omitempty"`
	Username  *string                         `json:"username,omitempty"`
}

// RBACRoleFieldForOrganization defines model for RBAC Role Field for Organization.
type RBACRoleFieldForOrganization struct {
	Name *string `json:"name,omitempty"`
}

// RemoveMemberResponse defines model for Remove Member Response.
type RemoveMemberResponse struct {
	Message *string `json:"message,omitempty"`
	Status  *bool   `json:"status,omitempty"`
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
	// GetOrganizationInfo request
	GetOrganizationInfo(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RemoveAMemberFromOrganization request
	RemoveAMemberFromOrganization(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetOrganizationInfo(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetOrganizationInfoRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RemoveAMemberFromOrganization(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRemoveAMemberFromOrganizationRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetOrganizationInfoRequest generates requests for GetOrganizationInfo
func NewGetOrganizationInfoRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/auth/organizations")
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

// NewRemoveAMemberFromOrganizationRequest generates requests for RemoveAMemberFromOrganization
func NewRemoveAMemberFromOrganizationRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/auth/organizations/remove-member")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
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
	// GetOrganizationInfoWithResponse request
	GetOrganizationInfoWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetOrganizationInfoResponse, error)

	// RemoveAMemberFromOrganizationWithResponse request
	RemoveAMemberFromOrganizationWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RemoveAMemberFromOrganizationResponse, error)
}

type GetOrganizationInfoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *OrganizationModel
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r GetOrganizationInfoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetOrganizationInfoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RemoveAMemberFromOrganizationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RemoveMemberResponse
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r RemoveAMemberFromOrganizationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RemoveAMemberFromOrganizationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetOrganizationInfoWithResponse request returning *GetOrganizationInfoResponse
func (c *ClientWithResponses) GetOrganizationInfoWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetOrganizationInfoResponse, error) {
	rsp, err := c.GetOrganizationInfo(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetOrganizationInfoResponse(rsp)
}

// RemoveAMemberFromOrganizationWithResponse request returning *RemoveAMemberFromOrganizationResponse
func (c *ClientWithResponses) RemoveAMemberFromOrganizationWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RemoveAMemberFromOrganizationResponse, error) {
	rsp, err := c.RemoveAMemberFromOrganization(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRemoveAMemberFromOrganizationResponse(rsp)
}

// ParseGetOrganizationInfoResponse parses an HTTP response from a GetOrganizationInfoWithResponse call
func ParseGetOrganizationInfoResponse(rsp *http.Response) (*GetOrganizationInfoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetOrganizationInfoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OrganizationModel
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

// ParseRemoveAMemberFromOrganizationResponse parses an HTTP response from a RemoveAMemberFromOrganizationWithResponse call
func ParseRemoveAMemberFromOrganizationResponse(rsp *http.Response) (*RemoveAMemberFromOrganizationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RemoveAMemberFromOrganizationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RemoveMemberResponse
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
