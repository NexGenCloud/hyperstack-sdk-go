// Package clusters provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package clusters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/NexGenCloud/hyperstack-sdk-go/lib/time"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// ClusterFields defines model for ClusterFields.
type ClusterFields struct {
	ApiAddress        *string               `json:"api_address,omitempty"`
	CreatedAt         *time.CustomTime      `json:"created_at,omitempty"`
	EnvironmentName   *string               `json:"environment_name,omitempty"`
	Id                *int                  `json:"id,omitempty"`
	KeypairName       *string               `json:"keypair_name,omitempty"`
	KubeConfig        *string               `json:"kube_config,omitempty"`
	KubernetesVersion *string               `json:"kubernetes_version,omitempty"`
	Name              *string               `json:"name,omitempty"`
	NodeCount         *int                  `json:"node_count,omitempty"`
	NodeFlavor        *InstanceFlavorFields `json:"node_flavor,omitempty"`
	Status            *string               `json:"status,omitempty"`
	StatusReason      *string               `json:"status_reason,omitempty"`
}

// ClusterListResponse defines model for ClusterListResponse.
type ClusterListResponse struct {
	Clusters *[]ClusterFields `json:"clusters,omitempty"`
	Message  *string          `json:"message,omitempty"`
	Status   *bool            `json:"status,omitempty"`
}

// ClusterResponse defines model for ClusterResponse.
type ClusterResponse struct {
	Cluster *ClusterFields `json:"cluster,omitempty"`
	Message *string        `json:"message,omitempty"`
	Status  *bool          `json:"status,omitempty"`
}

// ClusterVersions defines model for ClusterVersions.
type ClusterVersions struct {
	Message  *string   `json:"message,omitempty"`
	Status   *bool     `json:"status,omitempty"`
	Versions *[]string `json:"versions,omitempty"`
}

// CreateClusterPayload defines model for CreateClusterPayload.
type CreateClusterPayload struct {
	EnvironmentName   string `json:"environment_name"`
	ImageName         string `json:"image_name"`
	KeypairName       string `json:"keypair_name"`
	KubernetesVersion string `json:"kubernetes_version"`
	MasterFlavorName  string `json:"master_flavor_name"`
	Name              string `json:"name"`
	NodeCount         int    `json:"node_count"`
	NodeFlavorName    string `json:"node_flavor_name"`
}

// ErrorResponseModel defines model for ErrorResponseModel.
type ErrorResponseModel struct {
	ErrorReason *string `json:"error_reason,omitempty"`
	Message     *string `json:"message,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}

// InstanceFlavorFields defines model for InstanceFlavorFields.
type InstanceFlavorFields struct {
	Cpu       *int     `json:"cpu,omitempty"`
	Disk      *int     `json:"disk,omitempty"`
	Ephemeral *int     `json:"ephemeral,omitempty"`
	Gpu       *string  `json:"gpu,omitempty"`
	GpuCount  *int     `json:"gpu_count,omitempty"`
	Id        *int     `json:"id,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Ram       *float32 `json:"ram,omitempty"`
}

// ResponseModel defines model for ResponseModel.
type ResponseModel struct {
	Message *string `json:"message,omitempty"`
	Status  *bool   `json:"status,omitempty"`
}

