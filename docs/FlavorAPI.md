# \FlavorAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFlavors**](FlavorAPI.md#ListFlavors) | **Get** /core/flavors | List Flavors



## ListFlavors

> FlavorListResponse ListFlavors(ctx).Region(region).Execute()

List Flavors



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
	region := "region_example" // string | Region Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlavorAPI.ListFlavors(context.Background()).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlavorAPI.ListFlavors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFlavors`: FlavorListResponse
	fmt.Fprintf(os.Stdout, "Response from `FlavorAPI.ListFlavors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** | Region Name | 

### Return type

[**FlavorListResponse**](FlavorListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

