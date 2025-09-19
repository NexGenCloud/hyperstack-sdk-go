# \ImageAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**DeleteImage**](ImageAPI.md#DeleteImage) | **Delete** /core/images/{id} | Delete an image
[**FetchImageNameAvailability**](ImageAPI.md#FetchImageNameAvailability) | **Get** /core/image/name-availability/{name} | Fetch name availability for Images
[**GetImageDetails**](ImageAPI.md#GetImageDetails) | **Get** /core/images/{id} | Get Private Image Details
[**ListImages2**](ImageAPI.md#ListImages2) | **Get** /core/images | List Images



## DeleteImage

> ResponseModel DeleteImage(ctx, id).Execute()
=======
[**DeleteAnImage**](ImageAPI.md#DeleteAnImage) | **Delete** /core/images/{id} | Delete an image
[**FetchNameAvailabilityForImages**](ImageAPI.md#FetchNameAvailabilityForImages) | **Get** /core/image/name-availability/{name} | Fetch name availability for Images
[**GetPrivateImageDetails**](ImageAPI.md#GetPrivateImageDetails) | **Get** /core/images/{id} | Get Private Image Details
[**ListImages**](ImageAPI.md#ListImages) | **Get** /core/images | List Images



## DeleteAnImage

> ResponseModel DeleteAnImage(ctx, id).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.ImageAPI.DeleteImage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.DeleteImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteImage`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.DeleteImage`: %v\n", resp)
=======
	resp, r, err := apiClient.ImageAPI.DeleteAnImage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.DeleteAnImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAnImage`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.DeleteAnImage`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiDeleteImageRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiDeleteAnImageRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseModel**](ResponseModel.md)

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
## FetchImageNameAvailability

> NameAvailableModel FetchImageNameAvailability(ctx, name).Execute()
=======
## FetchNameAvailabilityForImages

> NameAvailableModel FetchNameAvailabilityForImages(ctx, name).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.ImageAPI.FetchImageNameAvailability(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.FetchImageNameAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchImageNameAvailability`: NameAvailableModel
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.FetchImageNameAvailability`: %v\n", resp)
=======
	resp, r, err := apiClient.ImageAPI.FetchNameAvailabilityForImages(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.FetchNameAvailabilityForImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchNameAvailabilityForImages`: NameAvailableModel
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.FetchNameAvailabilityForImages`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiFetchImageNameAvailabilityRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiFetchNameAvailabilityForImagesRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NameAvailableModel**](NameAvailableModel.md)

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
## GetImageDetails

> Image GetImageDetails(ctx, id).IncludeRelatedVms(includeRelatedVms).Execute()
=======
## GetPrivateImageDetails

> Image GetPrivateImageDetails(ctx, id).IncludeRelatedVms(includeRelatedVms).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.ImageAPI.GetImageDetails(context.Background(), id).IncludeRelatedVms(includeRelatedVms).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetImageDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageDetails`: Image
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetImageDetails`: %v\n", resp)
=======
	resp, r, err := apiClient.ImageAPI.GetPrivateImageDetails(context.Background(), id).IncludeRelatedVms(includeRelatedVms).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetPrivateImageDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateImageDetails`: Image
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetPrivateImageDetails`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGetImageDetailsRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiGetPrivateImageDetailsRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeRelatedVms** | **bool** |  | 

### Return type

[**Image**](Image.md)

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
## ListImages2

> Images ListImages2(ctx).Region(region).IncludePublic(includePublic).Search(search).Page(page).PerPage(perPage).Execute()
=======
## ListImages

> Images ListImages(ctx).Region(region).IncludePublic(includePublic).Search(search).Page(page).PerPage(perPage).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.ImageAPI.ListImages2(context.Background()).Region(region).IncludePublic(includePublic).Search(search).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.ListImages2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImages2`: Images
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.ListImages2`: %v\n", resp)
=======
	resp, r, err := apiClient.ImageAPI.ListImages(context.Background()).Region(region).IncludePublic(includePublic).Search(search).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.ListImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImages`: Images
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.ListImages`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters



### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiListImages2Request struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiListImagesRequest struct via the builder pattern
>>>>>>> main


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

