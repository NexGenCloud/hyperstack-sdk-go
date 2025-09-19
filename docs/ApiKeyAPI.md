# \ApiKeyAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**DeleteAPIKey**](ApiKeyAPI.md#DeleteAPIKey) | **Delete** /api-key/{api_key_id} | Delete API Key
[**GenerateAPIKey**](ApiKeyAPI.md#GenerateAPIKey) | **Post** /api-key/generate | Generate API Key
[**RetrieveAPIKey**](ApiKeyAPI.md#RetrieveAPIKey) | **Get** /api-key | Retrieve API Keys
[**UpdateAPIKey**](ApiKeyAPI.md#UpdateAPIKey) | **Put** /api-key/{api_key_id} | Update API Key



## DeleteAPIKey

> CommonResponseModel DeleteAPIKey(ctx, apiKeyId).Execute()
=======
[**DeleteApiKey**](ApiKeyAPI.md#DeleteApiKey) | **Delete** /api-key/{api_key_id} | Delete API Key
[**GenerateApiKey**](ApiKeyAPI.md#GenerateApiKey) | **Post** /api-key/generate | Generate API Key
[**RetrieveApiKeys**](ApiKeyAPI.md#RetrieveApiKeys) | **Get** /api-key | Retrieve API Keys
[**UpdateApiKey**](ApiKeyAPI.md#UpdateApiKey) | **Put** /api-key/{api_key_id} | Update API Key



## DeleteApiKey

> CommonResponseModel DeleteApiKey(ctx, apiKeyId).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.ApiKeyAPI.DeleteAPIKey(context.Background(), apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.DeleteAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAPIKey`: CommonResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.DeleteAPIKey`: %v\n", resp)
=======
	resp, r, err := apiClient.ApiKeyAPI.DeleteApiKey(context.Background(), apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.DeleteApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiKey`: CommonResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.DeleteApiKey`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiDeleteAPIKeyRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiDeleteApiKeyRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommonResponseModel**](CommonResponseModel.md)

### Authorization

<<<<<<< HEAD
[apiKey](../README.md#apiKey)
=======
[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)
>>>>>>> main

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


<<<<<<< HEAD
## GenerateAPIKey

> GenerateUpdateApiKeyResponseModel GenerateAPIKey(ctx).Payload(payload).Execute()
=======
## GenerateApiKey

> GenerateUpdateApiKeyResponseModel GenerateApiKey(ctx).Payload(payload).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.ApiKeyAPI.GenerateAPIKey(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.GenerateAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateAPIKey`: GenerateUpdateApiKeyResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.GenerateAPIKey`: %v\n", resp)
=======
	resp, r, err := apiClient.ApiKeyAPI.GenerateApiKey(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.GenerateApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateApiKey`: GenerateUpdateApiKeyResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.GenerateApiKey`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters



### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGenerateAPIKeyRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiGenerateApiKeyRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**GenerateUpdateApiKeyPayload**](GenerateUpdateApiKeyPayload.md) |  | 

### Return type

[**GenerateUpdateApiKeyResponseModel**](GenerateUpdateApiKeyResponseModel.md)

### Authorization

<<<<<<< HEAD
No authorization required
=======
[accessToken](../README.md#accessToken)
>>>>>>> main

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


<<<<<<< HEAD
## RetrieveAPIKey

> GetApiKeysResponseModel RetrieveAPIKey(ctx).Execute()
=======
## RetrieveApiKeys

> GetApiKeysResponseModel RetrieveApiKeys(ctx).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.ApiKeyAPI.RetrieveAPIKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.RetrieveAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveAPIKey`: GetApiKeysResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.RetrieveAPIKey`: %v\n", resp)
=======
	resp, r, err := apiClient.ApiKeyAPI.RetrieveApiKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.RetrieveApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveApiKeys`: GetApiKeysResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.RetrieveApiKeys`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiRetrieveAPIKeyRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveApiKeysRequest struct via the builder pattern
>>>>>>> main


### Return type

[**GetApiKeysResponseModel**](GetApiKeysResponseModel.md)

### Authorization

<<<<<<< HEAD
[apiKey](../README.md#apiKey)
=======
[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)
>>>>>>> main

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


<<<<<<< HEAD
## UpdateAPIKey

> GenerateUpdateApiKeyResponseModel UpdateAPIKey(ctx, apiKeyId).Payload(payload).Execute()
=======
## UpdateApiKey

> GenerateUpdateApiKeyResponseModel UpdateApiKey(ctx, apiKeyId).Payload(payload).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.ApiKeyAPI.UpdateAPIKey(context.Background(), apiKeyId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.UpdateAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAPIKey`: GenerateUpdateApiKeyResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.UpdateAPIKey`: %v\n", resp)
=======
	resp, r, err := apiClient.ApiKeyAPI.UpdateApiKey(context.Background(), apiKeyId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.UpdateApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiKey`: GenerateUpdateApiKeyResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.UpdateApiKey`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiUpdateAPIKeyRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiUpdateApiKeyRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**GenerateUpdateApiKeyPayload**](GenerateUpdateApiKeyPayload.md) |  | 

### Return type

[**GenerateUpdateApiKeyResponseModel**](GenerateUpdateApiKeyResponseModel.md)

### Authorization

<<<<<<< HEAD
[apiKey](../README.md#apiKey)
=======
[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)
>>>>>>> main

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

