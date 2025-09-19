# \BillingAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**GetLastDayCost**](BillingAPI.md#GetLastDayCost) | **Get** /billing/billing/last-day-cost | GET: Last Day Cost
[**GetOrganizationThreshold**](BillingAPI.md#GetOrganizationThreshold) | **Get** /billing/billing/threshold | GET: All Thresholds for Organization
[**GetUsage2**](BillingAPI.md#GetUsage2) | **Get** /billing/billing/usage | GET: Billing usage
[**GetUserBillingBucketDetailsGraph**](BillingAPI.md#GetUserBillingBucketDetailsGraph) | **Get** /billing/billing/history/bucket/{bucket_id}/graph | Retrieve hourly cost datapoints of a Specific Bucket for a specific billing cycle
[**GetUserBillingClusterDetailsGraph**](BillingAPI.md#GetUserBillingClusterDetailsGraph) | **Get** /billing/billing/history/cluster/{cluster_id}/graph | Retrieve hourly cost datapoints of a specific Cluster for a specific billing cycle
[**GetUserBillingDataSynthesisDetailsGraph**](BillingAPI.md#GetUserBillingDataSynthesisDetailsGraph) | **Get** /billing/billing/history/data_synthesis/{resource_id}/graph | Retrieve hourly cost datapoints of a Specific Data Synthesis for a specific
[**GetUserBillingFineTuningDetailsGraph**](BillingAPI.md#GetUserBillingFineTuningDetailsGraph) | **Get** /billing/billing/history/fine_tuning/{resource_id}/graph | Retrieve hourly cost datapoints of a Specific Fine Tuning for a specific billing cycle
[**GetUserBillingHistory2**](BillingAPI.md#GetUserBillingHistory2) | **Get** /billing/billing/history | Retrieve Billing History for a specific Billing Cycle
[**GetUserBillingHistoryBucket2**](BillingAPI.md#GetUserBillingHistoryBucket2) | **Get** /billing/billing/history/bucket | Retrieve Billing History of Volume for a specific Billing Cycle
[**GetUserBillingHistoryBucketDetails**](BillingAPI.md#GetUserBillingHistoryBucketDetails) | **Get** /billing/billing/history/bucket/{bucket_id} | Retrieve Billing History of a Specific Snapshot for a specific Billing Cycle
[**GetUserBillingHistoryCluster**](BillingAPI.md#GetUserBillingHistoryCluster) | **Get** /billing/billing/history/cluster | Retrieve Billing History of Clusters for a specific Billing Cycle
[**GetUserBillingHistoryClusterDetails**](BillingAPI.md#GetUserBillingHistoryClusterDetails) | **Get** /billing/billing/history/cluster/{cluster_id} | Retrieve Billing History of a Specific Cluster for a specific Billing Cycle
[**GetUserBillingHistoryContract**](BillingAPI.md#GetUserBillingHistoryContract) | **Get** /billing/billing/history/contract | Retrieve Billing History of Contract for a specific Billing Cycle
[**GetUserBillingHistoryDataSynthesis**](BillingAPI.md#GetUserBillingHistoryDataSynthesis) | **Get** /billing/billing/history/data_synthesis | Retrieve Billing History of data synthesis for a specific Billing Cycle
[**GetUserBillingHistoryDataSynthesisDetails**](BillingAPI.md#GetUserBillingHistoryDataSynthesisDetails) | **Get** /billing/billing/history/data_synthesis/{resource_id} | 
[**GetUserBillingHistoryFineTuning**](BillingAPI.md#GetUserBillingHistoryFineTuning) | **Get** /billing/billing/history/fine_tuning | Retrieve Billing History of model evaluation for a specific Billing Cycle
[**GetUserBillingHistoryFineTuningDetails**](BillingAPI.md#GetUserBillingHistoryFineTuningDetails) | **Get** /billing/billing/history/fine_tuning/{resource_id} | Retrieve Billing History of a Specific Fine Tuning for a specific Billing Cycle
[**GetUserBillingHistoryModelEvaluation**](BillingAPI.md#GetUserBillingHistoryModelEvaluation) | **Get** /billing/billing/history/model_evaluation | Retrieve Billing History of model evaluation for a specific Billing Cycle
[**GetUserBillingHistoryModelEvaluationDetails**](BillingAPI.md#GetUserBillingHistoryModelEvaluationDetails) | **Get** /billing/billing/history/model_evaluation/{resource_id} | 
[**GetUserBillingHistoryServerlessInference**](BillingAPI.md#GetUserBillingHistoryServerlessInference) | **Get** /billing/billing/history/serverless_inference | Retrieve Billing History of serverless inference for a specific Billing Cycle
[**GetUserBillingHistoryServerlessInferenceDetails**](BillingAPI.md#GetUserBillingHistoryServerlessInferenceDetails) | **Get** /billing/billing/history/serverless_inference/{resource_id} | 
[**GetUserBillingHistorySnapshot**](BillingAPI.md#GetUserBillingHistorySnapshot) | **Get** /billing/billing/history/snapshot | Retrieve Billing History of Snapshot for a specific Billing Cycle
[**GetUserBillingHistorySnapshotDetails**](BillingAPI.md#GetUserBillingHistorySnapshotDetails) | **Get** /billing/billing/history/snapshot/{snapshot_id} | Retrieve Billing History of a Specific Snapshot for a specific Billing Cycle
[**GetUserBillingHistoryVm2**](BillingAPI.md#GetUserBillingHistoryVm2) | **Get** /billing/billing/history/virtual-machine | Retrieve Billing History of Virtual Machine for a specific Billing Cycle
[**GetUserBillingHistoryVmDetails2**](BillingAPI.md#GetUserBillingHistoryVmDetails2) | **Get** /billing/billing/history/virtual-machine/{vm_id} | Retrieve Billing History of a Specific Virtual Machine for a specific Billing Cycle
[**GetUserBillingHistoryVmSubResourceGraph2**](BillingAPI.md#GetUserBillingHistoryVmSubResourceGraph2) | **Get** /billing/billing/virtual-machine/{vm_id}/sub-resource/graph | Retrieve Sub-Resources Historical Cost datapoints of a Virtual
[**GetUserBillingHistoryVmTotalCosts**](BillingAPI.md#GetUserBillingHistoryVmTotalCosts) | **Get** /billing/billing/virtual-machine/{vm_id}/sub-resource | Retrieve Total Costs and Non Discount Costs for Sub Resources
[**GetUserBillingHistoryVolume2**](BillingAPI.md#GetUserBillingHistoryVolume2) | **Get** /billing/billing/history/volume | Retrieve Billing History of Volume for a specific Billing Cycle
[**GetUserBillingHistoryVolumeDetails2**](BillingAPI.md#GetUserBillingHistoryVolumeDetails2) | **Get** /billing/billing/history/volume/{volume_id} | Retrieve Billing History of a Specific Volume for a specific Billing Cycle
[**GetUserBillingModelEvaluationDetailsGraph**](BillingAPI.md#GetUserBillingModelEvaluationDetailsGraph) | **Get** /billing/billing/history/model_evaluation/{resource_id}/graph | Retrieve hourly cost datapoints of a Specific Model Evaluation for a specific
[**GetUserBillingServerlessInferenceDetailsGraph**](BillingAPI.md#GetUserBillingServerlessInferenceDetailsGraph) | **Get** /billing/billing/history/serverless_inference/{resource_id}/graph | Retrieve hourly cost datapoints of a Specific Serverless Inference for a specific
[**GetUserBillingSnapshotDetailsGraph**](BillingAPI.md#GetUserBillingSnapshotDetailsGraph) | **Get** /billing/billing/history/snapshot/{snapshot_id}/graph | Retrieve hourly cost datapoints of a Specific Snapshot for a specific billing cycle
[**GetUserBillingVmDetailsGraph2**](BillingAPI.md#GetUserBillingVmDetailsGraph2) | **Get** /billing/billing/history/virtual-machine/{vm_id}/graph | Retrieve hourly cost datapoints of a Specific Virtual Machine for a specific billing cycle
[**GetUserBillingVolumeDetailsGraph**](BillingAPI.md#GetUserBillingVolumeDetailsGraph) | **Get** /billing/billing/history/volume/{volume_id}/graph | Retrieve hourly cost datapoints of a Specific Volume for a specific billing cycle
[**GetUserVmBillingEvents**](BillingAPI.md#GetUserVmBillingEvents) | **Get** /billing/billing/virtual-machine/{vm_id}/billing-events | Retrieve VM billing events history
[**GetUserVolumeBillingEvents**](BillingAPI.md#GetUserVolumeBillingEvents) | **Get** /billing/billing/volume/{volume_id}/billing-events | Retrieve Volume billing events history
[**PutOrganizationThreshold**](BillingAPI.md#PutOrganizationThreshold) | **Put** /billing/billing/threshold/{threshold_id} | Update: Subscribe or Unsubscribe Notification Threshold



## GetLastDayCost

> LastDayCostResponse GetLastDayCost(ctx).Execute()
=======
[**GetAllThresholdsForOrganization**](BillingAPI.md#GetAllThresholdsForOrganization) | **Get** /billing/billing/threshold | GET: All Thresholds for Organization
[**GetBillingUsage**](BillingAPI.md#GetBillingUsage) | **Get** /billing/billing/usage | GET: Billing usage
[**GetLastDayCost**](BillingAPI.md#GetLastDayCost) | **Get** /billing/billing/last-day-cost | GET: Last Day Cost
[**RetrieveBillingHistoryForASpecificBillingCycle**](BillingAPI.md#RetrieveBillingHistoryForASpecificBillingCycle) | **Get** /billing/billing/history | Retrieve Billing History for a specific Billing Cycle
[**RetrieveBillingHistoryOfASpecificSnapshotForASpecificBillingCycle**](BillingAPI.md#RetrieveBillingHistoryOfASpecificSnapshotForASpecificBillingCycle) | **Get** /billing/billing/history/snapshot/{snapshot_id} | Retrieve Billing History of a Specific Snapshot for a specific Billing Cycle
[**RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle**](BillingAPI.md#RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle) | **Get** /billing/billing/history/virtual-machine/{vm_id} | Retrieve Billing History of a Specific Virtual Machine for a specific Billing Cycle
[**RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle**](BillingAPI.md#RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle) | **Get** /billing/billing/history/volume/{volume_id} | Retrieve Billing History of a Specific Volume for a specific Billing Cycle
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
>>>>>>> main

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
<<<<<<< HEAD
	// response from `GetLastDayCost`: LastDayCostResponse
=======
	// response from `GetLastDayCost`: Lastdaycostresponse
>>>>>>> main
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetLastDayCost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastDayCostRequest struct via the builder pattern


### Return type

<<<<<<< HEAD
[**LastDayCostResponse**](LastDayCostResponse.md)

### Authorization

[apiKey](../README.md#apiKey)
=======
[**Lastdaycostresponse**](Lastdaycostresponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)
>>>>>>> main

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


<<<<<<< HEAD
## GetOrganizationThreshold

> OrganizationThresholdsResponse GetOrganizationThreshold(ctx).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetOrganizationThreshold(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetOrganizationThreshold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationThreshold`: OrganizationThresholdsResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetOrganizationThreshold`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationThresholdRequest struct via the builder pattern


### Return type

[**OrganizationThresholdsResponse**](OrganizationThresholdsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsage2

> BillingMetricesResponse GetUsage2(ctx).Deleted(deleted).Environment(environment).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetUsage2(context.Background()).Deleted(deleted).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUsage2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsage2`: BillingMetricesResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUsage2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsage2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleted** | **string** | &#x60;true&#x60; will return inactive resources and &#x60;false&#x60; will return active resources. By defualt(&#x60;deleted&#x3D;false&#x60;) | 
 **environment** | **string** | Filter resources by environment ID or Name | 

### Return type

[**BillingMetricesResponse**](BillingMetricesResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingBucketDetailsGraph

> ResourceLevelGraphBillingDetailsBucket GetUserBillingBucketDetailsGraph(ctx, bucketId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve hourly cost datapoints of a Specific Bucket for a specific billing cycle



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
	bucketId := int32(56) // int32 | 
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingBucketDetailsGraph(context.Background(), bucketId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingBucketDetailsGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingBucketDetailsGraph`: ResourceLevelGraphBillingDetailsBucket
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingBucketDetailsGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingBucketDetailsGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelGraphBillingDetailsBucket**](ResourceLevelGraphBillingDetailsBucket.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingClusterDetailsGraph

> ResourceLevelClusterGraphBillingDetailsResponseModel GetUserBillingClusterDetailsGraph(ctx, clusterId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve hourly cost datapoints of a specific Cluster for a specific billing cycle



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
	clusterId := int32(56) // int32 | 
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingClusterDetailsGraph(context.Background(), clusterId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingClusterDetailsGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingClusterDetailsGraph`: ResourceLevelClusterGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingClusterDetailsGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingClusterDetailsGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelClusterGraphBillingDetailsResponseModel**](ResourceLevelClusterGraphBillingDetailsResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingDataSynthesisDetailsGraph

> DataSynthesisBillingHistoryDetailsResponseSchema GetUserBillingDataSynthesisDetailsGraph(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve hourly cost datapoints of a Specific Data Synthesis for a specific



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
	resourceId := int32(56) // int32 | 
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingDataSynthesisDetailsGraph(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingDataSynthesisDetailsGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingDataSynthesisDetailsGraph`: DataSynthesisBillingHistoryDetailsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingDataSynthesisDetailsGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingDataSynthesisDetailsGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**DataSynthesisBillingHistoryDetailsResponseSchema**](DataSynthesisBillingHistoryDetailsResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingFineTuningDetailsGraph

> ResourceLevelVolumeGraphBillingDetailsResponseModel GetUserBillingFineTuningDetailsGraph(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve hourly cost datapoints of a Specific Fine Tuning for a specific billing cycle



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
	resourceId := int32(56) // int32 | 
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingFineTuningDetailsGraph(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingFineTuningDetailsGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingFineTuningDetailsGraph`: ResourceLevelVolumeGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingFineTuningDetailsGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingFineTuningDetailsGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVolumeGraphBillingDetailsResponseModel**](ResourceLevelVolumeGraphBillingDetailsResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistory2

> OrganizationLevelBillingHistoryResponseModel GetUserBillingHistory2(ctx).StartDate(startDate).EndDate(endDate).Graph(graph).Execute()
=======
## RetrieveBillingHistoryForASpecificBillingCycle

> OrganizationLevelBillingHistoryResponseModel RetrieveBillingHistoryForASpecificBillingCycle(ctx).StartDate(startDate).EndDate(endDate).Graph(graph).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistory2(context.Background()).StartDate(startDate).EndDate(endDate).Graph(graph).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistory2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistory2`: OrganizationLevelBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistory2`: %v\n", resp)
=======
	resp, r, err := apiClient.BillingAPI.RetrieveBillingHistoryForASpecificBillingCycle(context.Background()).StartDate(startDate).EndDate(endDate).Graph(graph).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveBillingHistoryForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBillingHistoryForASpecificBillingCycle`: OrganizationLevelBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveBillingHistoryForASpecificBillingCycle`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters



### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGetUserBillingHistory2Request struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveBillingHistoryForASpecificBillingCycleRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **graph** | **string** | Set this value to \&quot;true\&quot; for getting graph value | 

### Return type

[**OrganizationLevelBillingHistoryResponseModel**](OrganizationLevelBillingHistoryResponseModel.md)

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
## GetUserBillingHistoryBucket2

> ResourceLevelBucketBillingHistoryResponseModel GetUserBillingHistoryBucket2(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

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
	search := "search_example" // string | Search by resource \"Name\" or \"ID\" (optional)
	perPage := int32(56) // int32 | Number of items to return per page (optional)
	page := int32(56) // int32 | Page number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryBucket2(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryBucket2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryBucket2`: ResourceLevelBucketBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryBucket2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryBucket2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by resource \&quot;Name\&quot; or \&quot;ID\&quot; | 
 **perPage** | **int32** | Number of items to return per page | 
 **page** | **int32** | Page number | 

### Return type

[**ResourceLevelBucketBillingHistoryResponseModel**](ResourceLevelBucketBillingHistoryResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryBucketDetails

> ResourceLevelBucketBillingDetailsResponseModel GetUserBillingHistoryBucketDetails(ctx, bucketId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve Billing History of a Specific Snapshot for a specific Billing Cycle



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
	bucketId := int32(56) // int32 | 
	startDate := "startDate_example" // string | Datetime should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Datetime should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryBucketDetails(context.Background(), bucketId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryBucketDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryBucketDetails`: ResourceLevelBucketBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryBucketDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryBucketDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Datetime should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Datetime should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelBucketBillingDetailsResponseModel**](ResourceLevelBucketBillingDetailsResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryCluster

> ResourceLevelClusterBillingHistoryResponseModel GetUserBillingHistoryCluster(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

Retrieve Billing History of Clusters for a specific Billing Cycle



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
	search := "search_example" // string | Search by resource \"Name\" or \"ID\" (optional)
	perPage := int32(56) // int32 | Number of items to return per page (optional)
	page := int32(56) // int32 | Page number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryCluster(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryCluster`: ResourceLevelClusterBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by resource \&quot;Name\&quot; or \&quot;ID\&quot; | 
 **perPage** | **int32** | Number of items to return per page | 
 **page** | **int32** | Page number | 

### Return type

[**ResourceLevelClusterBillingHistoryResponseModel**](ResourceLevelClusterBillingHistoryResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryClusterDetails

> ResourceLevelClusterBillingDetailsResponseModel GetUserBillingHistoryClusterDetails(ctx, clusterId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve Billing History of a Specific Cluster for a specific Billing Cycle



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
	clusterId := int32(56) // int32 | 
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryClusterDetails(context.Background(), clusterId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryClusterDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryClusterDetails`: ResourceLevelClusterBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryClusterDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryClusterDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelClusterBillingDetailsResponseModel**](ResourceLevelClusterBillingDetailsResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryContract

> GetUserBillingHistoryContract(ctx).StartDate(startDate).EndDate(endDate).Search(search).Execute()

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
	r, err := apiClient.BillingAPI.GetUserBillingHistoryContract(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by Contract \&quot;Description\&quot; or \&quot;ID\&quot; | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryDataSynthesis

> TokenBasedBillingHistoryResponse GetUserBillingHistoryDataSynthesis(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

Retrieve Billing History of data synthesis for a specific Billing Cycle



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
	search := "search_example" // string | Search by resource \"Name\" or \"ID\" (optional)
	perPage := int32(56) // int32 | Number of items to return per page (optional)
	page := int32(56) // int32 | Page number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryDataSynthesis(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryDataSynthesis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryDataSynthesis`: TokenBasedBillingHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryDataSynthesis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryDataSynthesisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by resource \&quot;Name\&quot; or \&quot;ID\&quot; | 
 **perPage** | **int32** | Number of items to return per page | 
 **page** | **int32** | Page number | 

### Return type

[**TokenBasedBillingHistoryResponse**](TokenBasedBillingHistoryResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryDataSynthesisDetails

> DataSynthesisBillingHistoryDetailsResponseSchema GetUserBillingHistoryDataSynthesisDetails(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()





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
	resourceId := int32(56) // int32 | 
	startDate := "startDate_example" // string | YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryDataSynthesisDetails(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryDataSynthesisDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryDataSynthesisDetails`: DataSynthesisBillingHistoryDetailsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryDataSynthesisDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryDataSynthesisDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | YYYY-MM-DDTHH:MM:SS | 

### Return type

[**DataSynthesisBillingHistoryDetailsResponseSchema**](DataSynthesisBillingHistoryDetailsResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryFineTuning

> WorkloadBillingHistoryResponse GetUserBillingHistoryFineTuning(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

Retrieve Billing History of model evaluation for a specific Billing Cycle



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
	search := "search_example" // string | Search by resource \"Name\" or \"ID\" (optional)
	perPage := int32(56) // int32 | Number of items to return per page (optional)
	page := int32(56) // int32 | Page number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryFineTuning(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryFineTuning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryFineTuning`: WorkloadBillingHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryFineTuning`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryFineTuningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by resource \&quot;Name\&quot; or \&quot;ID\&quot; | 
 **perPage** | **int32** | Number of items to return per page | 
 **page** | **int32** | Page number | 

### Return type

[**WorkloadBillingHistoryResponse**](WorkloadBillingHistoryResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryFineTuningDetails

> ResourceLevelVolumeBillingDetailsResponseModel GetUserBillingHistoryFineTuningDetails(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve Billing History of a Specific Fine Tuning for a specific Billing Cycle



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
	resourceId := int32(56) // int32 | 
	startDate := "startDate_example" // string | Datetime should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Datetime should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryFineTuningDetails(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryFineTuningDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryFineTuningDetails`: ResourceLevelVolumeBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryFineTuningDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryFineTuningDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Datetime should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Datetime should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVolumeBillingDetailsResponseModel**](ResourceLevelVolumeBillingDetailsResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryModelEvaluation

> TokenBasedBillingHistoryResponse GetUserBillingHistoryModelEvaluation(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

Retrieve Billing History of model evaluation for a specific Billing Cycle



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
	search := "search_example" // string | Search by resource \"Name\" or \"ID\" (optional)
	perPage := int32(56) // int32 | Number of items to return per page (optional)
	page := int32(56) // int32 | Page number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryModelEvaluation(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryModelEvaluation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryModelEvaluation`: TokenBasedBillingHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryModelEvaluation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryModelEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by resource \&quot;Name\&quot; or \&quot;ID\&quot; | 
 **perPage** | **int32** | Number of items to return per page | 
 **page** | **int32** | Page number | 

### Return type

[**TokenBasedBillingHistoryResponse**](TokenBasedBillingHistoryResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryModelEvaluationDetails

> ModelEvaluationBillingHistoryDetailsResponseSchema GetUserBillingHistoryModelEvaluationDetails(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()





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
	resourceId := int32(56) // int32 | 
	startDate := "startDate_example" // string | YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryModelEvaluationDetails(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryModelEvaluationDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryModelEvaluationDetails`: ModelEvaluationBillingHistoryDetailsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryModelEvaluationDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryModelEvaluationDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ModelEvaluationBillingHistoryDetailsResponseSchema**](ModelEvaluationBillingHistoryDetailsResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryServerlessInference

> TokenBasedBillingHistoryResponse GetUserBillingHistoryServerlessInference(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

Retrieve Billing History of serverless inference for a specific Billing Cycle



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
	search := "search_example" // string | Search by resource \"Name\" or \"ID\" (optional)
	perPage := int32(56) // int32 | Number of items to return per page (optional)
	page := int32(56) // int32 | Page number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryServerlessInference(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryServerlessInference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryServerlessInference`: TokenBasedBillingHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryServerlessInference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryServerlessInferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by resource \&quot;Name\&quot; or \&quot;ID\&quot; | 
 **perPage** | **int32** | Number of items to return per page | 
 **page** | **int32** | Page number | 

### Return type

[**TokenBasedBillingHistoryResponse**](TokenBasedBillingHistoryResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryServerlessInferenceDetails

> ServerlessInferencedBillingHistoryDetailsResponseSchema GetUserBillingHistoryServerlessInferenceDetails(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()





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
	resourceId := int32(56) // int32 | 
	startDate := "startDate_example" // string | YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryServerlessInferenceDetails(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryServerlessInferenceDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryServerlessInferenceDetails`: ServerlessInferencedBillingHistoryDetailsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryServerlessInferenceDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryServerlessInferenceDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ServerlessInferencedBillingHistoryDetailsResponseSchema**](ServerlessInferencedBillingHistoryDetailsResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistorySnapshot

> ResourceLevelVolumeBillingHistoryResponseModel GetUserBillingHistorySnapshot(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

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
	search := "search_example" // string | Search by resource \"Name\" or \"ID\" (optional)
	perPage := int32(56) // int32 | Number of items to return per page (optional)
	page := int32(56) // int32 | Page number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistorySnapshot(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistorySnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistorySnapshot`: ResourceLevelVolumeBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistorySnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistorySnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by resource \&quot;Name\&quot; or \&quot;ID\&quot; | 
 **perPage** | **int32** | Number of items to return per page | 
 **page** | **int32** | Page number | 

### Return type

[**ResourceLevelVolumeBillingHistoryResponseModel**](ResourceLevelVolumeBillingHistoryResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistorySnapshotDetails

> ResourceLevelVolumeBillingDetailsResponseModel GetUserBillingHistorySnapshotDetails(ctx, snapshotId).StartDate(startDate).EndDate(endDate).Execute()
=======
## RetrieveBillingHistoryOfASpecificSnapshotForASpecificBillingCycle

> ResourceLevelVolumeBillingDetailsResponseModel RetrieveBillingHistoryOfASpecificSnapshotForASpecificBillingCycle(ctx, snapshotId).StartDate(startDate).EndDate(endDate).Execute()
>>>>>>> main

Retrieve Billing History of a Specific Snapshot for a specific Billing Cycle



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
<<<<<<< HEAD
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistorySnapshotDetails(context.Background(), snapshotId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistorySnapshotDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistorySnapshotDetails`: ResourceLevelVolumeBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistorySnapshotDetails`: %v\n", resp)
=======
	resp, r, err := apiClient.BillingAPI.RetrieveBillingHistoryOfASpecificSnapshotForASpecificBillingCycle(context.Background(), snapshotId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveBillingHistoryOfASpecificSnapshotForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBillingHistoryOfASpecificSnapshotForASpecificBillingCycle`: ResourceLevelVolumeBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveBillingHistoryOfASpecificSnapshotForASpecificBillingCycle`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGetUserBillingHistorySnapshotDetailsRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveBillingHistoryOfASpecificSnapshotForASpecificBillingCycleRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Datetime should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Datetime should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVolumeBillingDetailsResponseModel**](ResourceLevelVolumeBillingDetailsResponseModel.md)

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
## GetUserBillingHistoryVm2

> ResourceLevelVmBillingHistoryResponseModel GetUserBillingHistoryVm2(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

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
	search := "search_example" // string | Search by resource \"Name\" or \"ID\" (optional)
	perPage := int32(56) // int32 | Number of items to return per page (optional)
	page := int32(56) // int32 | Page number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryVm2(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryVm2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryVm2`: ResourceLevelVmBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryVm2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryVm2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by resource \&quot;Name\&quot; or \&quot;ID\&quot; | 
 **perPage** | **int32** | Number of items to return per page | 
 **page** | **int32** | Page number | 

### Return type

[**ResourceLevelVmBillingHistoryResponseModel**](ResourceLevelVmBillingHistoryResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryVmDetails2

> ResourceLevelVMBillingDetailsResponseModel GetUserBillingHistoryVmDetails2(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()
=======
## RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle

> ResourceLevelVMBillingDetailsResponseModel RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryVmDetails2(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryVmDetails2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryVmDetails2`: ResourceLevelVMBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryVmDetails2`: %v\n", resp)
=======
	resp, r, err := apiClient.BillingAPI.RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle`: ResourceLevelVMBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycle`: %v\n", resp)
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
Other parameters are passed through a pointer to a apiGetUserBillingHistoryVmDetails2Request struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveBillingHistoryOfASpecificVirtualMachineForASpecificBillingCycleRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVMBillingDetailsResponseModel**](ResourceLevelVMBillingDetailsResponseModel.md)

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
## GetUserBillingHistoryVmSubResourceGraph2

> SubResourcesGraphResponseModel GetUserBillingHistoryVmSubResourceGraph2(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryVmSubResourceGraph2(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryVmSubResourceGraph2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryVmSubResourceGraph2`: SubResourcesGraphResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryVmSubResourceGraph2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryVmSubResourceGraph2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**SubResourcesGraphResponseModel**](SubResourcesGraphResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryVmTotalCosts

> SubResourcesCostsResponseModel GetUserBillingHistoryVmTotalCosts(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryVmTotalCosts(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryVmTotalCosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryVmTotalCosts`: SubResourcesCostsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryVmTotalCosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryVmTotalCostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**SubResourcesCostsResponseModel**](SubResourcesCostsResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryVolume2

> ResourceLevelVolumeBillingHistoryResponseModel GetUserBillingHistoryVolume2(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

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
	search := "search_example" // string | Search by resource \"Name\" or \"ID\" (optional)
	perPage := int32(56) // int32 | Number of items to return per page (optional)
	page := int32(56) // int32 | Page number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryVolume2(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryVolume2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryVolume2`: ResourceLevelVolumeBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryVolume2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryVolume2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by resource \&quot;Name\&quot; or \&quot;ID\&quot; | 
 **perPage** | **int32** | Number of items to return per page | 
 **page** | **int32** | Page number | 

### Return type

[**ResourceLevelVolumeBillingHistoryResponseModel**](ResourceLevelVolumeBillingHistoryResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingHistoryVolumeDetails2

> ResourceLevelVolumeBillingDetailsResponseModel GetUserBillingHistoryVolumeDetails2(ctx, volumeId).StartDate(startDate).EndDate(endDate).Execute()
=======
## RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle

> ResourceLevelVolumeBillingDetailsResponseModel RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle(ctx, volumeId).StartDate(startDate).EndDate(endDate).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistoryVolumeDetails2(context.Background(), volumeId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistoryVolumeDetails2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistoryVolumeDetails2`: ResourceLevelVolumeBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistoryVolumeDetails2`: %v\n", resp)
=======
	resp, r, err := apiClient.BillingAPI.RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle(context.Background(), volumeId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle`: ResourceLevelVolumeBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycle`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGetUserBillingHistoryVolumeDetails2Request struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveBillingHistoryOfASpecificVolumeForASpecificBillingCycleRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVolumeBillingDetailsResponseModel**](ResourceLevelVolumeBillingDetailsResponseModel.md)

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
## GetUserBillingModelEvaluationDetailsGraph

> ModelEvaluationBillingHistoryDetailsResponseSchema GetUserBillingModelEvaluationDetailsGraph(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve hourly cost datapoints of a Specific Model Evaluation for a specific
=======
## RetrieveBillingHistoryOfContractForASpecificBillingCycle

> RetrieveBillingHistoryOfContractForASpecificBillingCycle(ctx).StartDate(startDate).EndDate(endDate).Search(search).Execute()

Retrieve Billing History of Contract for a specific Billing Cycle
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
	resourceId := int32(56) // int32 | 
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingModelEvaluationDetailsGraph(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingModelEvaluationDetailsGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingModelEvaluationDetailsGraph`: ModelEvaluationBillingHistoryDetailsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingModelEvaluationDetailsGraph`: %v\n", resp)
=======
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
>>>>>>> main
}
```

### Path Parameters


<<<<<<< HEAD
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingModelEvaluationDetailsGraphRequest struct via the builder pattern
=======

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBillingHistoryOfContractForASpecificBillingCycleRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
<<<<<<< HEAD

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ModelEvaluationBillingHistoryDetailsResponseSchema**](ModelEvaluationBillingHistoryDetailsResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)
=======
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by Contract \&quot;Description\&quot; or \&quot;ID\&quot; | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)
>>>>>>> main

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


<<<<<<< HEAD
## GetUserBillingServerlessInferenceDetailsGraph

> ServerlessInferencedBillingHistoryDetailsResponseSchema GetUserBillingServerlessInferenceDetailsGraph(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve hourly cost datapoints of a Specific Serverless Inference for a specific
=======
## RetrieveBillingHistoryOfSnapshotForASpecificBillingCycle

> ResourceLevelVolumeBillingHistoryResponseModel RetrieveBillingHistoryOfSnapshotForASpecificBillingCycle(ctx).StartDate(startDate).EndDate(endDate).Search(search).Page(page).PerPage(perPage).Execute()

Retrieve Billing History of Snapshot for a specific Billing Cycle
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
	resourceId := int32(56) // int32 | 
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetUserBillingServerlessInferenceDetailsGraph(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingServerlessInferenceDetailsGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingServerlessInferenceDetailsGraph`: ServerlessInferencedBillingHistoryDetailsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingServerlessInferenceDetailsGraph`: %v\n", resp)
=======
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
>>>>>>> main
}
```

### Path Parameters


<<<<<<< HEAD
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingServerlessInferenceDetailsGraphRequest struct via the builder pattern
=======

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBillingHistoryOfSnapshotForASpecificBillingCycleRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
<<<<<<< HEAD

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ServerlessInferencedBillingHistoryDetailsResponseSchema**](ServerlessInferencedBillingHistoryDetailsResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)
=======
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **search** | **string** | Search by Volume \&quot;Name\&quot; or \&quot;ID\&quot; | 
 **page** | **string** | Page number | 
 **perPage** | **string** | Number of items to return per page | 

### Return type

[**ResourceLevelVolumeBillingHistoryResponseModel**](ResourceLevelVolumeBillingHistoryResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)
>>>>>>> main

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


<<<<<<< HEAD
## GetUserBillingSnapshotDetailsGraph

> ResourceLevelVolumeGraphBillingDetailsResponseModel GetUserBillingSnapshotDetailsGraph(ctx, snapshotId).StartDate(startDate).EndDate(endDate).Execute()
=======
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
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.BillingAPI.GetUserBillingSnapshotDetailsGraph(context.Background(), snapshotId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingSnapshotDetailsGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingSnapshotDetailsGraph`: ResourceLevelVolumeGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingSnapshotDetailsGraph`: %v\n", resp)
=======
	resp, r, err := apiClient.BillingAPI.RetrieveHourlyCostDatapointsOfASpecificSnapshotForASpecificBillingCycle(context.Background(), snapshotId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveHourlyCostDatapointsOfASpecificSnapshotForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveHourlyCostDatapointsOfASpecificSnapshotForASpecificBillingCycle`: ResourceLevelVolumeGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveHourlyCostDatapointsOfASpecificSnapshotForASpecificBillingCycle`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGetUserBillingSnapshotDetailsGraphRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveHourlyCostDatapointsOfASpecificSnapshotForASpecificBillingCycleRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVolumeGraphBillingDetailsResponseModel**](ResourceLevelVolumeGraphBillingDetailsResponseModel.md)

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
## GetUserBillingVmDetailsGraph2

> ResourceLevelVmGraphBillingDetailsResponseModel GetUserBillingVmDetailsGraph2(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()
=======
## RetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycle

> ResourceLevelVmGraphBillingDetailsResponseModel RetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycle(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.BillingAPI.GetUserBillingVmDetailsGraph2(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingVmDetailsGraph2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingVmDetailsGraph2`: ResourceLevelVmGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingVmDetailsGraph2`: %v\n", resp)
=======
	resp, r, err := apiClient.BillingAPI.RetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycle(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycle`: ResourceLevelVmGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycle`: %v\n", resp)
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
Other parameters are passed through a pointer to a apiGetUserBillingVmDetailsGraph2Request struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveHourlyCostDatapointsOfASpecificVirtualMachineForASpecificBillingCycleRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVmGraphBillingDetailsResponseModel**](ResourceLevelVmGraphBillingDetailsResponseModel.md)

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
## GetUserBillingVolumeDetailsGraph

> ResourceLevelVolumeGraphBillingDetailsResponseModel GetUserBillingVolumeDetailsGraph(ctx, volumeId).StartDate(startDate).EndDate(endDate).Execute()
=======
## RetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycle

> ResourceLevelVolumeGraphBillingDetailsResponseModel RetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycle(ctx, volumeId).StartDate(startDate).EndDate(endDate).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.BillingAPI.GetUserBillingVolumeDetailsGraph(context.Background(), volumeId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingVolumeDetailsGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingVolumeDetailsGraph`: ResourceLevelVolumeGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingVolumeDetailsGraph`: %v\n", resp)
=======
	resp, r, err := apiClient.BillingAPI.RetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycle(context.Background(), volumeId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycle`: ResourceLevelVolumeGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycle`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGetUserBillingVolumeDetailsGraphRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveHourlyCostDatapointsOfASpecificVolumeForASpecificBillingCycleRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVolumeGraphBillingDetailsResponseModel**](ResourceLevelVolumeGraphBillingDetailsResponseModel.md)

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
## GetUserVmBillingEvents

> ResourceBillingEventsHistoryResponse GetUserVmBillingEvents(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()
=======
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
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.BillingAPI.GetUserVmBillingEvents(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserVmBillingEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserVmBillingEvents`: ResourceBillingEventsHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserVmBillingEvents`: %v\n", resp)
=======
	resp, r, err := apiClient.BillingAPI.RetrieveVmBillingEventsHistory(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveVmBillingEventsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVmBillingEventsHistory`: ResourceBillingEventsHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveVmBillingEventsHistory`: %v\n", resp)
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
Other parameters are passed through a pointer to a apiGetUserVmBillingEventsRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveVmBillingEventsHistoryRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceBillingEventsHistoryResponse**](ResourceBillingEventsHistoryResponse.md)

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
## GetUserVolumeBillingEvents

> ResourceBillingEventsHistoryResponse GetUserVolumeBillingEvents(ctx, volumeId).StartDate(startDate).EndDate(endDate).Execute()
=======
## RetrieveVolumeBillingEventsHistory

> ResourceBillingEventsHistoryResponse RetrieveVolumeBillingEventsHistory(ctx, volumeId).StartDate(startDate).EndDate(endDate).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.BillingAPI.GetUserVolumeBillingEvents(context.Background(), volumeId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserVolumeBillingEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserVolumeBillingEvents`: ResourceBillingEventsHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserVolumeBillingEvents`: %v\n", resp)
=======
	resp, r, err := apiClient.BillingAPI.RetrieveVolumeBillingEventsHistory(context.Background(), volumeId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetrieveVolumeBillingEventsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVolumeBillingEventsHistory`: ResourceBillingEventsHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetrieveVolumeBillingEventsHistory`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGetUserVolumeBillingEventsRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveVolumeBillingEventsHistoryRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceBillingEventsHistoryResponse**](ResourceBillingEventsHistoryResponse.md)

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
## PutOrganizationThreshold

> OrganizationThresholdUpdateResponse PutOrganizationThreshold(ctx, thresholdId).Payload(payload).Execute()
=======
## UpdateSubscribeOrUnsubscribeNotificationThreshold

> Organizationthresholdupdateresponse UpdateSubscribeOrUnsubscribeNotificationThreshold(ctx, thresholdId).Payload(payload).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	payload := *openapiclient.NewSubscribeOrUnsubscribeUpdatePayload(false) // SubscribeOrUnsubscribeUpdatePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.PutOrganizationThreshold(context.Background(), thresholdId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PutOrganizationThreshold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrganizationThreshold`: OrganizationThresholdUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.PutOrganizationThreshold`: %v\n", resp)
=======
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
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thresholdId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiPutOrganizationThresholdRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiUpdateSubscribeOrUnsubscribeNotificationThresholdRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

<<<<<<< HEAD
 **payload** | [**SubscribeOrUnsubscribeUpdatePayload**](SubscribeOrUnsubscribeUpdatePayload.md) |  | 

### Return type

[**OrganizationThresholdUpdateResponse**](OrganizationThresholdUpdateResponse.md)

### Authorization

[apiKey](../README.md#apiKey)
=======
 **payload** | [**Subscribeorunsubscribeupdatepayload**](Subscribeorunsubscribeupdatepayload.md) |  | 

### Return type

[**Organizationthresholdupdateresponse**](Organizationthresholdupdateresponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)
>>>>>>> main

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

