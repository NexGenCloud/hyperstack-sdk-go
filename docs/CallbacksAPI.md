# \CallbacksAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachCallbackToVirtualMachine**](CallbacksAPI.md#AttachCallbackToVirtualMachine) | **Post** /core/virtual-machines/{vm_id}/attach-callback | Attach callback to virtual machine
[**AttachCallbackToVolume**](CallbacksAPI.md#AttachCallbackToVolume) | **Post** /core/volumes/{volume_id}/attach-callback | Attach callback to volume
[**DeleteVirtualMachineCallback**](CallbacksAPI.md#DeleteVirtualMachineCallback) | **Delete** /core/virtual-machines/{vm_id}/delete-callback | Delete virtual machine callback
[**DeleteVolumeCallback**](CallbacksAPI.md#DeleteVolumeCallback) | **Delete** /core/volumes/{volume_id}/delete-callback | Delete volume callback
[**UpdateVirtualMachineCallback**](CallbacksAPI.md#UpdateVirtualMachineCallback) | **Put** /core/virtual-machines/{vm_id}/update-callback | Update virtual machine callback
[**UpdateVolumeCallback**](CallbacksAPI.md#UpdateVolumeCallback) | **Put** /core/volumes/{volume_id}/update-callback | Update volume callback



## AttachCallbackToVirtualMachine

> AttachCallbackResponse AttachCallbackToVirtualMachine(ctx, vmId).Payload(payload).Execute()

Attach callback to virtual machine



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
	payload := *openapiclient.NewAttachCallbackPayload("Url_example") // AttachCallbackPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallbacksAPI.AttachCallbackToVirtualMachine(context.Background(), vmId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallbacksAPI.AttachCallbackToVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachCallbackToVirtualMachine`: AttachCallbackResponse
	fmt.Fprintf(os.Stdout, "Response from `CallbacksAPI.AttachCallbackToVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachCallbackToVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AttachCallbackPayload**](AttachCallbackPayload.md) |  | 

### Return type

[**AttachCallbackResponse**](AttachCallbackResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachCallbackToVolume

> AttachCallbackResponse AttachCallbackToVolume(ctx, volumeId).Payload(payload).Execute()

Attach callback to volume



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
	volumeId := int32(56) // int32 | 
	payload := *openapiclient.NewAttachCallbackPayload("Url_example") // AttachCallbackPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallbacksAPI.AttachCallbackToVolume(context.Background(), volumeId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallbacksAPI.AttachCallbackToVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachCallbackToVolume`: AttachCallbackResponse
	fmt.Fprintf(os.Stdout, "Response from `CallbacksAPI.AttachCallbackToVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachCallbackToVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AttachCallbackPayload**](AttachCallbackPayload.md) |  | 

### Return type

[**AttachCallbackResponse**](AttachCallbackResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualMachineCallback

> ResponseModel DeleteVirtualMachineCallback(ctx, vmId).Execute()

Delete virtual machine callback



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
	resp, r, err := apiClient.CallbacksAPI.DeleteVirtualMachineCallback(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallbacksAPI.DeleteVirtualMachineCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVirtualMachineCallback`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `CallbacksAPI.DeleteVirtualMachineCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualMachineCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolumeCallback

> ResponseModel DeleteVolumeCallback(ctx, volumeId).Execute()

Delete volume callback



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
	volumeId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallbacksAPI.DeleteVolumeCallback(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallbacksAPI.DeleteVolumeCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVolumeCallback`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `CallbacksAPI.DeleteVolumeCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualMachineCallback

> AttachCallbackResponse UpdateVirtualMachineCallback(ctx, vmId).Payload(payload).Execute()

Update virtual machine callback



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
	payload := *openapiclient.NewAttachCallbackPayload("Url_example") // AttachCallbackPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallbacksAPI.UpdateVirtualMachineCallback(context.Background(), vmId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallbacksAPI.UpdateVirtualMachineCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVirtualMachineCallback`: AttachCallbackResponse
	fmt.Fprintf(os.Stdout, "Response from `CallbacksAPI.UpdateVirtualMachineCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualMachineCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AttachCallbackPayload**](AttachCallbackPayload.md) |  | 

### Return type

[**AttachCallbackResponse**](AttachCallbackResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVolumeCallback

> AttachCallbackResponse UpdateVolumeCallback(ctx, volumeId).Payload(payload).Execute()

Update volume callback



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
	volumeId := int32(56) // int32 | 
	payload := *openapiclient.NewAttachCallbackPayload("Url_example") // AttachCallbackPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallbacksAPI.UpdateVolumeCallback(context.Background(), volumeId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallbacksAPI.UpdateVolumeCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVolumeCallback`: AttachCallbackResponse
	fmt.Fprintf(os.Stdout, "Response from `CallbacksAPI.UpdateVolumeCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AttachCallbackPayload**](AttachCallbackPayload.md) |  | 

### Return type

[**AttachCallbackResponse**](AttachCallbackResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

