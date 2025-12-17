# \AccessKeysAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessKeyEndpoint**](AccessKeysAPI.md#CreateAccessKeyEndpoint) | **Post** /object-storage/access-keys | Generate a new access key
[**DeleteAccessKeyEndpoint**](AccessKeysAPI.md#DeleteAccessKeyEndpoint) | **Delete** /object-storage/access-keys/{access_key_id} | Remove an existing access key
[**ListAccessKeysEndpoint**](AccessKeysAPI.md#ListAccessKeysEndpoint) | **Get** /object-storage/access-keys | List access keys



## CreateAccessKeyEndpoint

> ObjectStorageAccessKeyCreateResponse CreateAccessKeyEndpoint(ctx).Body(body).Execute()

Generate a new access key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/NexGenCloud/hyperstack-sdk-go/hyperstack"
)

func main() {
	body := *openapiclient.NewObjectStorageAccessKeyCreateRequest(openapiclient.object_storage__RegionsEnum("CANADA-1")) // ObjectStorageAccessKeyCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessKeysAPI.CreateAccessKeyEndpoint(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessKeysAPI.CreateAccessKeyEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccessKeyEndpoint`: ObjectStorageAccessKeyCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessKeysAPI.CreateAccessKeyEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessKeyEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ObjectStorageAccessKeyCreateRequest**](ObjectStorageAccessKeyCreateRequest.md) |  | 

### Return type

[**ObjectStorageAccessKeyCreateResponse**](ObjectStorageAccessKeyCreateResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessKeyEndpoint

> ObjectStorageDeleteResponse DeleteAccessKeyEndpoint(ctx, accessKeyId).Execute()

Remove an existing access key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/NexGenCloud/hyperstack-sdk-go/hyperstack"
)

func main() {
	accessKeyId := "accessKeyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessKeysAPI.DeleteAccessKeyEndpoint(context.Background(), accessKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessKeysAPI.DeleteAccessKeyEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccessKeyEndpoint`: ObjectStorageDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessKeysAPI.DeleteAccessKeyEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessKeyEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStorageDeleteResponse**](ObjectStorageDeleteResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessKeysEndpoint

> ObjectStorageAccessKeyListResponse ListAccessKeysEndpoint(ctx).Search(search).Page(page).PageSize(pageSize).Execute()

List access keys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/NexGenCloud/hyperstack-sdk-go/hyperstack"
)

func main() {
	search := "search_example" // string |  (optional)
	page := "page_example" // string |  (optional)
	pageSize := "pageSize_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessKeysAPI.ListAccessKeysEndpoint(context.Background()).Search(search).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessKeysAPI.ListAccessKeysEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccessKeysEndpoint`: ObjectStorageAccessKeyListResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessKeysAPI.ListAccessKeysEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessKeysEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 
 **page** | **string** |  | 
 **pageSize** | **string** |  | 

### Return type

[**ObjectStorageAccessKeyListResponse**](ObjectStorageAccessKeyListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

