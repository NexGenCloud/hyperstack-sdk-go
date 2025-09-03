# \ImageAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteImage**](ImageAPI.md#DeleteImage) | **Delete** /core/images/{id} | Delete an image
[**FetchImageNameAvailability**](ImageAPI.md#FetchImageNameAvailability) | **Get** /core/image/name-availability/{name} | Fetch name availability for Images
[**GetImageDetails**](ImageAPI.md#GetImageDetails) | **Get** /core/images/{id} | Get Private Image Details
[**ListImages2**](ImageAPI.md#ListImages2) | **Get** /core/images | List Images



## DeleteImage

> ResponseModel DeleteImage(ctx, id).Execute()

Delete an image



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
	resp, r, err := apiClient.ImageAPI.DeleteImage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.DeleteImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteImage`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.DeleteImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRequest struct via the builder pattern


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


## FetchImageNameAvailability

> NameAvailableModel FetchImageNameAvailability(ctx, name).Execute()

Fetch name availability for Images



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
	resp, r, err := apiClient.ImageAPI.FetchImageNameAvailability(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.FetchImageNameAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchImageNameAvailability`: NameAvailableModel
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.FetchImageNameAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchImageNameAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NameAvailableModel**](NameAvailableModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageDetails

> Image GetImageDetails(ctx, id).IncludeRelatedVms(includeRelatedVms).Execute()

Get Private Image Details



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
	includeRelatedVms := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetImageDetails(context.Background(), id).IncludeRelatedVms(includeRelatedVms).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetImageDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageDetails`: Image
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetImageDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeRelatedVms** | **bool** |  | 

### Return type

[**Image**](Image.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImages2

> Images ListImages2(ctx).Region(region).IncludePublic(includePublic).Search(search).Page(page).PerPage(perPage).Execute()

List Images



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
	includePublic := true // bool | Flag to include public images in the response (true/false). Default is true. (optional)
	search := "search_example" // string | Search query to filter images by name (optional)
	page := int32(56) // int32 | Page number for pagination (optional)
	perPage := int32(56) // int32 | Number of Images per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.ListImages2(context.Background()).Region(region).IncludePublic(includePublic).Search(search).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.ListImages2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImages2`: Images
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.ListImages2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListImages2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** | Region Name | 
 **includePublic** | **bool** | Flag to include public images in the response (true/false). Default is true. | 
 **search** | **string** | Search query to filter images by name | 
 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of Images per page | 

### Return type

[**Images**](Images.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

