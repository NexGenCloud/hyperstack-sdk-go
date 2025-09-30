# \SnapshotsAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateANewCustomImage**](SnapshotsAPI.md#CreateANewCustomImage) | **Post** /core/snapshots/{snapshot_id}/image | Create an image from a snapshot
[**DeleteAnExistingSnapshot**](SnapshotsAPI.md#DeleteAnExistingSnapshot) | **Delete** /core/snapshots/{id} | Delete snapshot
[**FetchSnapshotNameAvailability**](SnapshotsAPI.md#FetchSnapshotNameAvailability) | **Get** /core/snapshots/name-availability/{name} | Fetch snapshot name availability
[**RestoreASnapshot**](SnapshotsAPI.md#RestoreASnapshot) | **Post** /core/snapshots/{id}/restore | Restore a snapshot
[**RetrieveAnExistingSnapshot**](SnapshotsAPI.md#RetrieveAnExistingSnapshot) | **Get** /core/snapshots/{id} | Retrieve a snapshot
[**RetrievesAListOfSnapshots**](SnapshotsAPI.md#RetrievesAListOfSnapshots) | **Get** /core/snapshots | Retrieve list of snapshots with pagination



## CreateANewCustomImage

> CreateImage CreateANewCustomImage(ctx, snapshotId).Payload(payload).Execute()

Create an image from a snapshot



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
	payload := *openapiclient.NewCreateImagePayload("Name_example") // CreateImagePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.CreateANewCustomImage(context.Background(), snapshotId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.CreateANewCustomImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateANewCustomImage`: CreateImage
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.CreateANewCustomImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateANewCustomImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**CreateImagePayload**](CreateImagePayload.md) |  | 

### Return type

[**CreateImage**](CreateImage.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnExistingSnapshot

> ResponseModel DeleteAnExistingSnapshot(ctx, id).Execute()

Delete snapshot



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
	resp, r, err := apiClient.SnapshotsAPI.DeleteAnExistingSnapshot(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.DeleteAnExistingSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAnExistingSnapshot`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.DeleteAnExistingSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnExistingSnapshotRequest struct via the builder pattern


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


## FetchSnapshotNameAvailability

> NameAvailableModel FetchSnapshotNameAvailability(ctx, name).Execute()

Fetch snapshot name availability



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
	resp, r, err := apiClient.SnapshotsAPI.FetchSnapshotNameAvailability(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.FetchSnapshotNameAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchSnapshotNameAvailability`: NameAvailableModel
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.FetchSnapshotNameAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSnapshotNameAvailabilityRequest struct via the builder pattern


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


## RestoreASnapshot

> Instance RestoreASnapshot(ctx, id).Payload(payload).Execute()

Restore a snapshot



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
	payload := *openapiclient.NewSnapshotRestoreRequest("NewVmName_example") // SnapshotRestoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.RestoreASnapshot(context.Background(), id).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RestoreASnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreASnapshot`: Instance
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.RestoreASnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreASnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**SnapshotRestoreRequest**](SnapshotRestoreRequest.md) |  | 

### Return type

[**Instance**](Instance.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAnExistingSnapshot

> SnapshotRetrieve RetrieveAnExistingSnapshot(ctx, id).Execute()

Retrieve a snapshot



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
	resp, r, err := apiClient.SnapshotsAPI.RetrieveAnExistingSnapshot(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RetrieveAnExistingSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveAnExistingSnapshot`: SnapshotRetrieve
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.RetrieveAnExistingSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAnExistingSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SnapshotRetrieve**](SnapshotRetrieve.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievesAListOfSnapshots

> Snapshots RetrievesAListOfSnapshots(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

Retrieve list of snapshots with pagination



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
	search := "search_example" // string | Search By Snapshot ID or Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.RetrievesAListOfSnapshots(context.Background()).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RetrievesAListOfSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrievesAListOfSnapshots`: Snapshots
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.RetrievesAListOfSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrievesAListOfSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page Number | 
 **pageSize** | **string** | Data Per Page | 
 **search** | **string** | Search By Snapshot ID or Name | 

### Return type

[**Snapshots**](Snapshots.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

