# \PricebookAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**GetPricebook**](PricebookAPI.md#GetPricebook) | **Get** /pricebook | 



## GetPricebook

> []PricebookModel GetPricebook(ctx).Execute()


=======
[**RetrivePricebook**](PricebookAPI.md#RetrivePricebook) | **Get** /pricebook | 



## RetrivePricebook

> []PricebookModel RetrivePricebook(ctx).Execute()
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.PricebookAPI.GetPricebook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricebookAPI.GetPricebook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricebook`: []PricebookModel
	fmt.Fprintf(os.Stdout, "Response from `PricebookAPI.GetPricebook`: %v\n", resp)
=======
	resp, r, err := apiClient.PricebookAPI.RetrivePricebook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricebookAPI.RetrivePricebook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrivePricebook`: []PricebookModel
	fmt.Fprintf(os.Stdout, "Response from `PricebookAPI.RetrivePricebook`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGetPricebookRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrivePricebookRequest struct via the builder pattern
>>>>>>> main


### Return type

[**[]PricebookModel**](PricebookModel.md)

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

