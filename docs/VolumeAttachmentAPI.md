# \VolumeAttachmentAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**AttachVolumesToVirtualMachine**](VolumeAttachmentAPI.md#AttachVolumesToVirtualMachine) | **Post** /core/virtual-machines/{vm_id}/attach-volumes | Attach volumes to virtual machine
[**DetachVolumesFromVirtualMachine**](VolumeAttachmentAPI.md#DetachVolumesFromVirtualMachine) | **Post** /core/virtual-machines/{vm_id}/detach-volumes | Detach volumes from virtual machine
[**UpdateAVolumeAttachment**](VolumeAttachmentAPI.md#UpdateAVolumeAttachment) | **Patch** /core/volume-attachments/{volume_attachment_id} | Update a volume attachment
=======
[**AttachVolumesToVirtualMachine**](VolumeAttachmentAPI.md#AttachVolumesToVirtualMachine) | **Post** /core/virtual-machines/{virtual_machine_id}/attach-volumes | Attach volumes to virtual machine
[**DetachVolumesFromVirtualMachine**](VolumeAttachmentAPI.md#DetachVolumesFromVirtualMachine) | **Post** /core/virtual-machines/{virtual_machine_id}/detach-volumes | Detach volumes from virtual machine
>>>>>>> main



## AttachVolumesToVirtualMachine

<<<<<<< HEAD
> AttachVolumes AttachVolumesToVirtualMachine(ctx, vmId).Payload(payload).Execute()
=======
> AttachVolumes AttachVolumesToVirtualMachine(ctx, virtualMachineId).Payload(payload).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	vmId := int32(56) // int32 | 
=======
	virtualMachineId := int32(56) // int32 | 
>>>>>>> main
	payload := *openapiclient.NewAttachVolumesPayload() // AttachVolumesPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.VolumeAttachmentAPI.AttachVolumesToVirtualMachine(context.Background(), vmId).Payload(payload).Execute()
=======
	resp, r, err := apiClient.VolumeAttachmentAPI.AttachVolumesToVirtualMachine(context.Background(), virtualMachineId).Payload(payload).Execute()
>>>>>>> main
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAttachmentAPI.AttachVolumesToVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachVolumesToVirtualMachine`: AttachVolumes
	fmt.Fprintf(os.Stdout, "Response from `VolumeAttachmentAPI.AttachVolumesToVirtualMachine`: %v\n", resp)
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

### Other Parameters

Other parameters are passed through a pointer to a apiAttachVolumesToVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AttachVolumesPayload**](AttachVolumesPayload.md) |  | 

### Return type

[**AttachVolumes**](AttachVolumes.md)

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


## DetachVolumesFromVirtualMachine

<<<<<<< HEAD
> DetachVolumes DetachVolumesFromVirtualMachine(ctx, vmId).Payload(payload).Execute()
=======
> DetachVolumes DetachVolumesFromVirtualMachine(ctx, virtualMachineId).Payload(payload).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	vmId := int32(56) // int32 | 
=======
	virtualMachineId := int32(56) // int32 | 
>>>>>>> main
	payload := *openapiclient.NewDetachVolumesPayload() // DetachVolumesPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.VolumeAttachmentAPI.DetachVolumesFromVirtualMachine(context.Background(), vmId).Payload(payload).Execute()
=======
	resp, r, err := apiClient.VolumeAttachmentAPI.DetachVolumesFromVirtualMachine(context.Background(), virtualMachineId).Payload(payload).Execute()
>>>>>>> main
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAttachmentAPI.DetachVolumesFromVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachVolumesFromVirtualMachine`: DetachVolumes
	fmt.Fprintf(os.Stdout, "Response from `VolumeAttachmentAPI.DetachVolumesFromVirtualMachine`: %v\n", resp)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDetachVolumesFromVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**DetachVolumesPayload**](DetachVolumesPayload.md) |  | 

### Return type

[**DetachVolumes**](DetachVolumes.md)

### Authorization

<<<<<<< HEAD
[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAVolumeAttachment

> AttachVolumes UpdateAVolumeAttachment(ctx, volumeAttachmentId).Payload(payload).Execute()

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
	resp, r, err := apiClient.VolumeAttachmentAPI.UpdateAVolumeAttachment(context.Background(), volumeAttachmentId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAttachmentAPI.UpdateAVolumeAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAVolumeAttachment`: AttachVolumes
	fmt.Fprintf(os.Stdout, "Response from `VolumeAttachmentAPI.UpdateAVolumeAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeAttachmentId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAVolumeAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**UpdateVolumeAttachmentPayload**](UpdateVolumeAttachmentPayload.md) |  | 

### Return type

[**AttachVolumes**](AttachVolumes.md)

### Authorization

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

