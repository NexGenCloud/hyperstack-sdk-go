# \VirtualMachineEventsAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**ListVirtualMachineEvents**](VirtualMachineEventsAPI.md#ListVirtualMachineEvents) | **Get** /core/virtual-machines/{vm_id}/events | List virtual machine events
=======
[**ListVirtualMachineEvents**](VirtualMachineEventsAPI.md#ListVirtualMachineEvents) | **Get** /core/virtual-machines/{virtual_machine_id}/events | List virtual machine events
>>>>>>> main



## ListVirtualMachineEvents

<<<<<<< HEAD
> InstanceEvents ListVirtualMachineEvents(ctx, vmId).Execute()
=======
> InstanceEvents ListVirtualMachineEvents(ctx, virtualMachineId).Execute()
>>>>>>> main

List virtual machine events



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
	resp, r, err := apiClient.VirtualMachineEventsAPI.ListVirtualMachineEvents(context.Background(), vmId).Execute()
=======
	virtualMachineId := "virtualMachineId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineEventsAPI.ListVirtualMachineEvents(context.Background(), virtualMachineId).Execute()
>>>>>>> main
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineEventsAPI.ListVirtualMachineEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVirtualMachineEvents`: InstanceEvents
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineEventsAPI.ListVirtualMachineEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
<<<<<<< HEAD
**vmId** | **int32** |  | 
=======
**virtualMachineId** | **string** |  | 
>>>>>>> main

### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualMachineEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InstanceEvents**](InstanceEvents.md)

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

