# \BillingAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BucketsBillingHistoryHourlyChart**](BillingAPI.md#BucketsBillingHistoryHourlyChart) | **Get** /billing/billing/history/bucket/{bucket_id}/graph | Retrieve hourly cost datapoints of a Specific Bucket for a specific billing cycle
[**GetBucketBillingHistory**](BillingAPI.md#GetBucketBillingHistory) | **Get** /billing/billing/history/bucket/{bucket_id} | Retrieve Billing History of a Specific Snapshot for a specific Billing Cycle
[**GetClusterBillingHistory**](BillingAPI.md#GetClusterBillingHistory) | **Get** /billing/billing/history/cluster/{cluster_id} | Retrieve Billing History of a Specific Cluster for a specific Billing Cycle
[**GetClusterBillingHistoryGraph**](BillingAPI.md#GetClusterBillingHistoryGraph) | **Get** /billing/billing/history/cluster/{cluster_id}/graph | Retrieve hourly cost datapoints of a specific Cluster for a specific billing cycle
[**GetDataSynthesisBillingHistory**](BillingAPI.md#GetDataSynthesisBillingHistory) | **Get** /billing/billing/history/data_synthesis | Retrieve Billing History of data synthesis for a specific Billing Cycle
[**GetDataSynthesisBillingHistoryGraph**](BillingAPI.md#GetDataSynthesisBillingHistoryGraph) | **Get** /billing/billing/history/data_synthesis/{resource_id}/graph | Retrieve hourly cost datapoints of a Specific Data Synthesis for a specific
[**GetDataSynthesisHistoryForResource**](BillingAPI.md#GetDataSynthesisHistoryForResource) | **Get** /billing/billing/history/data_synthesis/{resource_id} | 
[**GetFineTuningBillingHistory**](BillingAPI.md#GetFineTuningBillingHistory) | **Get** /billing/billing/history/fine_tuning | Retrieve Billing History of model evaluation for a specific Billing Cycle
[**GetFineTuningBillingHistoryGraph**](BillingAPI.md#GetFineTuningBillingHistoryGraph) | **Get** /billing/billing/history/fine_tuning/{resource_id}/graph | Retrieve hourly cost datapoints of a Specific Fine Tuning for a specific billing cycle
[**GetLastDayCost**](BillingAPI.md#GetLastDayCost) | **Get** /billing/billing/last-day-cost | GET: Last Day Cost
[**GetModelEvaluationBillingHistory**](BillingAPI.md#GetModelEvaluationBillingHistory) | **Get** /billing/billing/history/model_evaluation | Retrieve Billing History of model evaluation for a specific Billing Cycle
[**GetModelEvaluationBillingHistoryGraph**](BillingAPI.md#GetModelEvaluationBillingHistoryGraph) | **Get** /billing/billing/history/model_evaluation/{resource_id}/graph | Retrieve hourly cost datapoints of a Specific Model Evaluation for a specific
[**GetNotificationThreshold**](BillingAPI.md#GetNotificationThreshold) | **Put** /billing/billing/threshold/{threshold_id} | Update: Subscribe or Unsubscribe Notification Threshold
[**GetResourceFineTuningBillingHistory**](BillingAPI.md#GetResourceFineTuningBillingHistory) | **Get** /billing/billing/history/fine_tuning/{resource_id} | Retrieve Billing History of a Specific Fine Tuning for a specific Billing Cycle
[**GetResourceModelEvaluationBillingHistory**](BillingAPI.md#GetResourceModelEvaluationBillingHistory) | **Get** /billing/billing/history/model_evaluation/{resource_id} | 
[**GetServerlessInferenceBillingHistoryGraph**](BillingAPI.md#GetServerlessInferenceBillingHistoryGraph) | **Get** /billing/billing/history/serverless_inference/{resource_id}/graph | Retrieve hourly cost datapoints of a Specific Serverless Inference for a specific
[**GetServerlessInferencesBillingHistory**](BillingAPI.md#GetServerlessInferencesBillingHistory) | **Get** /billing/billing/history/serverless_inference/{resource_id} | 
[**GetSnapshotBillingHistory**](BillingAPI.md#GetSnapshotBillingHistory) | **Get** /billing/billing/history/snapshot/{snapshot_id} | Retrieve Billing History of a Specific Snapshot for a specific Billing Cycle
[**GetSnapshotBillingHistoryGraph**](BillingAPI.md#GetSnapshotBillingHistoryGraph) | **Get** /billing/billing/history/snapshot/{snapshot_id}/graph | Retrieve hourly cost datapoints of a Specific Snapshot for a specific billing cycle
[**GetUsage**](BillingAPI.md#GetUsage) | **Get** /billing/billing/usage | GET: Billing usage
[**GetUserBillingHistory**](BillingAPI.md#GetUserBillingHistory) | **Get** /billing/billing/history | Retrieve Billing History for a specific Billing Cycle
[**GetVMBillingDetails**](BillingAPI.md#GetVMBillingDetails) | **Get** /billing/billing/history/virtual-machine/{vm_id} | Retrieve Billing History of a Specific Virtual Machine for a specific Billing Cycle
[**GetVMBillingEvents**](BillingAPI.md#GetVMBillingEvents) | **Get** /billing/billing/virtual-machine/{vm_id}/billing-events | Retrieve VM billing events history
[**GetVMBillingGraph**](BillingAPI.md#GetVMBillingGraph) | **Get** /billing/billing/history/virtual-machine/{vm_id}/graph | Retrieve hourly cost datapoints of a Specific Virtual Machine for a specific billing cycle
[**GetVMBillingHistory**](BillingAPI.md#GetVMBillingHistory) | **Get** /billing/billing/history/virtual-machine | Retrieve Billing History of Virtual Machine for a specific Billing Cycle
[**GetVMSubResourceCosts**](BillingAPI.md#GetVMSubResourceCosts) | **Get** /billing/billing/virtual-machine/{vm_id}/sub-resource | Retrieve Total Costs and Non Discount Costs for Sub Resources
[**GetVMSubResourceGraph**](BillingAPI.md#GetVMSubResourceGraph) | **Get** /billing/billing/virtual-machine/{vm_id}/sub-resource/graph | Retrieve Sub-Resources Historical Cost datapoints of a Virtual
[**GetVolumeBillingDetails**](BillingAPI.md#GetVolumeBillingDetails) | **Get** /billing/billing/history/volume/{volume_id} | Retrieve Billing History of a Specific Volume for a specific Billing Cycle
[**GetVolumeBillingEvents**](BillingAPI.md#GetVolumeBillingEvents) | **Get** /billing/billing/volume/{volume_id}/billing-events | Retrieve Volume billing events history
[**GetVolumeBillingHistory**](BillingAPI.md#GetVolumeBillingHistory) | **Get** /billing/billing/history/volume | Retrieve Billing History of Volume for a specific Billing Cycle
[**GetVolumeBillingHistoryGraph**](BillingAPI.md#GetVolumeBillingHistoryGraph) | **Get** /billing/billing/history/volume/{volume_id}/graph | Retrieve hourly cost datapoints of a Specific Volume for a specific billing cycle
[**ListBillingContractHistory**](BillingAPI.md#ListBillingContractHistory) | **Get** /billing/billing/history/contract | Retrieve Billing History of Contract for a specific Billing Cycle
[**ListBucketBillingHistory**](BillingAPI.md#ListBucketBillingHistory) | **Get** /billing/billing/history/bucket | Retrieve Billing History of a Bucket for a specific Billing Cycle
[**ListClustersBillingHistory**](BillingAPI.md#ListClustersBillingHistory) | **Get** /billing/billing/history/cluster | Retrieve Billing History of Clusters for a specific Billing Cycle
[**ListOrgNotificationThresholds**](BillingAPI.md#ListOrgNotificationThresholds) | **Get** /billing/billing/threshold | GET: All Thresholds for Organization
[**ListServerlessInferenceBillingHistory**](BillingAPI.md#ListServerlessInferenceBillingHistory) | **Get** /billing/billing/history/serverless_inference | Retrieve Billing History of serverless inference for a specific Billing Cycle
[**ListSnapshotBillingHistory**](BillingAPI.md#ListSnapshotBillingHistory) | **Get** /billing/billing/history/snapshot | Retrieve Billing History of Snapshot for a specific Billing Cycle



## BucketsBillingHistoryHourlyChart

> ResourceLevelGraphBillingDetailsBucket BucketsBillingHistoryHourlyChart(ctx, bucketId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.BucketsBillingHistoryHourlyChart(context.Background(), bucketId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BucketsBillingHistoryHourlyChart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BucketsBillingHistoryHourlyChart`: ResourceLevelGraphBillingDetailsBucket
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BucketsBillingHistoryHourlyChart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketsBillingHistoryHourlyChartRequest struct via the builder pattern


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


## GetBucketBillingHistory

> ResourceLevelBucketBillingDetailsResponseModel GetBucketBillingHistory(ctx, bucketId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetBucketBillingHistory(context.Background(), bucketId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBucketBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBucketBillingHistory`: ResourceLevelBucketBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBucketBillingHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketBillingHistoryRequest struct via the builder pattern


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


## GetClusterBillingHistory

> ResourceLevelClusterBillingDetailsResponseModel GetClusterBillingHistory(ctx, clusterId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetClusterBillingHistory(context.Background(), clusterId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetClusterBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterBillingHistory`: ResourceLevelClusterBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetClusterBillingHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterBillingHistoryRequest struct via the builder pattern


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


## GetClusterBillingHistoryGraph

> ResourceLevelClusterGraphBillingDetailsResponseModel GetClusterBillingHistoryGraph(ctx, clusterId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetClusterBillingHistoryGraph(context.Background(), clusterId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetClusterBillingHistoryGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterBillingHistoryGraph`: ResourceLevelClusterGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetClusterBillingHistoryGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterBillingHistoryGraphRequest struct via the builder pattern


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


## GetDataSynthesisBillingHistory

> TokenBasedBillingHistoryResponse GetDataSynthesisBillingHistory(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetDataSynthesisBillingHistory(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetDataSynthesisBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataSynthesisBillingHistory`: TokenBasedBillingHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetDataSynthesisBillingHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataSynthesisBillingHistoryRequest struct via the builder pattern


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


## GetDataSynthesisBillingHistoryGraph

> DataSynthesisBillingHistoryDetailsResponseSchema GetDataSynthesisBillingHistoryGraph(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetDataSynthesisBillingHistoryGraph(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetDataSynthesisBillingHistoryGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataSynthesisBillingHistoryGraph`: DataSynthesisBillingHistoryDetailsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetDataSynthesisBillingHistoryGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataSynthesisBillingHistoryGraphRequest struct via the builder pattern


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


## GetDataSynthesisHistoryForResource

> DataSynthesisBillingHistoryDetailsResponseSchema GetDataSynthesisHistoryForResource(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()





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
	resp, r, err := apiClient.BillingAPI.GetDataSynthesisHistoryForResource(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetDataSynthesisHistoryForResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataSynthesisHistoryForResource`: DataSynthesisBillingHistoryDetailsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetDataSynthesisHistoryForResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataSynthesisHistoryForResourceRequest struct via the builder pattern


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


## GetFineTuningBillingHistory

> WorkloadBillingHistoryResponse GetFineTuningBillingHistory(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetFineTuningBillingHistory(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetFineTuningBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFineTuningBillingHistory`: WorkloadBillingHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetFineTuningBillingHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFineTuningBillingHistoryRequest struct via the builder pattern


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


## GetFineTuningBillingHistoryGraph

> ResourceLevelVolumeGraphBillingDetailsResponseModel GetFineTuningBillingHistoryGraph(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetFineTuningBillingHistoryGraph(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetFineTuningBillingHistoryGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFineTuningBillingHistoryGraph`: ResourceLevelVolumeGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetFineTuningBillingHistoryGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFineTuningBillingHistoryGraphRequest struct via the builder pattern


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


## GetLastDayCost

> LastDayCostResponse GetLastDayCost(ctx).Execute()

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
	// response from `GetLastDayCost`: LastDayCostResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetLastDayCost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastDayCostRequest struct via the builder pattern


### Return type

[**LastDayCostResponse**](LastDayCostResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelEvaluationBillingHistory

> TokenBasedBillingHistoryResponse GetModelEvaluationBillingHistory(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetModelEvaluationBillingHistory(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetModelEvaluationBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelEvaluationBillingHistory`: TokenBasedBillingHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetModelEvaluationBillingHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetModelEvaluationBillingHistoryRequest struct via the builder pattern


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


## GetModelEvaluationBillingHistoryGraph

> ModelEvaluationBillingHistoryDetailsResponseSchema GetModelEvaluationBillingHistoryGraph(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve hourly cost datapoints of a Specific Model Evaluation for a specific



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
	resp, r, err := apiClient.BillingAPI.GetModelEvaluationBillingHistoryGraph(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetModelEvaluationBillingHistoryGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelEvaluationBillingHistoryGraph`: ModelEvaluationBillingHistoryDetailsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetModelEvaluationBillingHistoryGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelEvaluationBillingHistoryGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

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


## GetNotificationThreshold

> OrganizationThresholdUpdateResponse GetNotificationThreshold(ctx, thresholdId).Payload(payload).Execute()

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
	payload := *openapiclient.NewSubscribeOrUnsubscribeUpdatePayload(false) // SubscribeOrUnsubscribeUpdatePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetNotificationThreshold(context.Background(), thresholdId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetNotificationThreshold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationThreshold`: OrganizationThresholdUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetNotificationThreshold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thresholdId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**SubscribeOrUnsubscribeUpdatePayload**](SubscribeOrUnsubscribeUpdatePayload.md) |  | 

### Return type

[**OrganizationThresholdUpdateResponse**](OrganizationThresholdUpdateResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceFineTuningBillingHistory

> ResourceLevelVolumeBillingDetailsResponseModel GetResourceFineTuningBillingHistory(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetResourceFineTuningBillingHistory(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetResourceFineTuningBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceFineTuningBillingHistory`: ResourceLevelVolumeBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetResourceFineTuningBillingHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceFineTuningBillingHistoryRequest struct via the builder pattern


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


## GetResourceModelEvaluationBillingHistory

> ModelEvaluationBillingHistoryDetailsResponseSchema GetResourceModelEvaluationBillingHistory(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()





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
	resp, r, err := apiClient.BillingAPI.GetResourceModelEvaluationBillingHistory(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetResourceModelEvaluationBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceModelEvaluationBillingHistory`: ModelEvaluationBillingHistoryDetailsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetResourceModelEvaluationBillingHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceModelEvaluationBillingHistoryRequest struct via the builder pattern


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


## GetServerlessInferenceBillingHistoryGraph

> ServerlessInferencedBillingHistoryDetailsResponseSchema GetServerlessInferenceBillingHistoryGraph(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve hourly cost datapoints of a Specific Serverless Inference for a specific



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
	resp, r, err := apiClient.BillingAPI.GetServerlessInferenceBillingHistoryGraph(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetServerlessInferenceBillingHistoryGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerlessInferenceBillingHistoryGraph`: ServerlessInferencedBillingHistoryDetailsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetServerlessInferenceBillingHistoryGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerlessInferenceBillingHistoryGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

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


## GetServerlessInferencesBillingHistory

> ServerlessInferencedBillingHistoryDetailsResponseSchema GetServerlessInferencesBillingHistory(ctx, resourceId).StartDate(startDate).EndDate(endDate).Execute()





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
	resp, r, err := apiClient.BillingAPI.GetServerlessInferencesBillingHistory(context.Background(), resourceId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetServerlessInferencesBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerlessInferencesBillingHistory`: ServerlessInferencedBillingHistoryDetailsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetServerlessInferencesBillingHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerlessInferencesBillingHistoryRequest struct via the builder pattern


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


## GetSnapshotBillingHistory

> ResourceLevelVolumeBillingDetailsResponseModel GetSnapshotBillingHistory(ctx, snapshotId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetSnapshotBillingHistory(context.Background(), snapshotId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetSnapshotBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshotBillingHistory`: ResourceLevelVolumeBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetSnapshotBillingHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotBillingHistoryRequest struct via the builder pattern


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


## GetSnapshotBillingHistoryGraph

> ResourceLevelVolumeGraphBillingDetailsResponseModel GetSnapshotBillingHistoryGraph(ctx, snapshotId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetSnapshotBillingHistoryGraph(context.Background(), snapshotId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetSnapshotBillingHistoryGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshotBillingHistoryGraph`: ResourceLevelVolumeGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetSnapshotBillingHistoryGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotBillingHistoryGraphRequest struct via the builder pattern


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


## GetUsage

> BillingMetricesResponse GetUsage(ctx).Deleted(deleted).Environment(environment).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetUsage(context.Background()).Deleted(deleted).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsage`: BillingMetricesResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageRequest struct via the builder pattern


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


## GetUserBillingHistory

> OrganizationLevelBillingHistoryResponseModel GetUserBillingHistory(ctx).StartDate(startDate).EndDate(endDate).Graph(graph).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetUserBillingHistory(context.Background()).StartDate(startDate).EndDate(endDate).Graph(graph).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetUserBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBillingHistory`: OrganizationLevelBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetUserBillingHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **graph** | **string** | Set this value to \&quot;true\&quot; for getting graph value | 

### Return type

[**OrganizationLevelBillingHistoryResponseModel**](OrganizationLevelBillingHistoryResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMBillingDetails

> ResourceLevelVMBillingDetailsResponseModel GetVMBillingDetails(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetVMBillingDetails(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetVMBillingDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMBillingDetails`: ResourceLevelVMBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetVMBillingDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMBillingDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVMBillingDetailsResponseModel**](ResourceLevelVMBillingDetailsResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMBillingEvents

> ResourceBillingEventsHistoryResponse GetVMBillingEvents(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetVMBillingEvents(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetVMBillingEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMBillingEvents`: ResourceBillingEventsHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetVMBillingEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMBillingEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceBillingEventsHistoryResponse**](ResourceBillingEventsHistoryResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMBillingGraph

> ResourceLevelVmGraphBillingDetailsResponseModel GetVMBillingGraph(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetVMBillingGraph(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetVMBillingGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMBillingGraph`: ResourceLevelVmGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetVMBillingGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMBillingGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceLevelVmGraphBillingDetailsResponseModel**](ResourceLevelVmGraphBillingDetailsResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMBillingHistory

> ResourceLevelVmBillingHistoryResponseModel GetVMBillingHistory(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetVMBillingHistory(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetVMBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMBillingHistory`: ResourceLevelVmBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetVMBillingHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVMBillingHistoryRequest struct via the builder pattern


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


## GetVMSubResourceCosts

> SubResourcesCostsResponseModel GetVMSubResourceCosts(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetVMSubResourceCosts(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetVMSubResourceCosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMSubResourceCosts`: SubResourcesCostsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetVMSubResourceCosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMSubResourceCostsRequest struct via the builder pattern


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


## GetVMSubResourceGraph

> SubResourcesGraphResponseModel GetVMSubResourceGraph(ctx, vmId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetVMSubResourceGraph(context.Background(), vmId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetVMSubResourceGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMSubResourceGraph`: SubResourcesGraphResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetVMSubResourceGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMSubResourceGraphRequest struct via the builder pattern


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


## GetVolumeBillingDetails

> ResourceLevelVolumeBillingDetailsResponseModel GetVolumeBillingDetails(ctx, volumeId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetVolumeBillingDetails(context.Background(), volumeId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetVolumeBillingDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVolumeBillingDetails`: ResourceLevelVolumeBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetVolumeBillingDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeBillingDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

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


## GetVolumeBillingEvents

> ResourceBillingEventsHistoryResponse GetVolumeBillingEvents(ctx, volumeId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetVolumeBillingEvents(context.Background(), volumeId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetVolumeBillingEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVolumeBillingEvents`: ResourceBillingEventsHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetVolumeBillingEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeBillingEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ResourceBillingEventsHistoryResponse**](ResourceBillingEventsHistoryResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolumeBillingHistory

> ResourceLevelVolumeBillingHistoryResponseModel GetVolumeBillingHistory(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetVolumeBillingHistory(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetVolumeBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVolumeBillingHistory`: ResourceLevelVolumeBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetVolumeBillingHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeBillingHistoryRequest struct via the builder pattern


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


## GetVolumeBillingHistoryGraph

> ResourceLevelVolumeGraphBillingDetailsResponseModel GetVolumeBillingHistoryGraph(ctx, volumeId).StartDate(startDate).EndDate(endDate).Execute()

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
	resp, r, err := apiClient.BillingAPI.GetVolumeBillingHistoryGraph(context.Background(), volumeId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetVolumeBillingHistoryGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVolumeBillingHistoryGraph`: ResourceLevelVolumeGraphBillingDetailsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetVolumeBillingHistoryGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeBillingHistoryGraphRequest struct via the builder pattern


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


## ListBillingContractHistory

> ListBillingContractHistory(ctx).StartDate(startDate).EndDate(endDate).Search(search).Execute()

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
	r, err := apiClient.BillingAPI.ListBillingContractHistory(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListBillingContractHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBillingContractHistoryRequest struct via the builder pattern


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


## ListBucketBillingHistory

> ResourceLevelBucketBillingHistoryResponseModel ListBucketBillingHistory(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

Retrieve Billing History of a Bucket for a specific Billing Cycle



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
	resp, r, err := apiClient.BillingAPI.ListBucketBillingHistory(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListBucketBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBucketBillingHistory`: ResourceLevelBucketBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListBucketBillingHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBucketBillingHistoryRequest struct via the builder pattern


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


## ListClustersBillingHistory

> ResourceLevelClusterBillingHistoryResponseModel ListClustersBillingHistory(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

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
	resp, r, err := apiClient.BillingAPI.ListClustersBillingHistory(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListClustersBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClustersBillingHistory`: ResourceLevelClusterBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListClustersBillingHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClustersBillingHistoryRequest struct via the builder pattern


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


## ListOrgNotificationThresholds

> OrganizationThresholdsResponse ListOrgNotificationThresholds(ctx).Execute()

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
	resp, r, err := apiClient.BillingAPI.ListOrgNotificationThresholds(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListOrgNotificationThresholds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgNotificationThresholds`: OrganizationThresholdsResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListOrgNotificationThresholds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgNotificationThresholdsRequest struct via the builder pattern


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


## ListServerlessInferenceBillingHistory

> TokenBasedBillingHistoryResponse ListServerlessInferenceBillingHistory(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

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
	resp, r, err := apiClient.BillingAPI.ListServerlessInferenceBillingHistory(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListServerlessInferenceBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServerlessInferenceBillingHistory`: TokenBasedBillingHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListServerlessInferenceBillingHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServerlessInferenceBillingHistoryRequest struct via the builder pattern


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


## ListSnapshotBillingHistory

> ResourceLevelVolumeBillingHistoryResponseModel ListSnapshotBillingHistory(ctx).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()

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
	resp, r, err := apiClient.BillingAPI.ListSnapshotBillingHistory(context.Background()).StartDate(startDate).EndDate(endDate).Search(search).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListSnapshotBillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSnapshotBillingHistory`: ResourceLevelVolumeBillingHistoryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListSnapshotBillingHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSnapshotBillingHistoryRequest struct via the builder pattern


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

