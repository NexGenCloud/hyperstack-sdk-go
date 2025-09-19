# \VncUrlAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**GetVncUrl**](VncUrlAPI.md#GetVncUrl) | **Get** /core/virtual-machines/{vm_id}/console/{job_id} | Get VNC Console Link
[**GetVncUrl2**](VncUrlAPI.md#GetVncUrl2) | **Get** /core/virtual-machines/{vm_id}/request-console | Request Instance Console



## GetVncUrl

> VNCURL GetVncUrl(ctx, vmId, jobId).Execute()

Get VNC Console Link



=======
[**GetVncConsoleLink**](VncUrlAPI.md#GetVncConsoleLink) | **Get** /core/virtual-machines/{virtual_machine_id}/console/{job_id} | Get VNC Console Link
[**RequestInstanceConsole**](VncUrlAPI.md#RequestInstanceConsole) | **Get** /core/virtual-machines/{id}/request-console | Request Instance Console



## GetVncConsoleLink

> VNCURL GetVncConsoleLink(ctx, virtualMachineId, jobId).Execute()

Get VNC Console Link

>>>>>>> main
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
<<<<<<< HEAD
	vmId := int32(56) // int32 | 
=======
	virtualMachineId := int32(56) // int32 | 
>>>>>>> main
	jobId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.VncUrlAPI.GetVncUrl(context.Background(), vmId, jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VncUrlAPI.GetVncUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVncUrl`: VNCURL
	fmt.Fprintf(os.Stdout, "Response from `VncUrlAPI.GetVncUrl`: %v\n", resp)
=======
	resp, r, err := apiClient.VncUrlAPI.GetVncConsoleLink(context.Background(), virtualMachineId, jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VncUrlAPI.GetVncConsoleLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVncConsoleLink`: VNCURL
	fmt.Fprintf(os.Stdout, "Response from `VncUrlAPI.GetVncConsoleLink`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
<<<<<<< HEAD
**vmId** | **int32** |  | 
=======
**virtualMachineId** | **int32** |  | 
>>>>>>> main
**jobId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGetVncUrlRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiGetVncConsoleLinkRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VNCURL**](VNCURL.md)

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
## GetVncUrl2

> RequestConsole GetVncUrl2(ctx, vmId).Execute()

Request Instance Console



=======
## RequestInstanceConsole

> RequestConsole RequestInstanceConsole(ctx, id).Execute()

Request Instance Console

>>>>>>> main
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
<<<<<<< HEAD
	vmId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VncUrlAPI.GetVncUrl2(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VncUrlAPI.GetVncUrl2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVncUrl2`: RequestConsole
	fmt.Fprintf(os.Stdout, "Response from `VncUrlAPI.GetVncUrl2`: %v\n", resp)
=======
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
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
<<<<<<< HEAD
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVncUrl2Request struct via the builder pattern
=======
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestInstanceConsoleRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestConsole**](RequestConsole.md)

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

