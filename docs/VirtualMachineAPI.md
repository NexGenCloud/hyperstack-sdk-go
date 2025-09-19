# \VirtualMachineAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**DeleteInstance**](VirtualMachineAPI.md#DeleteInstance) | **Delete** /core/virtual-machines/{vm_id} | Delete virtual machine
[**DeleteSecurityRule**](VirtualMachineAPI.md#DeleteSecurityRule) | **Delete** /core/virtual-machines/{vm_id}/sg-rules/{sg_rule_id} | Delete firewall rule from virtual machine
[**FetchVirtualMachineNameAvailability**](VirtualMachineAPI.md#FetchVirtualMachineNameAvailability) | **Get** /core/virtual-machines/name-availability/{name} | Fetch virtual machine name availability
[**GetContractInstances**](VirtualMachineAPI.md#GetContractInstances) | **Get** /core/virtual-machines/contract/{contract_id}/virtual-machines | Retrieve virtual machines associated with a contract
[**GetInstance**](VirtualMachineAPI.md#GetInstance) | **Get** /core/virtual-machines | List virtual machines
[**GetInstance2**](VirtualMachineAPI.md#GetInstance2) | **Get** /core/virtual-machines/{vm_id} | Retrieve virtual machine details
[**GetInstance3**](VirtualMachineAPI.md#GetInstance3) | **Get** /core/virtual-machines/{vm_id}/hard-reboot | Hard reboot virtual machine
[**GetInstance4**](VirtualMachineAPI.md#GetInstance4) | **Get** /core/virtual-machines/{vm_id}/start | Start virtual machine
[**GetInstance5**](VirtualMachineAPI.md#GetInstance5) | **Get** /core/virtual-machines/{vm_id}/stop | Stop virtual machine
[**GetInstanceHibernate**](VirtualMachineAPI.md#GetInstanceHibernate) | **Get** /core/virtual-machines/{vm_id}/hibernate | Hibernate virtual machine
[**GetInstanceHibernateRestore**](VirtualMachineAPI.md#GetInstanceHibernateRestore) | **Get** /core/virtual-machines/{vm_id}/hibernate-restore | Restore virtual machine from hibernation
[**GetInstanceLogs**](VirtualMachineAPI.md#GetInstanceLogs) | **Get** /core/virtual-machines/{vm_id}/logs | Get virtual machine logs
[**GetInstanceMetrics**](VirtualMachineAPI.md#GetInstanceMetrics) | **Get** /core/virtual-machines/{vm_id}/metrics | Retrieve virtual machine performance metrics
[**PostInstance**](VirtualMachineAPI.md#PostInstance) | **Post** /core/virtual-machines | Create virtual machines
[**PostInstanceAttachFirewalls**](VirtualMachineAPI.md#PostInstanceAttachFirewalls) | **Post** /core/virtual-machines/{vm_id}/attach-firewalls | Attach firewalls to a virtual machine
[**PostInstanceLogs**](VirtualMachineAPI.md#PostInstanceLogs) | **Post** /core/virtual-machines/{vm_id}/logs | Request virtual machine logs
[**PostInstanceResize**](VirtualMachineAPI.md#PostInstanceResize) | **Post** /core/virtual-machines/{vm_id}/resize | Resize virtual machine
[**PostSecurityRule**](VirtualMachineAPI.md#PostSecurityRule) | **Post** /core/virtual-machines/{vm_id}/sg-rules | Add firewall rule to virtual machine
[**PostSnapshots**](VirtualMachineAPI.md#PostSnapshots) | **Post** /core/virtual-machines/{vm_id}/snapshots | Create snapshot from a virtual machine
[**PutLabels**](VirtualMachineAPI.md#PutLabels) | **Put** /core/virtual-machines/{vm_id}/label | Edit virtual machine labels



## DeleteInstance

> ResponseModel DeleteInstance(ctx, vmId).Execute()

Delete virtual machine
=======
[**AddFirewallRuleToVirtualMachine**](VirtualMachineAPI.md#AddFirewallRuleToVirtualMachine) | **Post** /core/virtual-machines/{id}/sg-rules | Add firewall rule to virtual machine
[**AttachFirewallsToAVirtualMachine**](VirtualMachineAPI.md#AttachFirewallsToAVirtualMachine) | **Post** /core/virtual-machines/{vm_id}/attach-firewalls | Attach firewalls to a virtual machine
[**CreateSnapshotFromAVirtualMachine**](VirtualMachineAPI.md#CreateSnapshotFromAVirtualMachine) | **Post** /core/virtual-machines/{vm_id}/snapshots | Create snapshot from a virtual machine
[**CreateVirtualMachines**](VirtualMachineAPI.md#CreateVirtualMachines) | **Post** /core/virtual-machines | Create virtual machines
[**DeleteFirewallRuleFromVirtualMachine**](VirtualMachineAPI.md#DeleteFirewallRuleFromVirtualMachine) | **Delete** /core/virtual-machines/{virtual_machine_id}/sg-rules/{sg_rule_id} | Delete firewall rule from virtual machine
[**DeleteVirtualMachine**](VirtualMachineAPI.md#DeleteVirtualMachine) | **Delete** /core/virtual-machines/{id} | Delete virtual machine
[**EditVirtualMachineLabels**](VirtualMachineAPI.md#EditVirtualMachineLabels) | **Put** /core/virtual-machines/{virtual_machine_id}/label | Edit virtual machine labels
[**FetchVirtualMachineNameAvailability**](VirtualMachineAPI.md#FetchVirtualMachineNameAvailability) | **Get** /core/virtual-machines/name-availability/{name} | Fetch virtual machine name availability
[**HardRebootVirtualMachine**](VirtualMachineAPI.md#HardRebootVirtualMachine) | **Get** /core/virtual-machines/{id}/hard-reboot | Hard reboot virtual machine
[**HibernateVirtualMachine**](VirtualMachineAPI.md#HibernateVirtualMachine) | **Get** /core/virtual-machines/{virtual_machine_id}/hibernate | Hibernate virtual machine
[**ListVirtualMachines**](VirtualMachineAPI.md#ListVirtualMachines) | **Get** /core/virtual-machines | List virtual machines
[**ResizeVirtualMachine**](VirtualMachineAPI.md#ResizeVirtualMachine) | **Post** /core/virtual-machines/{virtual_machine_id}/resize | Resize virtual machine
[**RestoreVirtualMachineFromHibernation**](VirtualMachineAPI.md#RestoreVirtualMachineFromHibernation) | **Get** /core/virtual-machines/{virtual_machine_id}/hibernate-restore | Restore virtual machine from hibernation
[**RetrieveVirtualMachineDetails**](VirtualMachineAPI.md#RetrieveVirtualMachineDetails) | **Get** /core/virtual-machines/{id} | Retrieve virtual machine details
[**RetrieveVirtualMachinePerformanceMetrics**](VirtualMachineAPI.md#RetrieveVirtualMachinePerformanceMetrics) | **Get** /core/virtual-machines/{virtual_machine_id}/metrics | Retrieve virtual machine performance metrics
[**RetrieveVirtualMachinesAssociatedWithAContract**](VirtualMachineAPI.md#RetrieveVirtualMachinesAssociatedWithAContract) | **Get** /core/virtual-machines/contract/{contract_id}/virtual-machines | Retrieve virtual machines associated with a contract
[**StartVirtualMachine**](VirtualMachineAPI.md#StartVirtualMachine) | **Get** /core/virtual-machines/{id}/start | Start virtual machine
[**StopVirtualMachine**](VirtualMachineAPI.md#StopVirtualMachine) | **Get** /core/virtual-machines/{id}/stop | Stop virtual machine



## AddFirewallRuleToVirtualMachine

> SecurityGroupRule AddFirewallRuleToVirtualMachine(ctx, id).Payload(payload).Execute()

Add firewall rule to virtual machine



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
	payload := *openapiclient.NewCreateSecurityRulePayload("Direction_example", "Ethertype_example", "any", "RemoteIpPrefix_example") // CreateSecurityRulePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.AddFirewallRuleToVirtualMachine(context.Background(), id).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.AddFirewallRuleToVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFirewallRuleToVirtualMachine`: SecurityGroupRule
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.AddFirewallRuleToVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFirewallRuleToVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**CreateSecurityRulePayload**](CreateSecurityRulePayload.md) |  | 

### Return type

[**SecurityGroupRule**](SecurityGroupRule.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachFirewallsToAVirtualMachine

> ResponseModel AttachFirewallsToAVirtualMachine(ctx, vmId).Payload(payload).Execute()

Attach firewalls to a virtual machine
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
	vmId := int32(56) // int32 | 
<<<<<<< HEAD

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.DeleteInstance(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.DeleteInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInstance`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.DeleteInstance`: %v\n", resp)
=======
	payload := *openapiclient.NewAttachFirewallsToVMPayload([]int32{int32(123)}) // AttachFirewallsToVMPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.AttachFirewallsToAVirtualMachine(context.Background(), vmId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.AttachFirewallsToAVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachFirewallsToAVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.AttachFirewallsToAVirtualMachine`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiAttachFirewallsToAVirtualMachineRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

<<<<<<< HEAD
=======
 **payload** | [**AttachFirewallsToVMPayload**](AttachFirewallsToVMPayload.md) |  | 
>>>>>>> main

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

<<<<<<< HEAD
[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
=======
[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
>>>>>>> main
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


<<<<<<< HEAD
## DeleteSecurityRule

> ResponseModel DeleteSecurityRule(ctx, vmId, sgRuleId).Execute()
=======
## CreateSnapshotFromAVirtualMachine

> CreateSnapshotResponse CreateSnapshotFromAVirtualMachine(ctx, vmId).Payload(payload).Execute()

Create snapshot from a virtual machine



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
	payload := *openapiclient.NewCreateSnapshotPayload("Description_example", false, "Name_example") // CreateSnapshotPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.CreateSnapshotFromAVirtualMachine(context.Background(), vmId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.CreateSnapshotFromAVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSnapshotFromAVirtualMachine`: CreateSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.CreateSnapshotFromAVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotFromAVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**CreateSnapshotPayload**](CreateSnapshotPayload.md) |  | 

### Return type

[**CreateSnapshotResponse**](CreateSnapshotResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVirtualMachines

> CreateInstancesResponse CreateVirtualMachines(ctx).Payload(payload).Execute()

Create virtual machines



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
	payload := *openapiclient.NewCreateInstancesPayload(int32(123), "EnvironmentName_example", "FlavorName_example", "KeyName_example", "Name_example") // CreateInstancesPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.CreateVirtualMachines(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.CreateVirtualMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVirtualMachines`: CreateInstancesResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.CreateVirtualMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateInstancesPayload**](CreateInstancesPayload.md) |  | 

### Return type

[**CreateInstancesResponse**](CreateInstancesResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewallRuleFromVirtualMachine

> ResponseModel DeleteFirewallRuleFromVirtualMachine(ctx, virtualMachineId, sgRuleId).Execute()
>>>>>>> main

Delete firewall rule from virtual machine



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
	sgRuleId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.VirtualMachineAPI.DeleteSecurityRule(context.Background(), vmId, sgRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.DeleteSecurityRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSecurityRule`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.DeleteSecurityRule`: %v\n", resp)
=======
	resp, r, err := apiClient.VirtualMachineAPI.DeleteFirewallRuleFromVirtualMachine(context.Background(), virtualMachineId, sgRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.DeleteFirewallRuleFromVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFirewallRuleFromVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.DeleteFirewallRuleFromVirtualMachine`: %v\n", resp)
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
**sgRuleId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiDeleteSecurityRuleRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiDeleteFirewallRuleFromVirtualMachineRequest struct via the builder pattern
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
=======
## DeleteVirtualMachine

> ResponseModel DeleteVirtualMachine(ctx, id).Execute()

Delete virtual machine



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
	resp, r, err := apiClient.VirtualMachineAPI.DeleteVirtualMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.DeleteVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.DeleteVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditVirtualMachineLabels

> ResponseModel EditVirtualMachineLabels(ctx, virtualMachineId).Payload(payload).Execute()

Edit virtual machine labels



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
	payload := *openapiclient.NewEditlabelofanexistingVMPayload() // EditlabelofanexistingVMPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.EditVirtualMachineLabels(context.Background(), virtualMachineId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.EditVirtualMachineLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditVirtualMachineLabels`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.EditVirtualMachineLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditVirtualMachineLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**EditlabelofanexistingVMPayload**](EditlabelofanexistingVMPayload.md) |  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


>>>>>>> main
## FetchVirtualMachineNameAvailability

> NameAvailableModel FetchVirtualMachineNameAvailability(ctx, name).Execute()

Fetch virtual machine name availability



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.FetchVirtualMachineNameAvailability(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.FetchVirtualMachineNameAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchVirtualMachineNameAvailability`: NameAvailableModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.FetchVirtualMachineNameAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchVirtualMachineNameAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NameAvailableModel**](NameAvailableModel.md)

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
## GetContractInstances

> ContractInstancesResponse GetContractInstances(ctx, contractId).Page(page).PageSize(pageSize).Search(search).Execute()
=======
## HardRebootVirtualMachine

> ResponseModel HardRebootVirtualMachine(ctx, id).Execute()

Hard reboot virtual machine



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
	resp, r, err := apiClient.VirtualMachineAPI.HardRebootVirtualMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.HardRebootVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HardRebootVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.HardRebootVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHardRebootVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HibernateVirtualMachine

> ResponseModel HibernateVirtualMachine(ctx, virtualMachineId).Execute()

Hibernate virtual machine



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.HibernateVirtualMachine(context.Background(), virtualMachineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.HibernateVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HibernateVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.HibernateVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHibernateVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualMachines

> Instances ListVirtualMachines(ctx).Page(page).PageSize(pageSize).Search(search).Environment(environment).Execute()

List virtual machines



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
	page := "page_example" // string | Page Number (optional)
	pageSize := "pageSize_example" // string | Data Per Page (optional)
	search := "search_example" // string |  (optional)
	environment := "environment_example" // string | Filter Environment ID or Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.ListVirtualMachines(context.Background()).Page(page).PageSize(pageSize).Search(search).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.ListVirtualMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVirtualMachines`: Instances
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.ListVirtualMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page Number | 
 **pageSize** | **string** | Data Per Page | 
 **search** | **string** |  | 
 **environment** | **string** | Filter Environment ID or Name | 

### Return type

[**Instances**](Instances.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResizeVirtualMachine

> ResponseModel ResizeVirtualMachine(ctx, virtualMachineId).Payload(payload).Execute()

Resize virtual machine



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
	payload := *openapiclient.NewInstanceResizePayload() // InstanceResizePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.ResizeVirtualMachine(context.Background(), virtualMachineId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.ResizeVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResizeVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.ResizeVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**InstanceResizePayload**](InstanceResizePayload.md) |  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreVirtualMachineFromHibernation

> ResponseModel RestoreVirtualMachineFromHibernation(ctx, virtualMachineId).Execute()

Restore virtual machine from hibernation



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.RestoreVirtualMachineFromHibernation(context.Background(), virtualMachineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.RestoreVirtualMachineFromHibernation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreVirtualMachineFromHibernation`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.RestoreVirtualMachineFromHibernation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreVirtualMachineFromHibernationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVirtualMachineDetails

> Instance RetrieveVirtualMachineDetails(ctx, id).Execute()

Retrieve virtual machine details



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
	resp, r, err := apiClient.VirtualMachineAPI.RetrieveVirtualMachineDetails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.RetrieveVirtualMachineDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVirtualMachineDetails`: Instance
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.RetrieveVirtualMachineDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVirtualMachineDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Instance**](Instance.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVirtualMachinePerformanceMetrics

> MetricsFields RetrieveVirtualMachinePerformanceMetrics(ctx, virtualMachineId).Duration(duration).Execute()

Retrieve virtual machine performance metrics



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
	duration := "duration_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.RetrieveVirtualMachinePerformanceMetrics(context.Background(), virtualMachineId).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.RetrieveVirtualMachinePerformanceMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVirtualMachinePerformanceMetrics`: MetricsFields
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.RetrieveVirtualMachinePerformanceMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVirtualMachinePerformanceMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duration** | **string** |  | 

### Return type

[**MetricsFields**](MetricsFields.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVirtualMachinesAssociatedWithAContract

> ContractInstancesResponse RetrieveVirtualMachinesAssociatedWithAContract(ctx, contractId).Page(page).PageSize(pageSize).Search(search).Execute()
>>>>>>> main

Retrieve virtual machines associated with a contract



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
	contractId := int32(56) // int32 | 
	page := "page_example" // string | Page Number (optional)
	pageSize := "pageSize_example" // string | Data Per Page (optional)
	search := "search_example" // string | Search By Instance ID or Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.VirtualMachineAPI.GetContractInstances(context.Background(), contractId).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.GetContractInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractInstances`: ContractInstancesResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.GetContractInstances`: %v\n", resp)
=======
	resp, r, err := apiClient.VirtualMachineAPI.RetrieveVirtualMachinesAssociatedWithAContract(context.Background(), contractId).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.RetrieveVirtualMachinesAssociatedWithAContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVirtualMachinesAssociatedWithAContract`: ContractInstancesResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.RetrieveVirtualMachinesAssociatedWithAContract`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGetContractInstancesRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveVirtualMachinesAssociatedWithAContractRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Page Number | 
 **pageSize** | **string** | Data Per Page | 
 **search** | **string** | Search By Instance ID or Name | 

### Return type

[**ContractInstancesResponse**](ContractInstancesResponse.md)

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
## GetInstance

> Instances GetInstance(ctx).Page(page).PageSize(pageSize).Search(search).Environment(environment).ExcludeFirewalls(excludeFirewalls).Execute()

List virtual machines



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
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	environment := "environment_example" // string |  (optional)
	excludeFirewalls := []int32{int32(123)} // []int32 | Comma-separated list of Security Group IDs to ignore instances attached (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.GetInstance(context.Background()).Page(page).PageSize(pageSize).Search(search).Environment(environment).ExcludeFirewalls(excludeFirewalls).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance`: Instances
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.GetInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **search** | **string** |  | 
 **environment** | **string** |  | 
 **excludeFirewalls** | **[]int32** | Comma-separated list of Security Group IDs to ignore instances attached | 

### Return type

[**Instances**](Instances.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance2

> Instance GetInstance2(ctx, vmId).Execute()

Retrieve virtual machine details



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
	resp, r, err := apiClient.VirtualMachineAPI.GetInstance2(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.GetInstance2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance2`: Instance
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.GetInstance2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstance2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Instance**](Instance.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance3

> ResponseModel GetInstance3(ctx, vmId).Execute()

Hard reboot virtual machine



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
	resp, r, err := apiClient.VirtualMachineAPI.GetInstance3(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.GetInstance3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance3`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.GetInstance3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstance3Request struct via the builder pattern


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


## GetInstance4

> ResponseModel GetInstance4(ctx, vmId).Execute()
=======
## StartVirtualMachine

> ResponseModel StartVirtualMachine(ctx, id).Execute()
>>>>>>> main

Start virtual machine



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
	resp, r, err := apiClient.VirtualMachineAPI.GetInstance4(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.GetInstance4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance4`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.GetInstance4`: %v\n", resp)
=======
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.StartVirtualMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.StartVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.StartVirtualMachine`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetInstance4Request struct via the builder pattern
=======
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartVirtualMachineRequest struct via the builder pattern
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
## GetInstance5

> ResponseModel GetInstance5(ctx, vmId).Execute()
=======
## StopVirtualMachine

> ResponseModel StopVirtualMachine(ctx, id).Execute()
>>>>>>> main

Stop virtual machine



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
	resp, r, err := apiClient.VirtualMachineAPI.GetInstance5(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.GetInstance5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance5`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.GetInstance5`: %v\n", resp)
=======
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.StopVirtualMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.StopVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.StopVirtualMachine`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetInstance5Request struct via the builder pattern
=======
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopVirtualMachineRequest struct via the builder pattern
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

## GetInstanceHibernate

> ResponseModel GetInstanceHibernate(ctx, vmId).Execute()

Hibernate virtual machine



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
	resp, r, err := apiClient.VirtualMachineAPI.GetInstanceHibernate(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.GetInstanceHibernate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceHibernate`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.GetInstanceHibernate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceHibernateRequest struct via the builder pattern


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


## GetInstanceHibernateRestore

> ResponseModel GetInstanceHibernateRestore(ctx, vmId).Execute()

Restore virtual machine from hibernation



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
	resp, r, err := apiClient.VirtualMachineAPI.GetInstanceHibernateRestore(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.GetInstanceHibernateRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceHibernateRestore`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.GetInstanceHibernateRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceHibernateRestoreRequest struct via the builder pattern


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


## GetInstanceLogs

> GetInstanceLogsResponse GetInstanceLogs(ctx, vmId).RequestId(requestId).Execute()

Get virtual machine logs



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
	requestId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.GetInstanceLogs(context.Background(), vmId).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.GetInstanceLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceLogs`: GetInstanceLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.GetInstanceLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestId** | **int32** |  | 

### Return type

[**GetInstanceLogsResponse**](GetInstanceLogsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceMetrics

> MetricsFields GetInstanceMetrics(ctx, vmId).Duration(duration).Execute()

Retrieve virtual machine performance metrics



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
	duration := "duration_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.GetInstanceMetrics(context.Background(), vmId).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.GetInstanceMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceMetrics`: MetricsFields
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.GetInstanceMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duration** | **string** |  | 

### Return type

[**MetricsFields**](MetricsFields.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInstance

> CreateInstancesResponse PostInstance(ctx).Payload(payload).Execute()

Create virtual machines



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
	payload := *openapiclient.NewCreateInstancesPayload(int32(123), "EnvironmentName_example", "FlavorName_example", "KeyName_example", "Name_example") // CreateInstancesPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.PostInstance(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.PostInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInstance`: CreateInstancesResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.PostInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateInstancesPayload**](CreateInstancesPayload.md) |  | 

### Return type

[**CreateInstancesResponse**](CreateInstancesResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInstanceAttachFirewalls

> ResponseModel PostInstanceAttachFirewalls(ctx, vmId).Payload(payload).Execute()

Attach firewalls to a virtual machine



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
	payload := *openapiclient.NewAttachFirewallsToVMPayload([]int32{int32(123)}) // AttachFirewallsToVMPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.PostInstanceAttachFirewalls(context.Background(), vmId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.PostInstanceAttachFirewalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInstanceAttachFirewalls`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.PostInstanceAttachFirewalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostInstanceAttachFirewallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AttachFirewallsToVMPayload**](AttachFirewallsToVMPayload.md) |  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInstanceLogs

> RequestInstanceLogsResponse PostInstanceLogs(ctx, vmId).Payload(payload).Execute()

Request virtual machine logs



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
	payload := *openapiclient.NewRequestInstanceLogsPayload() // RequestInstanceLogsPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.PostInstanceLogs(context.Background(), vmId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.PostInstanceLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInstanceLogs`: RequestInstanceLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.PostInstanceLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostInstanceLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**RequestInstanceLogsPayload**](RequestInstanceLogsPayload.md) |  | 

### Return type

[**RequestInstanceLogsResponse**](RequestInstanceLogsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInstanceResize

> ResponseModel PostInstanceResize(ctx, vmId).Payload(payload).Execute()

Resize virtual machine



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
	payload := *openapiclient.NewInstanceResizePayload() // InstanceResizePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.PostInstanceResize(context.Background(), vmId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.PostInstanceResize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInstanceResize`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.PostInstanceResize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostInstanceResizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**InstanceResizePayload**](InstanceResizePayload.md) |  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSecurityRule

> SecurityGroupRule PostSecurityRule(ctx, vmId).Payload(payload).Execute()

Add firewall rule to virtual machine



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
	payload := *openapiclient.NewCreateSecurityRulePayload("Direction_example", "Ethertype_example", "any", "RemoteIpPrefix_example") // CreateSecurityRulePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.PostSecurityRule(context.Background(), vmId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.PostSecurityRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSecurityRule`: SecurityGroupRule
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.PostSecurityRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSecurityRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**CreateSecurityRulePayload**](CreateSecurityRulePayload.md) |  | 

### Return type

[**SecurityGroupRule**](SecurityGroupRule.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSnapshots

> CreateSnapshotResponse PostSnapshots(ctx, vmId).Payload(payload).Execute()

Create snapshot from a virtual machine



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
	payload := *openapiclient.NewCreateSnapshotPayload("Description_example", "Name_example") // CreateSnapshotPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.PostSnapshots(context.Background(), vmId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.PostSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSnapshots`: CreateSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.PostSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**CreateSnapshotPayload**](CreateSnapshotPayload.md) |  | 

### Return type

[**CreateSnapshotResponse**](CreateSnapshotResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLabels

> ResponseModel PutLabels(ctx, vmId).Payload(payload).Execute()

Edit virtual machine labels



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
	payload := *openapiclient.NewEditLabelOfAnExistingVMPayload() // EditLabelOfAnExistingVMPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.PutLabels(context.Background(), vmId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.PutLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLabels`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.PutLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**EditLabelOfAnExistingVMPayload**](EditLabelOfAnExistingVMPayload.md) |  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

=======
>>>>>>> main
