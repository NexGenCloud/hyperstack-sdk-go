# \VolumeAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolume**](VolumeAPI.md#CreateVolume) | **Post** /core/volumes | Create volume
[**DeleteVolume**](VolumeAPI.md#DeleteVolume) | **Delete** /core/volumes/{volume_id} | Delete volume
[**FetchVolumeDetails**](VolumeAPI.md#FetchVolumeDetails) | **Get** /core/volumes/{volume_id} | Fetch Volume Details
[**FetchVolumeNameAvailability**](VolumeAPI.md#FetchVolumeNameAvailability) | **Get** /core/volume/name-availability/{name} | Fetch volume name availability
[**ListVolumeTypes**](VolumeAPI.md#ListVolumeTypes) | **Get** /core/volume-types | List volume types
[**ListVolumes**](VolumeAPI.md#ListVolumes) | **Get** /core/volumes | List volumes
[**UpdateVolume**](VolumeAPI.md#UpdateVolume) | **Patch** /core/volumes/{volume_id} | Update volume fields



## CreateVolume

> Volume CreateVolume(ctx).Payload(payload).Execute()

Create volume



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
	payload := *openapiclient.NewCreateVolumePayload("EnvironmentName_example", "Name_example", int32(123), "VolumeType_example") // CreateVolumePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.CreateVolume(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.CreateVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolume`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.CreateVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateVolumePayload**](CreateVolumePayload.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolume

> ResponseModel DeleteVolume(ctx, volumeId).Execute()

Delete volume



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.DeleteVolume(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.DeleteVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVolume`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.DeleteVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeRequest struct via the builder pattern


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


## FetchVolumeDetails

> Volume FetchVolumeDetails(ctx, volumeId).Execute()

Fetch Volume Details



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.FetchVolumeDetails(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.FetchVolumeDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchVolumeDetails`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.FetchVolumeDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchVolumeDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Volume**](Volume.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVolumeNameAvailability

> NameAvailableModel FetchVolumeNameAvailability(ctx, name).Execute()

Fetch volume name availability



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
	resp, r, err := apiClient.VolumeAPI.FetchVolumeNameAvailability(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.FetchVolumeNameAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchVolumeNameAvailability`: NameAvailableModel
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.FetchVolumeNameAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchVolumeNameAvailabilityRequest struct via the builder pattern


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


## ListVolumeTypes

> VolumeTypes ListVolumeTypes(ctx).Execute()

List volume types



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
	resp, r, err := apiClient.VolumeAPI.ListVolumeTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.ListVolumeTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumeTypes`: VolumeTypes
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.ListVolumeTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVolumeTypesRequest struct via the builder pattern


### Return type

[**VolumeTypes**](VolumeTypes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumes

> Volumes ListVolumes(ctx).Page(page).PageSize(pageSize).Search(search).Environment(environment).Execute()

List volumes



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
	page := "page_example" // string | Page Number (optional)
	pageSize := "pageSize_example" // string | Data Per Page (optional)
	search := "search_example" // string |  (optional)
	environment := "environment_example" // string | Filter Environment ID or Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.ListVolumes(context.Background()).Page(page).PageSize(pageSize).Search(search).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.ListVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumes`: Volumes
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.ListVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page Number | 
 **pageSize** | **string** | Data Per Page | 
 **search** | **string** |  | 
 **environment** | **string** | Filter Environment ID or Name | 

### Return type

[**Volumes**](Volumes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVolume

> UpdateVolumeResponse UpdateVolume(ctx, volumeId).Payload(payload).Execute()

Update volume fields



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
	payload := *openapiclient.NewUpdateVolumePayload("EnvironmentName_example") // UpdateVolumePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.UpdateVolume(context.Background(), volumeId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.UpdateVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVolume`: UpdateVolumeResponse
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.UpdateVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**UpdateVolumePayload**](UpdateVolumePayload.md) |  | 

### Return type

[**UpdateVolumeResponse**](UpdateVolumeResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

