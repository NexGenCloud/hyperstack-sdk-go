# \StockAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**RetrieveGPUStocks**](StockAPI.md#RetrieveGPUStocks) | **Get** /core/stocks | Retrieve GPU stocks



## RetrieveGPUStocks

> NewStockRetriveResponse RetrieveGPUStocks(ctx).Execute()
=======
[**RetrieveGpuStocks**](StockAPI.md#RetrieveGpuStocks) | **Get** /core/stocks | Retrieve GPU stocks



## RetrieveGpuStocks

> NewStockRetriveResponse RetrieveGpuStocks(ctx).Execute()
>>>>>>> main

Retrieve GPU stocks



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
	resp, r, err := apiClient.StockAPI.RetrieveGPUStocks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StockAPI.RetrieveGPUStocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveGPUStocks`: NewStockRetriveResponse
	fmt.Fprintf(os.Stdout, "Response from `StockAPI.RetrieveGPUStocks`: %v\n", resp)
=======
	resp, r, err := apiClient.StockAPI.RetrieveGpuStocks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StockAPI.RetrieveGpuStocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveGpuStocks`: NewStockRetriveResponse
	fmt.Fprintf(os.Stdout, "Response from `StockAPI.RetrieveGpuStocks`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiRetrieveGPUStocksRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveGpuStocksRequest struct via the builder pattern
>>>>>>> main


### Return type

[**NewStockRetriveResponse**](NewStockRetriveResponse.md)

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

