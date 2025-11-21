# \VolumeAttachmentAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachVolumesToVM**](VolumeAttachmentAPI.md#AttachVolumesToVM) | **Post** /core/virtual-machines/{vm_id}/attach-volumes | Attach volumes to virtual machine
[**DetachVolumesFromVM**](VolumeAttachmentAPI.md#DetachVolumesFromVM) | **Post** /core/virtual-machines/{vm_id}/detach-volumes | Detach volumes from virtual machine
[**UpdateVolumeAttachment**](VolumeAttachmentAPI.md#UpdateVolumeAttachment) | **Patch** /core/volume-attachments/{volume_attachment_id} | Update a volume attachment



## AttachVolumesToVM

> AttachVolumes AttachVolumesToVM(ctx, vmId).Payload(payload).Execute()

Attach volumes to virtual machine



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
	payload := *openapiclient.NewAttachVolumesPayload() // AttachVolumesPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAttachmentAPI.AttachVolumesToVM(context.Background(), vmId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAttachmentAPI.AttachVolumesToVM``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachVolumesToVM`: AttachVolumes
	fmt.Fprintf(os.Stdout, "Response from `VolumeAttachmentAPI.AttachVolumesToVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachVolumesToVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AttachVolumesPayload**](AttachVolumesPayload.md) |  | 

### Return type

[**AttachVolumes**](AttachVolumes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachVolumesFromVM

> DetachVolumes DetachVolumesFromVM(ctx, vmId).Payload(payload).Execute()

Detach volumes from virtual machine



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
	payload := *openapiclient.NewDetachVolumesPayload() // DetachVolumesPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAttachmentAPI.DetachVolumesFromVM(context.Background(), vmId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAttachmentAPI.DetachVolumesFromVM``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachVolumesFromVM`: DetachVolumes
	fmt.Fprintf(os.Stdout, "Response from `VolumeAttachmentAPI.DetachVolumesFromVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachVolumesFromVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**DetachVolumesPayload**](DetachVolumesPayload.md) |  | 

### Return type

[**DetachVolumes**](DetachVolumes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVolumeAttachment

> AttachVolumes UpdateVolumeAttachment(ctx, volumeAttachmentId).Payload(payload).Execute()

Update a volume attachment

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
	volumeAttachmentId := int32(56) // int32 | 
	payload := *openapiclient.NewUpdateVolumeAttachmentPayload() // UpdateVolumeAttachmentPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAttachmentAPI.UpdateVolumeAttachment(context.Background(), volumeAttachmentId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAttachmentAPI.UpdateVolumeAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVolumeAttachment`: AttachVolumes
	fmt.Fprintf(os.Stdout, "Response from `VolumeAttachmentAPI.UpdateVolumeAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeAttachmentId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**UpdateVolumeAttachmentPayload**](UpdateVolumeAttachmentPayload.md) |  | 

### Return type

[**AttachVolumes**](AttachVolumes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

