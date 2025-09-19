# \BetaAccessAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateABetaAccessRequest**](BetaAccessAPI.md#CreateABetaAccessRequest) | **Post** /auth/beta-access/requests | Create a new beta access request
[**GetBetaAccessStatus**](BetaAccessAPI.md#GetBetaAccessStatus) | **Get** /auth/beta-access/requests | Check the status of all beta access requests
[**GetBetaAccessStatus2**](BetaAccessAPI.md#GetBetaAccessStatus2) | **Get** /auth/beta-access/requests/{program} | Check the status of beta access requests



## CreateABetaAccessRequest

> BetaAccessRequestResponseModel CreateABetaAccessRequest(ctx).Payload(payload).Execute()

Create a new beta access request



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
	payload := *openapiclient.NewBetaAccessRequestPayload("Program_example") // BetaAccessRequestPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAccessAPI.CreateABetaAccessRequest(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAccessAPI.CreateABetaAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateABetaAccessRequest`: BetaAccessRequestResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BetaAccessAPI.CreateABetaAccessRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateABetaAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**BetaAccessRequestPayload**](BetaAccessRequestPayload.md) |  | 

### Return type

[**BetaAccessRequestResponseModel**](BetaAccessRequestResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBetaAccessStatus

> BetaAccessStatusResponseModel GetBetaAccessStatus(ctx).Execute()

Check the status of all beta access requests



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
	resp, r, err := apiClient.BetaAccessAPI.GetBetaAccessStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAccessAPI.GetBetaAccessStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBetaAccessStatus`: BetaAccessStatusResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BetaAccessAPI.GetBetaAccessStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBetaAccessStatusRequest struct via the builder pattern


### Return type

[**BetaAccessStatusResponseModel**](BetaAccessStatusResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBetaAccessStatus2

> BetaAccessStatusResponseModel GetBetaAccessStatus2(ctx, program).Execute()

Check the status of beta access requests



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
	program := "program_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAccessAPI.GetBetaAccessStatus2(context.Background(), program).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAccessAPI.GetBetaAccessStatus2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBetaAccessStatus2`: BetaAccessStatusResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BetaAccessAPI.GetBetaAccessStatus2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**program** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBetaAccessStatus2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BetaAccessStatusResponseModel**](BetaAccessStatusResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

