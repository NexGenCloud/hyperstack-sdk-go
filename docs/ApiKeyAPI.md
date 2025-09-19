# \ApiKeyAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAPIKey**](ApiKeyAPI.md#DeleteAPIKey) | **Delete** /api-key/{api_key_id} | Delete API Key
[**GenerateAPIKey**](ApiKeyAPI.md#GenerateAPIKey) | **Post** /api-key/generate | Generate API Key
[**RetrieveAPIKey**](ApiKeyAPI.md#RetrieveAPIKey) | **Get** /api-key | Retrieve API Keys
[**UpdateAPIKey**](ApiKeyAPI.md#UpdateAPIKey) | **Put** /api-key/{api_key_id} | Update API Key



## DeleteAPIKey

> CommonResponseModel DeleteAPIKey(ctx, apiKeyId).Execute()

Delete API Key



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
	apiKeyId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.DeleteAPIKey(context.Background(), apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.DeleteAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAPIKey`: CommonResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.DeleteAPIKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommonResponseModel**](CommonResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateAPIKey

> GenerateUpdateApiKeyResponseModel GenerateAPIKey(ctx).Payload(payload).Execute()

Generate API Key



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
	payload := *openapiclient.NewGenerateUpdateApiKeyPayload("Name_example") // GenerateUpdateApiKeyPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.GenerateAPIKey(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.GenerateAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateAPIKey`: GenerateUpdateApiKeyResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.GenerateAPIKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**GenerateUpdateApiKeyPayload**](GenerateUpdateApiKeyPayload.md) |  | 

### Return type

[**GenerateUpdateApiKeyResponseModel**](GenerateUpdateApiKeyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAPIKey

> GetApiKeysResponseModel RetrieveAPIKey(ctx).Execute()

Retrieve API Keys



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.RetrieveAPIKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.RetrieveAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveAPIKey`: GetApiKeysResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.RetrieveAPIKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAPIKeyRequest struct via the builder pattern


### Return type

[**GetApiKeysResponseModel**](GetApiKeysResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAPIKey

> GenerateUpdateApiKeyResponseModel UpdateAPIKey(ctx, apiKeyId).Payload(payload).Execute()

Update API Key



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
	apiKeyId := int32(56) // int32 | 
	payload := *openapiclient.NewGenerateUpdateApiKeyPayload("Name_example") // GenerateUpdateApiKeyPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.UpdateAPIKey(context.Background(), apiKeyId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.UpdateAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAPIKey`: GenerateUpdateApiKeyResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.UpdateAPIKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**GenerateUpdateApiKeyPayload**](GenerateUpdateApiKeyPayload.md) |  | 

### Return type

[**GenerateUpdateApiKeyResponseModel**](GenerateUpdateApiKeyResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

