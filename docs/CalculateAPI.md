# \CalculateAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculateResourceBillingRate**](CalculateAPI.md#CalculateResourceBillingRate) | **Get** /pricebook/calculate/resource/{resource_type}/{id} | Retrieve Billing Rate for Resource



## CalculateResourceBillingRate

> ResourceBillingResponseForCustomer CalculateResourceBillingRate(ctx, resourceType, id).Execute()

Retrieve Billing Rate for Resource



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
	resourceType := "resourceType_example" // string | 
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalculateAPI.CalculateResourceBillingRate(context.Background(), resourceType, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalculateAPI.CalculateResourceBillingRate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculateResourceBillingRate`: ResourceBillingResponseForCustomer
	fmt.Fprintf(os.Stdout, "Response from `CalculateAPI.CalculateResourceBillingRate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCalculateResourceBillingRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResourceBillingResponseForCustomer**](ResourceBillingResponseForCustomer.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

