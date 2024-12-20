# \CustomerContractAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListContracts**](CustomerContractAPI.md#ListContracts) | **Get** /pricebook/contracts | List Contracts
[**RetrieveContractDetails**](CustomerContractAPI.md#RetrieveContractDetails) | **Get** /pricebook/contracts/{contract_id} | Retrieve Contract Details
[**RetrieveGpuAllocationGraphForContract**](CustomerContractAPI.md#RetrieveGpuAllocationGraphForContract) | **Get** /pricebook/contracts/{contract_id}/gpu_allocation_graph | Retrieve GPU Allocation Graph for Contract



## ListContracts

> GetCustomerContractsListResponseModel ListContracts(ctx).Page(page).PerPage(perPage).Execute()

List Contracts



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
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerContractAPI.ListContracts(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerContractAPI.ListContracts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContracts`: GetCustomerContractsListResponseModel
	fmt.Fprintf(os.Stdout, "Response from `CustomerContractAPI.ListContracts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**GetCustomerContractsListResponseModel**](GetCustomerContractsListResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveContractDetails

> CustomerContractDetailResponseModel RetrieveContractDetails(ctx, contractId).Execute()

Retrieve Contract Details



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerContractAPI.RetrieveContractDetails(context.Background(), contractId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerContractAPI.RetrieveContractDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveContractDetails`: CustomerContractDetailResponseModel
	fmt.Fprintf(os.Stdout, "Response from `CustomerContractAPI.RetrieveContractDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveContractDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerContractDetailResponseModel**](CustomerContractDetailResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveGpuAllocationGraphForContract

> ContractGPUAllocationGraphResponse RetrieveGpuAllocationGraphForContract(ctx, contractId).StartDate(startDate).EndDate(endDate).Execute()

Retrieve GPU Allocation Graph for Contract



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
	startDate := "startDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)
	endDate := "endDate_example" // string | Date should be formatted in YYYY-MM-DDTHH:MM:SS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerContractAPI.RetrieveGpuAllocationGraphForContract(context.Background(), contractId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerContractAPI.RetrieveGpuAllocationGraphForContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveGpuAllocationGraphForContract`: ContractGPUAllocationGraphResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerContractAPI.RetrieveGpuAllocationGraphForContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveGpuAllocationGraphForContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 
 **endDate** | **string** | Date should be formatted in YYYY-MM-DDTHH:MM:SS | 

### Return type

[**ContractGPUAllocationGraphResponse**](ContractGPUAllocationGraphResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

