# \VncUrlAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVncConsoleLink**](VncUrlAPI.md#GetVncConsoleLink) | **Get** /core/virtual-machines/{virtual_machine_id}/console/{job_id} | Get VNC Console Link
[**RequestInstanceConsole**](VncUrlAPI.md#RequestInstanceConsole) | **Get** /core/virtual-machines/{id}/request-console | Request Instance Console



## GetVncConsoleLink

> VNCURL GetVncConsoleLink(ctx, virtualMachineId, jobId).Execute()

Get VNC Console Link

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
	virtualMachineId := int32(56) // int32 | 
	jobId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VncUrlAPI.GetVncConsoleLink(context.Background(), virtualMachineId, jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VncUrlAPI.GetVncConsoleLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVncConsoleLink`: VNCURL
	fmt.Fprintf(os.Stdout, "Response from `VncUrlAPI.GetVncConsoleLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | **int32** |  | 
**jobId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVncConsoleLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VNCURL**](VNCURL.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestInstanceConsole

> RequestConsole RequestInstanceConsole(ctx, id).Execute()

Request Instance Console

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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VncUrlAPI.RequestInstanceConsole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VncUrlAPI.RequestInstanceConsole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestInstanceConsole`: RequestConsole
	fmt.Fprintf(os.Stdout, "Response from `VncUrlAPI.RequestInstanceConsole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestInstanceConsoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestConsole**](RequestConsole.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

