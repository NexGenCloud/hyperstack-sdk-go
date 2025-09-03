# \VncUrlAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVncUrl**](VncUrlAPI.md#GetVncUrl) | **Get** /core/virtual-machines/{vm_id}/console/{job_id} | Get VNC Console Link
[**GetVncUrl2**](VncUrlAPI.md#GetVncUrl2) | **Get** /core/virtual-machines/{vm_id}/request-console | Request Instance Console



## GetVncUrl

> VNCURL GetVncUrl(ctx, vmId, jobId).Execute()

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
	vmId := int32(56) // int32 | 
	jobId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VncUrlAPI.GetVncUrl(context.Background(), vmId, jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VncUrlAPI.GetVncUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVncUrl`: VNCURL
	fmt.Fprintf(os.Stdout, "Response from `VncUrlAPI.GetVncUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 
**jobId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVncUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VNCURL**](VNCURL.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVncUrl2

> RequestConsole GetVncUrl2(ctx, vmId).Execute()

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
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVncUrl2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestConsole**](RequestConsole.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

