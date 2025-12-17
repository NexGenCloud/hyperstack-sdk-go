# \BucketsAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBucketEndpoint**](BucketsAPI.md#DeleteBucketEndpoint) | **Delete** /object-storage/buckets/{bucket_name} | Delete a bucket
[**ListBucketsEndpoint**](BucketsAPI.md#ListBucketsEndpoint) | **Get** /object-storage/buckets | List buckets
[**RetrieveBucketEndpoint**](BucketsAPI.md#RetrieveBucketEndpoint) | **Get** /object-storage/buckets/{bucket_name} | Retrieve a bucket



## DeleteBucketEndpoint

> ObjectStorageDeleteResponse DeleteBucketEndpoint(ctx, bucketName).Region(region).Execute()

Delete a bucket

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
	bucketName := "bucketName_example" // string | 
	region := "region_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketsAPI.DeleteBucketEndpoint(context.Background(), bucketName).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketsAPI.DeleteBucketEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBucketEndpoint`: ObjectStorageDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `BucketsAPI.DeleteBucketEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** |  | 

### Return type

[**ObjectStorageDeleteResponse**](ObjectStorageDeleteResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBucketsEndpoint

> ObjectStorageBucketListResponse ListBucketsEndpoint(ctx).Search(search).Execute()

List buckets

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
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketsAPI.ListBucketsEndpoint(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketsAPI.ListBucketsEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBucketsEndpoint`: ObjectStorageBucketListResponse
	fmt.Fprintf(os.Stdout, "Response from `BucketsAPI.ListBucketsEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBucketsEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 

### Return type

[**ObjectStorageBucketListResponse**](ObjectStorageBucketListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBucketEndpoint

> ObjectStorageBucketResponse RetrieveBucketEndpoint(ctx, bucketName).Region(region).Execute()

Retrieve a bucket

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
	bucketName := "bucketName_example" // string | 
	region := "region_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketsAPI.RetrieveBucketEndpoint(context.Background(), bucketName).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketsAPI.RetrieveBucketEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBucketEndpoint`: ObjectStorageBucketResponse
	fmt.Fprintf(os.Stdout, "Response from `BucketsAPI.RetrieveBucketEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBucketEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** |  | 

### Return type

[**ObjectStorageBucketResponse**](ObjectStorageBucketResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

