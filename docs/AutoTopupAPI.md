# \AutoTopupAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutoTopUp**](AutoTopupAPI.md#CreateAutoTopUp) | **Post** /billing/auto-topup | Create an auto top up configuration and initiate Stripe setup
[**DisableAutoTopUp**](AutoTopupAPI.md#DisableAutoTopUp) | **Delete** /billing/auto-topup | Disable auto top up, preventing any future automatic charges
[**GetAutoTopUp**](AutoTopupAPI.md#GetAutoTopUp) | **Get** /billing/auto-topup | Retrieve the current auto top up configuration and transaction history
[**GetAutoTopUpStatus**](AutoTopupAPI.md#GetAutoTopUpStatus) | **Get** /billing/auto-topup/status | Get auto top-up status and configuration
[**UpdateAutoTopUp**](AutoTopupAPI.md#UpdateAutoTopUp) | **Put** /billing/auto-topup | Update an existing active auto top up configuration



## CreateAutoTopUp

> CreateAutoTopupResponse CreateAutoTopUp(ctx).Payload(payload).Execute()

Create an auto top up configuration and initiate Stripe setup

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
	payload := *openapiclient.NewCreateAutoTopupPayload(float32(123), float32(123)) // CreateAutoTopupPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTopupAPI.CreateAutoTopUp(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTopupAPI.CreateAutoTopUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAutoTopUp`: CreateAutoTopupResponse
	fmt.Fprintf(os.Stdout, "Response from `AutoTopupAPI.CreateAutoTopUp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoTopUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateAutoTopupPayload**](CreateAutoTopupPayload.md) |  | 

### Return type

[**CreateAutoTopupResponse**](CreateAutoTopupResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableAutoTopUp

> DisableAutoTopupResponse DisableAutoTopUp(ctx).Execute()

Disable auto top up, preventing any future automatic charges

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
	resp, r, err := apiClient.AutoTopupAPI.DisableAutoTopUp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTopupAPI.DisableAutoTopUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableAutoTopUp`: DisableAutoTopupResponse
	fmt.Fprintf(os.Stdout, "Response from `AutoTopupAPI.DisableAutoTopUp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisableAutoTopUpRequest struct via the builder pattern


### Return type

[**DisableAutoTopupResponse**](DisableAutoTopupResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoTopUp

> GetAutoTopupResponse GetAutoTopUp(ctx).Execute()

Retrieve the current auto top up configuration and transaction history

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
	resp, r, err := apiClient.AutoTopupAPI.GetAutoTopUp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTopupAPI.GetAutoTopUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoTopUp`: GetAutoTopupResponse
	fmt.Fprintf(os.Stdout, "Response from `AutoTopupAPI.GetAutoTopUp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoTopUpRequest struct via the builder pattern


### Return type

[**GetAutoTopupResponse**](GetAutoTopupResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoTopUpStatus

> AutoTopupStatusSchema GetAutoTopUpStatus(ctx).Execute()

Get auto top-up status and configuration



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
	resp, r, err := apiClient.AutoTopupAPI.GetAutoTopUpStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTopupAPI.GetAutoTopUpStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoTopUpStatus`: AutoTopupStatusSchema
	fmt.Fprintf(os.Stdout, "Response from `AutoTopupAPI.GetAutoTopUpStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoTopUpStatusRequest struct via the builder pattern


### Return type

[**AutoTopupStatusSchema**](AutoTopupStatusSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoTopUp

> UpdateAutoTopupResponse UpdateAutoTopUp(ctx).Payload(payload).Execute()

Update an existing active auto top up configuration

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
	payload := *openapiclient.NewUpdateAutoTopupPayload(float32(123), float32(123)) // UpdateAutoTopupPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTopupAPI.UpdateAutoTopUp(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTopupAPI.UpdateAutoTopUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoTopUp`: UpdateAutoTopupResponse
	fmt.Fprintf(os.Stdout, "Response from `AutoTopupAPI.UpdateAutoTopUp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoTopUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**UpdateAutoTopupPayload**](UpdateAutoTopupPayload.md) |  | 

### Return type

[**UpdateAutoTopupResponse**](UpdateAutoTopupResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