// CreateClusterJSONRequestBody defines body for CreateCluster for application/json ContentType.
type CreateClusterJSONRequestBody = CreateClusterPayload

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
	// ListClusters request
	ListClusters(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateClusterWithBody request with any body
	CreateClusterWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateCluster(ctx context.Context, body CreateClusterJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetClusterVersions request
	GetClusterVersions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteACluster request
	DeleteACluster(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GettingClusterDetail request
	GettingClusterDetail(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ListClusters(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListClustersRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateClusterWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateClusterRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateCluster(ctx context.Context, body CreateClusterJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateClusterRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetClusterVersions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetClusterVersionsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteACluster(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteAClusterRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GettingClusterDetail(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGettingClusterDetailRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListClustersRequest generates requests for ListClusters
func NewListClustersRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/core/clusters")
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

// NewCreateClusterRequest calls the generic CreateCluster builder with application/json body
func NewCreateClusterRequest(server string, body CreateClusterJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateClusterRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateClusterRequestWithBody generates requests for CreateCluster with any type of body
func NewCreateClusterRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/core/clusters")
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

// NewGetClusterVersionsRequest generates requests for GetClusterVersions
func NewGetClusterVersionsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/core/clusters/versions")
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

// NewDeleteAClusterRequest generates requests for DeleteACluster
func NewDeleteAClusterRequest(server string, id int) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/core/clusters/%s", pathParam0)
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

// NewGettingClusterDetailRequest generates requests for GettingClusterDetail
func NewGettingClusterDetailRequest(server string, id int) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/core/clusters/%s", pathParam0)
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
	// ListClustersWithResponse request
	ListClustersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListClustersResponse, error)

	// CreateClusterWithBodyWithResponse request with any body
	CreateClusterWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateClusterResponse, error)

	CreateClusterWithResponse(ctx context.Context, body CreateClusterJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateClusterResponse, error)

	// GetClusterVersionsWithResponse request
	GetClusterVersionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetClusterVersionsResponse, error)

	// DeleteAClusterWithResponse request
	DeleteAClusterWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteAClusterResponse, error)

	// GettingClusterDetailWithResponse request
	GettingClusterDetailWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GettingClusterDetailResponse, error)
}

type ListClustersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ClusterListResponse
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r ListClustersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListClustersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateClusterResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *ClusterResponse
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
	JSON409      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r CreateClusterResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateClusterResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetClusterVersionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ClusterVersions
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r GetClusterVersionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetClusterVersionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteAClusterResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ResponseModel
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r DeleteAClusterResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteAClusterResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GettingClusterDetailResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ClusterResponse
	JSON400      *ErrorResponseModel
	JSON401      *ErrorResponseModel
	JSON404      *ErrorResponseModel
}

// Status returns HTTPResponse.Status
func (r GettingClusterDetailResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GettingClusterDetailResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListClustersWithResponse request returning *ListClustersResponse
func (c *ClientWithResponses) ListClustersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListClustersResponse, error) {
	rsp, err := c.ListClusters(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListClustersResponse(rsp)
}

// CreateClusterWithBodyWithResponse request with arbitrary body returning *CreateClusterResponse
func (c *ClientWithResponses) CreateClusterWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateClusterResponse, error) {
	rsp, err := c.CreateClusterWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateClusterResponse(rsp)
}

func (c *ClientWithResponses) CreateClusterWithResponse(ctx context.Context, body CreateClusterJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateClusterResponse, error) {
	rsp, err := c.CreateCluster(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateClusterResponse(rsp)
}

// GetClusterVersionsWithResponse request returning *GetClusterVersionsResponse
func (c *ClientWithResponses) GetClusterVersionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetClusterVersionsResponse, error) {
	rsp, err := c.GetClusterVersions(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetClusterVersionsResponse(rsp)
}

// DeleteAClusterWithResponse request returning *DeleteAClusterResponse
func (c *ClientWithResponses) DeleteAClusterWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteAClusterResponse, error) {
	rsp, err := c.DeleteACluster(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteAClusterResponse(rsp)
}

// GettingClusterDetailWithResponse request returning *GettingClusterDetailResponse
func (c *ClientWithResponses) GettingClusterDetailWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GettingClusterDetailResponse, error) {
	rsp, err := c.GettingClusterDetail(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGettingClusterDetailResponse(rsp)
}

// ParseListClustersResponse parses an HTTP response from a ListClustersWithResponse call
func ParseListClustersResponse(rsp *http.Response) (*ListClustersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListClustersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ClusterListResponse
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

// ParseCreateClusterResponse parses an HTTP response from a CreateClusterWithResponse call
func ParseCreateClusterResponse(rsp *http.Response) (*CreateClusterResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateClusterResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest ClusterResponse
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

// ParseGetClusterVersionsResponse parses an HTTP response from a GetClusterVersionsWithResponse call
func ParseGetClusterVersionsResponse(rsp *http.Response) (*GetClusterVersionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetClusterVersionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ClusterVersions
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

// ParseDeleteAClusterResponse parses an HTTP response from a DeleteAClusterWithResponse call
func ParseDeleteAClusterResponse(rsp *http.Response) (*DeleteAClusterResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteAClusterResponse{
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

// ParseGettingClusterDetailResponse parses an HTTP response from a GettingClusterDetailWithResponse call
func ParseGettingClusterDetailResponse(rsp *http.Response) (*GettingClusterDetailResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GettingClusterDetailResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ClusterResponse
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
