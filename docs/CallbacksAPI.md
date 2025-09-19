# \CallbacksAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**AttachCallbackToVirtualMachine**](CallbacksAPI.md#AttachCallbackToVirtualMachine) | **Post** /core/virtual-machines/{vm_id}/attach-callback | Attach callback to virtual machine
[**AttachCallbackToVolume**](CallbacksAPI.md#AttachCallbackToVolume) | **Post** /core/volumes/{volume_id}/attach-callback | Attach callback to volume
[**DeleteVirtualMachineCallback**](CallbacksAPI.md#DeleteVirtualMachineCallback) | **Delete** /core/virtual-machines/{vm_id}/delete-callback | Delete virtual machine callback
[**DeleteVolumeCallback**](CallbacksAPI.md#DeleteVolumeCallback) | **Delete** /core/volumes/{volume_id}/delete-callback | Delete volume callback
[**UpdateVirtualMachineCallback**](CallbacksAPI.md#UpdateVirtualMachineCallback) | **Put** /core/virtual-machines/{vm_id}/update-callback | Update virtual machine callback
[**UpdateVolumeCallback**](CallbacksAPI.md#UpdateVolumeCallback) | **Put** /core/volumes/{volume_id}/update-callback | Update volume callback
=======
[**AttachCallbackToVirtualMachine**](CallbacksAPI.md#AttachCallbackToVirtualMachine) | **Post** /core/virtual-machines/{id}/attach-callback | Attach callback to virtual machine
[**AttachCallbackToVolume**](CallbacksAPI.md#AttachCallbackToVolume) | **Post** /core/volumes/{id}/attach-callback | Attach callback to volume
[**DeleteVirtualMachineCallback**](CallbacksAPI.md#DeleteVirtualMachineCallback) | **Delete** /core/virtual-machines/{id}/delete-callback | Delete virtual machine callback
[**DeleteVolumeCallback**](CallbacksAPI.md#DeleteVolumeCallback) | **Delete** /core/volumes/{id}/delete-callback | Delete volume callback
[**UpdateVirtualMachineCallback**](CallbacksAPI.md#UpdateVirtualMachineCallback) | **Put** /core/virtual-machines/{id}/update-callback | Update virtual machine callback
[**UpdateVolumeCallback**](CallbacksAPI.md#UpdateVolumeCallback) | **Put** /core/volumes/{id}/update-callback | Update volume callback
>>>>>>> main



## AttachCallbackToVirtualMachine

<<<<<<< HEAD
> AttachCallbackResponse AttachCallbackToVirtualMachine(ctx, vmId).Payload(payload).Execute()
=======
> AttachCallbackResponse AttachCallbackToVirtualMachine(ctx, id).Payload(payload).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	vmId := int32(56) // int32 | 
=======
	id := int32(56) // int32 | 
>>>>>>> main
	payload := *openapiclient.NewAttachCallbackPayload("Url_example") // AttachCallbackPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.CallbacksAPI.AttachCallbackToVirtualMachine(context.Background(), vmId).Payload(payload).Execute()
=======
	resp, r, err := apiClient.CallbacksAPI.AttachCallbackToVirtualMachine(context.Background(), id).Payload(payload).Execute()
>>>>>>> main
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
<<<<<<< HEAD
**vmId** | **int32** |  | 
=======
**id** | **int32** |  | 
>>>>>>> main

### Other Parameters

Other parameters are passed through a pointer to a apiAttachCallbackToVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AttachCallbackPayload**](AttachCallbackPayload.md) |  | 

### Return type

[**AttachCallbackResponse**](AttachCallbackResponse.md)

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


## AttachCallbackToVolume

<<<<<<< HEAD
> AttachCallbackResponse AttachCallbackToVolume(ctx, volumeId).Payload(payload).Execute()
=======
> AttachCallbackResponse AttachCallbackToVolume(ctx, id).Payload(payload).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	volumeId := int32(56) // int32 | 
=======
	id := int32(56) // int32 | 
>>>>>>> main
	payload := *openapiclient.NewAttachCallbackPayload("Url_example") // AttachCallbackPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.CallbacksAPI.AttachCallbackToVolume(context.Background(), volumeId).Payload(payload).Execute()
=======
	resp, r, err := apiClient.CallbacksAPI.AttachCallbackToVolume(context.Background(), id).Payload(payload).Execute()
>>>>>>> main
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
<<<<<<< HEAD
**volumeId** | **int32** |  | 
=======
**id** | **int32** |  | 
>>>>>>> main

### Other Parameters

Other parameters are passed through a pointer to a apiAttachCallbackToVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AttachCallbackPayload**](AttachCallbackPayload.md) |  | 

### Return type

[**AttachCallbackResponse**](AttachCallbackResponse.md)

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


## DeleteVirtualMachineCallback

<<<<<<< HEAD
> ResponseModel DeleteVirtualMachineCallback(ctx, vmId).Execute()
=======
> ResponseModel DeleteVirtualMachineCallback(ctx, id).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	vmId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallbacksAPI.DeleteVirtualMachineCallback(context.Background(), vmId).Execute()
=======
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallbacksAPI.DeleteVirtualMachineCallback(context.Background(), id).Execute()
>>>>>>> main
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
<<<<<<< HEAD
**vmId** | **int32** |  | 
=======
**id** | **int32** |  | 
>>>>>>> main

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualMachineCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseModel**](ResponseModel.md)

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


## DeleteVolumeCallback

<<<<<<< HEAD
> ResponseModel DeleteVolumeCallback(ctx, volumeId).Execute()
=======
> ResponseModel DeleteVolumeCallback(ctx, id).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	volumeId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallbacksAPI.DeleteVolumeCallback(context.Background(), volumeId).Execute()
=======
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallbacksAPI.DeleteVolumeCallback(context.Background(), id).Execute()
>>>>>>> main
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
<<<<<<< HEAD
**volumeId** | **int32** |  | 
=======
**id** | **int32** |  | 
>>>>>>> main

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseModel**](ResponseModel.md)

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


## UpdateVirtualMachineCallback

<<<<<<< HEAD
> AttachCallbackResponse UpdateVirtualMachineCallback(ctx, vmId).Payload(payload).Execute()
=======
> AttachCallbackResponse UpdateVirtualMachineCallback(ctx, id).Payload(payload).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	vmId := int32(56) // int32 | 
=======
	id := int32(56) // int32 | 
>>>>>>> main
	payload := *openapiclient.NewAttachCallbackPayload("Url_example") // AttachCallbackPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.CallbacksAPI.UpdateVirtualMachineCallback(context.Background(), vmId).Payload(payload).Execute()
=======
	resp, r, err := apiClient.CallbacksAPI.UpdateVirtualMachineCallback(context.Background(), id).Payload(payload).Execute()
>>>>>>> main
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
<<<<<<< HEAD
**vmId** | **int32** |  | 
=======
**id** | **int32** |  | 
>>>>>>> main

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualMachineCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AttachCallbackPayload**](AttachCallbackPayload.md) |  | 

### Return type

[**AttachCallbackResponse**](AttachCallbackResponse.md)

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


## UpdateVolumeCallback

<<<<<<< HEAD
> AttachCallbackResponse UpdateVolumeCallback(ctx, volumeId).Payload(payload).Execute()
=======
> AttachCallbackResponse UpdateVolumeCallback(ctx, id).Payload(payload).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	volumeId := int32(56) // int32 | 
=======
	id := int32(56) // int32 | 
>>>>>>> main
	payload := *openapiclient.NewAttachCallbackPayload("Url_example") // AttachCallbackPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.CallbacksAPI.UpdateVolumeCallback(context.Background(), volumeId).Payload(payload).Execute()
=======
	resp, r, err := apiClient.CallbacksAPI.UpdateVolumeCallback(context.Background(), id).Payload(payload).Execute()
>>>>>>> main
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
<<<<<<< HEAD
**volumeId** | **int32** |  | 
=======
**id** | **int32** |  | 
>>>>>>> main

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AttachCallbackPayload**](AttachCallbackPayload.md) |  | 

### Return type

[**AttachCallbackResponse**](AttachCallbackResponse.md)

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

