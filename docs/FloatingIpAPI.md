# \FloatingIpAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**AttachPublicIPToVirtualMachine**](FloatingIpAPI.md#AttachPublicIPToVirtualMachine) | **Post** /core/virtual-machines/{vm_id}/attach-floatingip | Attach public IP to virtual machine
[**DetachPublicIPFromVirtualMachine**](FloatingIpAPI.md#DetachPublicIPFromVirtualMachine) | **Post** /core/virtual-machines/{vm_id}/detach-floatingip | Detach public IP from virtual machine



## AttachPublicIPToVirtualMachine

> ResponseModel AttachPublicIPToVirtualMachine(ctx, vmId).Execute()
=======
[**AttachPublicIpToVirtualMachine**](FloatingIpAPI.md#AttachPublicIpToVirtualMachine) | **Post** /core/virtual-machines/{id}/attach-floatingip | Attach public IP to virtual machine
[**DetachPublicIpFromVirtualMachine**](FloatingIpAPI.md#DetachPublicIpFromVirtualMachine) | **Post** /core/virtual-machines/{id}/detach-floatingip | Detach public IP from virtual machine



## AttachPublicIpToVirtualMachine

> ResponseModel AttachPublicIpToVirtualMachine(ctx, id).Execute()
>>>>>>> main

Attach public IP to virtual machine



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
	resp, r, err := apiClient.FloatingIpAPI.AttachPublicIPToVirtualMachine(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FloatingIpAPI.AttachPublicIPToVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachPublicIPToVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FloatingIpAPI.AttachPublicIPToVirtualMachine`: %v\n", resp)
=======
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FloatingIpAPI.AttachPublicIpToVirtualMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FloatingIpAPI.AttachPublicIpToVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachPublicIpToVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FloatingIpAPI.AttachPublicIpToVirtualMachine`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAttachPublicIPToVirtualMachineRequest struct via the builder pattern
=======
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachPublicIpToVirtualMachineRequest struct via the builder pattern
>>>>>>> main


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


<<<<<<< HEAD
## DetachPublicIPFromVirtualMachine

> ResponseModel DetachPublicIPFromVirtualMachine(ctx, vmId).Execute()
=======
## DetachPublicIpFromVirtualMachine

> ResponseModel DetachPublicIpFromVirtualMachine(ctx, id).Execute()
>>>>>>> main

Detach public IP from virtual machine



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
	resp, r, err := apiClient.FloatingIpAPI.DetachPublicIPFromVirtualMachine(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FloatingIpAPI.DetachPublicIPFromVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachPublicIPFromVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FloatingIpAPI.DetachPublicIPFromVirtualMachine`: %v\n", resp)
=======
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FloatingIpAPI.DetachPublicIpFromVirtualMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FloatingIpAPI.DetachPublicIpFromVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachPublicIpFromVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FloatingIpAPI.DetachPublicIpFromVirtualMachine`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDetachPublicIPFromVirtualMachineRequest struct via the builder pattern
=======
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachPublicIpFromVirtualMachineRequest struct via the builder pattern
>>>>>>> main


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

