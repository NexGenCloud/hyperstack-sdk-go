# \BillingAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllThresholdsForOrganization**](BillingAPI.md#GetAllThresholdsForOrganization) | **Get** /billing/billing/threshold | GET: All Thresholds for Organization
[**GetBillingUsage**](BillingAPI.md#GetBillingUsage) | **Get** /billing/billing/usage | GET: Billing usage
[**GetLastDayCost**](BillingAPI.md#GetLastDayCost) | **Get** /billing/billing/last-day-cost | GET: Last Day Cost
[**RetrieveBillingHistoryForASpecificBillingCycle**](BillingAPI.md#RetrieveBillingHistoryForASpecificBillingCycle) | **Get** /billing/billing/history | Retrieve Billing History for a specific Billing Cycle
[**RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle**](BillingAPI.md#RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle) | **Get** /billing/billing/history/virtual-machine/{vm_id} | Retrieve Billing History of a Specific Virtual Machine for a specific Billing Cycle
[**RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle**](BillingAPI.md#RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle) | **Get** /billing/billing/history/snapshot/{snapshot_id} | Retrieve Billing History of a Specific Volume for a specific Billing Cycle
[**RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle_0**](BillingAPI.md#RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle_0) | **Get** /billing/billing/history/volume/{volume_id} | Retrieve Billing History of a Specific Volume for a specific Billing Cycle
[**RetrieveBillingHistoryOfContractForASpecificBillingCycle**](BillingAPI.md#RetrieveBillingHistoryOfContractForASpecificBillingCycle) | **Get** /billing/billing/history/contract | Retrieve Billing History of Contract for a specific Billing Cycle
[**RetrieveBillingHistoryOfSnapshotForASpecificBillingCycle**](BillingAPI.md#RetrieveBillingHistoryOfSnapshotForASpecificBillingCycle) | **Get** /billing/billing/history/snapshot | Retrieve Billing History of Snapshot for a specific Billing Cycle
[**RetrieveBillingHistoryOfVirtualMachineForASpecificBillingCycle**](BillingAPI.md#RetrieveBillingHistoryOfVirtualMachineForASpecificBillingCycle) | **Get** /billing/billing/history/virtual-machine | Retrieve Billing History of Virtual Machine for a specific Billing Cycle
[**RetrieveBillingHistoryOfVolumeForASpecificBillingCycle**](BillingAPI.md#RetrieveBillingHistoryOfVolumeForASpecificBillingCycle) | **Get** /billing/billing/history/volume | Retrieve Billing History of Volume for a specific Billing Cycle
[**RetrieveHourlyCostDatapointsOfASpecificSnapshotForASpecificBillingCycle**](BillingAPI.md#RetrieveHourlyCostDatapointsOfASpecificSnapshotForASpecificBillingCycle) | **Get** /billing/billing/history/snapshot/{snapshot_id}/graph | Retrieve hourly cost datapoints of a Specific Snapshot for a specific billing cycle
[**RetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycle**](BillingAPI.md#RetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycle) | **Get** /billing/billing/history/virtual-machine/{vm_id}/graph | Retrieve hourly cost datapoints of a Specific Virtual Machine for a specific billing cycle
[**RetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycle**](BillingAPI.md#RetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycle) | **Get** /billing/billing/history/volume/{volume_id}/graph | Retrieve hourly cost datapoints of a Specific Volume for a specific billing cycle
[**RetrieveSubResourcesHistoricalCostDatapointsOfAVirtual**](BillingAPI.md#RetrieveSubResourcesHistoricalCostDatapointsOfAVirtual) | **Get** /billing/billing/virtual-machine/{vm_id}/sub-resource/graph | Retrieve Sub-Resources Historical Cost datapoints of a Virtual
[**RetrieveTotalCostsAndNonDiscountCostsForSubResources**](BillingAPI.md#RetrieveTotalCostsAndNonDiscountCostsForSubResources) | **Get** /billing/billing/virtual-machine/{vm_id}/sub-resource | Retrieve Total Costs and Non Discount Costs for Sub Resources
[**RetrieveVmBillingEventsHistory**](BillingAPI.md#RetrieveVmBillingEventsHistory) | **Get** /billing/billing/virtual-machine/{vm_id}/billing-events | Retrieve VM billing events history
[**RetrieveVolumeBillingEventsHistory**](BillingAPI.md#RetrieveVolumeBillingEventsHistory) | **Get** /billing/billing/volume/{volume_id}/billing-events | Retrieve Volume billing events history
[**UpdateSubscribeOrUnsubscribeNotificationThreshold**](BillingAPI.md#UpdateSubscribeOrUnsubscribeNotificationThreshold) | **Put** /billing/billing/threshold/{threshold_id} | Update: Subscribe or Unsubscribe Notification Threshold



## GetAllThresholdsForOrganization

> Organizationthresholdsresponse GetAllThresholdsForOrganization(ctx).Execute()

GET: All Thresholds for Organization



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetAllThresholdsForOrganization(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetAllThresholdsForOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllThresholdsForOrganization`: Organizationthresholdsresponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetAllThresholdsForOrganization`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllThresholdsForOrganizationRequest struct via the builder pattern


### Return type

[**Organizationthresholdsresponse**](Organizationthresholdsresponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingUsage

> Billingmetricesresponse GetBillingUsage(ctx).Deleted(deleted).Environment(environment).Execute()

GET: Billing usage



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
	deleted := "deleted_example" // string | `true` will return inactive resources and `false` will return active resources. By defualt(`deleted=false`) (optional)
	environment := "environment_example" // string | Filter resources by environment ID or Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetBillingUsage(context.Background()).Deleted(deleted).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingUsage`: Billingmetricesresponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleted** | **string** | &#x60;true&#x60; will return inactive resources and &#x60;false&#x60; will return active resources. By defualt(&#x60;deleted&#x3D;false&#x60;) | 
 **environment** | **string** | Filter resources by environment ID or Name | 

### Return type

[**Billingmetricesresponse**](Billingmetricesresponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastDayCost

> Lastdaycostresponse GetLastDayCost(ctx).Execute()

GET: Last Day Cost



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetLastDayCost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetLastDayCost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLastDayCost`: Lastdaycostresponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetLastDayCost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastDayCostRequest struct via the builder pattern


### Return type

[**Lastdaycostresponse**](Lastdaycostresponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBillingHistoryForASpecificBillingCycle

> OrganizationLevelBillingHistoryResponseModel RetrieveBillingHistoryForASpecificBillingCycle(ctx).StartDate(startDate).EndDate(endDate).Graph(graph).Execute()

Retrieve Billing History for a specific Billing Cycle



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
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	graph := "graph_example" // string | Set this value to \"true\" for getting graph value (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RetrieveBillingHistoryForASpecificBillingCycle(context.Background()).StartDate(startDate).EndDate(endDate).Graph(graph).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveBillingHistoryForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBillingHistoryForASpecificBillingCycle`: OrganizationLevelBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveBillingHistoryForASpecificBillingCycle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBillingHistoryForASpecificBillingCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **graph** | **string** | Set this value to \&quot;true\&quot; for getting graph value | 

### Return type

[**OrganizationLevelBillingHistoryResponseModel**](OrganizationLevelBillingHistoryResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle

> ResourceLevelVMBillingDetailsResponseModel RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve Billing History of a Specific Virtual Machine for a specific Billing Cycle



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
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle`: ResourceLevelVMBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVMBillingDetailsResponseModel**](ResourceLevelVMBillingDetailsResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle

> ResourceLevelVolumeBillingDetailsResponseModel RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle(ctx, snapshotId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve Billing History of a Specific Volume for a specific Billing Cycle



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
	snapshotId := int32(56) // int32 | 
	startDate := "startDate_example" // string | Datetime should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Datetime should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle(context.Background(), snapshotId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle`: ResourceLevelVolumeBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Datetime should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Datetime should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVolumeBillingDetailsResponseModel**](ResourceLevelVolumeBillingDetailsResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle_0

> ResourceLevelVolumeBillingDetailsResponseModel RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle_0(ctx, volumeId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve Billing History of a Specific Volume for a specific Billing Cycle



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
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle_0(context.Background(), volumeId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle_0`: ResourceLevelVolumeBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVolumeBillingDetailsResponseModel**](ResourceLevelVolumeBillingDetailsResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBillingHistoryOfContractForASpecificBillingCycle

> RetrieveBillingHistoryOfContractForASpecificBillingCycle(ctx).StartDate(startDate).EndDate(endDate).Search(search).Execute()

Retrieve Billing History of Contract for a specific Billing Cycle



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
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	search := "search_example" // string | Search by Contract \"Description\" or \"ID\" (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BillingAPI.RetrieveBillingHistoryOfContractForASpecificBillingCycle(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveBillingHistoryOfContractForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBillingHistoryOfContractForASpecificBillingCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by Contract \&quot;Description\&quot; or \&quot;ID\&quot; | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBillingHistoryOfSnapshotForASpecificBillingCycle

> ResourceLevelVolumeBillingHistoryResponseModel RetrieveBillingHistoryOfSnapshotForASpecificBillingCycle(ctx).StartDate(startDate).EndDate(endDate).Search(search).Page(page).PerPage(perPage).Execute()

Retrieve Billing History of Snapshot for a specific Billing Cycle



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
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	search := "search_example" // string | Search by Volume \"Name\" or \"ID\" (optional)
	page := "page_example" // string | Page number (optional)
	perPage := "perPage_example" // string | Number of items to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RetrieveBillingHistoryOfSnapshotForASpecificBillingCycle(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveBillingHistoryOfSnapshotForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBillingHistoryOfSnapshotForASpecificBillingCycle`: ResourceLevelVolumeBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveBillingHistoryOfSnapshotForASpecificBillingCycle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBillingHistoryOfSnapshotForASpecificBillingCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by Volume \&quot;Name\&quot; or \&quot;ID\&quot; | 
 **page** | **string** | Page number | 
 **perPage** | **string** | Number of items to return per page | 

### Return type

[**ResourceLevelVolumeBillingHistoryResponseModel**](ResourceLevelVolumeBillingHistoryResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBillingHistoryOfVirtualMachineForASpecificBillingCycle

> ResourceLevelVmBillingHistoryResponseModel RetrieveBillingHistoryOfVirtualMachineForASpecificBillingCycle(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

Retrieve Billing History of Virtual Machine for a specific Billing Cycle



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
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	search := "search_example" // string | Search by Virtual Machine \"Name\" or \"ID\" (optional)
	perPage := "perPage_example" // string | Number of items to return per page (optional)
	page := "page_example" // string | Page number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RetrieveBillingHistoryOfVirtualMachineForASpecificBillingCycle(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveBillingHistoryOfVirtualMachineForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBillingHistoryOfVirtualMachineForASpecificBillingCycle`: ResourceLevelVmBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveBillingHistoryOfVirtualMachineForASpecificBillingCycle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBillingHistoryOfVirtualMachineForASpecificBillingCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by Virtual Machine \&quot;Name\&quot; or \&quot;ID\&quot; | 
 **perPage** | **string** | Number of items to return per page | 
 **page** | **string** | Page number | 

### Return type

[**ResourceLevelVmBillingHistoryResponseModel**](ResourceLevelVmBillingHistoryResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBillingHistoryOfVolumeForASpecificBillingCycle

> ResourceLevelVolumeBillingHistoryResponseModel RetrieveBillingHistoryOfVolumeForASpecificBillingCycle(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

Retrieve Billing History of Volume for a specific Billing Cycle



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
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	search := "search_example" // string | Search by Volume \"Name\" or \"ID\" (optional)
	perPage := "perPage_example" // string | Number of items to return per page (optional)
	page := "page_example" // string | Page number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RetrieveBillingHistoryOfVolumeForASpecificBillingCycle(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveBillingHistoryOfVolumeForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBillingHistoryOfVolumeForASpecificBillingCycle`: ResourceLevelVolumeBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveBillingHistoryOfVolumeForASpecificBillingCycle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBillingHistoryOfVolumeForASpecificBillingCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by Volume \&quot;Name\&quot; or \&quot;ID\&quot; | 
 **perPage** | **string** | Number of items to return per page | 
 **page** | **string** | Page number | 

### Return type

[**ResourceLevelVolumeBillingHistoryResponseModel**](ResourceLevelVolumeBillingHistoryResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveHourlyCostDatapointsOfASpecificSnapshotForASpecificBillingCycle

> ResourceLevelVolumeGraphBillingDetailsResponseModel RetrieveHourlyCostDatapointsOfASpecificSnapshotForASpecificBillingCycle(ctx, snapshotId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve hourly cost datapoints of a Specific Snapshot for a specific billing cycle



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
	snapshotId := int32(56) // int32 | 
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RetrieveHourlyCostDatapointsOfASpecificSnapshotForASpecificBillingCycle(context.Background(), snapshotId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveHourlyCostDatapointsOfASpecificSnapshotForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveHourlyCostDatapointsOfASpecificSnapshotForASpecificBillingCycle`: ResourceLevelVolumeGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveHourlyCostDatapointsOfASpecificSnapshotForASpecificBillingCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveHourlyCostDatapointsOfASpecificSnapshotForASpecificBillingCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVolumeGraphBillingDetailsResponseModel**](ResourceLevelVolumeGraphBillingDetailsResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycle

> ResourceLevelVmGraphBillingDetailsResponseModel RetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycle(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve hourly cost datapoints of a Specific Virtual Machine for a specific billing cycle



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
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycle(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycle`: ResourceLevelVmGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVmGraphBillingDetailsResponseModel**](ResourceLevelVmGraphBillingDetailsResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycle

> ResourceLevelVolumeGraphBillingDetailsResponseModel RetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycle(ctx, volumeId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve hourly cost datapoints of a Specific Volume for a specific billing cycle



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
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycle(context.Background(), volumeId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycle`: ResourceLevelVolumeGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVolumeGraphBillingDetailsResponseModel**](ResourceLevelVolumeGraphBillingDetailsResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSubResourcesHistoricalCostDatapointsOfAVirtual

> SubResourcesGraphResponseModel RetrieveSubResourcesHistoricalCostDatapointsOfAVirtual(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve Sub-Resources Historical Cost datapoints of a Virtual



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
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RetrieveSubResourcesHistoricalCostDatapointsOfAVirtual(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveSubResourcesHistoricalCostDatapointsOfAVirtual``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveSubResourcesHistoricalCostDatapointsOfAVirtual`: SubResourcesGraphResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveSubResourcesHistoricalCostDatapointsOfAVirtual`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSubResourcesHistoricalCostDatapointsOfAVirtualRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**SubResourcesGraphResponseModel**](SubResourcesGraphResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveTotalCostsAndNonDiscountCostsForSubResources

> SubResourcesCostsResponseModel RetrieveTotalCostsAndNonDiscountCostsForSubResources(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve Total Costs and Non Discount Costs for Sub Resources



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
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RetrieveTotalCostsAndNonDiscountCostsForSubResources(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveTotalCostsAndNonDiscountCostsForSubResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveTotalCostsAndNonDiscountCostsForSubResources`: SubResourcesCostsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveTotalCostsAndNonDiscountCostsForSubResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveTotalCostsAndNonDiscountCostsForSubResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**SubResourcesCostsResponseModel**](SubResourcesCostsResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVmBillingEventsHistory

> ResourceBillingEventsHistoryResponse RetrieveVmBillingEventsHistory(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve VM billing events history



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
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RetrieveVmBillingEventsHistory(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveVmBillingEventsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVmBillingEventsHistory`: ResourceBillingEventsHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveVmBillingEventsHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVmBillingEventsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceBillingEventsHistoryResponse**](ResourceBillingEventsHistoryResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVolumeBillingEventsHistory

> ResourceBillingEventsHistoryResponse RetrieveVolumeBillingEventsHistory(ctx, volumeId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve Volume billing events history



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
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RetrieveVolumeBillingEventsHistory(context.Background(), volumeId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveVolumeBillingEventsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVolumeBillingEventsHistory`: ResourceBillingEventsHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveVolumeBillingEventsHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVolumeBillingEventsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceBillingEventsHistoryResponse**](ResourceBillingEventsHistoryResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscribeOrUnsubscribeNotificationThreshold

> Organizationthresholdupdateresponse UpdateSubscribeOrUnsubscribeNotificationThreshold(ctx, thresholdId).Payload(payload).Execute()

Update: Subscribe or Unsubscribe Notification Threshold



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
	thresholdId := int32(56) // int32 | 
	payload := *openapiclient.NewSubscribeorunsubscribeupdatepayload(false) // Subscribeorunsubscribeupdatepayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.UpdateSubscribeOrUnsubscribeNotificationThreshold(context.Background(), thresholdId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.UpdateSubscribeOrUnsubscribeNotificationThreshold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubscribeOrUnsubscribeNotificationThreshold`: Organizationthresholdupdateresponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.UpdateSubscribeOrUnsubscribeNotificationThreshold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thresholdId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscribeOrUnsubscribeNotificationThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**Subscribeorunsubscribeupdatepayload**](Subscribeorunsubscribeupdatepayload.md) |  | 

### Return type

[**Organizationthresholdupdateresponse**](Organizationthresholdupdateresponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

