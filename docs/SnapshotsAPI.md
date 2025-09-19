# \SnapshotsAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**CreateANewCustomImage**](SnapshotsAPI.md#CreateANewCustomImage) | **Post** /core/snapshots/{snapshot_id}/image | Create an image from a snapshot
[**DeleteSnapshot**](SnapshotsAPI.md#DeleteSnapshot) | **Delete** /core/snapshots/{id} | Delete snapshot
[**FetchSnapshotNameAvailability**](SnapshotsAPI.md#FetchSnapshotNameAvailability) | **Get** /core/snapshots/name-availability/{name} | Fetch snapshot name availability
[**GetSnapshot**](SnapshotsAPI.md#GetSnapshot) | **Get** /core/snapshots/{id} | Retrieve a snapshot
[**GetSnapshots**](SnapshotsAPI.md#GetSnapshots) | **Get** /core/snapshots | Retrieve list of snapshots with pagination
[**PostSnapshotRestore**](SnapshotsAPI.md#PostSnapshotRestore) | **Post** /core/snapshots/{id}/restore | Restore a snapshot



## CreateANewCustomImage

> CreateImage CreateANewCustomImage(ctx, snapshotId).Payload(payload).Execute()
=======
[**CreateAnImageFromASnapshot**](SnapshotsAPI.md#CreateAnImageFromASnapshot) | **Post** /core/snapshots/{snapshot_id}/image | Create an image from a snapshot
[**DeleteSnapshot**](SnapshotsAPI.md#DeleteSnapshot) | **Delete** /core/snapshots/{id} | Delete snapshot
[**FetchSnapshotNameAvailability**](SnapshotsAPI.md#FetchSnapshotNameAvailability) | **Get** /core/snapshots/name-availability/{name} | Fetch snapshot name availability
[**RestoreASnapshot**](SnapshotsAPI.md#RestoreASnapshot) | **Post** /core/snapshots/{id}/restore | Restore a snapshot
[**RetrieveASnapshot**](SnapshotsAPI.md#RetrieveASnapshot) | **Get** /core/snapshots/{id} | Retrieve a snapshot
[**RetrieveListOfSnapshotsWithPagination**](SnapshotsAPI.md#RetrieveListOfSnapshotsWithPagination) | **Get** /core/snapshots | Retrieve list of snapshots with pagination



## CreateAnImageFromASnapshot

> CreateImage CreateAnImageFromASnapshot(ctx, snapshotId).Payload(payload).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.SnapshotsAPI.CreateANewCustomImage(context.Background(), snapshotId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.CreateANewCustomImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateANewCustomImage`: CreateImage
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.CreateANewCustomImage`: %v\n", resp)
=======
	resp, r, err := apiClient.SnapshotsAPI.CreateAnImageFromASnapshot(context.Background(), snapshotId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.CreateAnImageFromASnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAnImageFromASnapshot`: CreateImage
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.CreateAnImageFromASnapshot`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiCreateANewCustomImageRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiCreateAnImageFromASnapshotRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**CreateImagePayload**](CreateImagePayload.md) |  | 

### Return type

[**CreateImage**](CreateImage.md)

### Authorization

<<<<<<< HEAD
[apiKey](../README.md#apiKey)
=======
[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)
>>>>>>> main

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSnapshot

> ResponseModel DeleteSnapshot(ctx, id).Execute()

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
	resp, r, err := apiClient.SnapshotsAPI.DeleteSnapshot(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.DeleteSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSnapshot`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.DeleteSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotRequest struct via the builder pattern


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
## GetSnapshot

> SnapshotRetrieve GetSnapshot(ctx, id).Execute()
=======
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

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveASnapshot

> SnapshotRetrieve RetrieveASnapshot(ctx, id).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.SnapshotsAPI.GetSnapshot(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.GetSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshot`: SnapshotRetrieve
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.GetSnapshot`: %v\n", resp)
=======
	resp, r, err := apiClient.SnapshotsAPI.RetrieveASnapshot(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RetrieveASnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveASnapshot`: SnapshotRetrieve
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.RetrieveASnapshot`: %v\n", resp)
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
Other parameters are passed through a pointer to a apiGetSnapshotRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveASnapshotRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SnapshotRetrieve**](SnapshotRetrieve.md)

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
## GetSnapshots

> Snapshots GetSnapshots(ctx).Page(page).PageSize(pageSize).Search(search).Execute()
=======
## RetrieveListOfSnapshotsWithPagination

> Snapshots RetrieveListOfSnapshotsWithPagination(ctx).Page(page).PageSize(pageSize).Search(search).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.SnapshotsAPI.GetSnapshots(context.Background()).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.GetSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshots`: Snapshots
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.GetSnapshots`: %v\n", resp)
=======
	resp, r, err := apiClient.SnapshotsAPI.RetrieveListOfSnapshotsWithPagination(context.Background()).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RetrieveListOfSnapshotsWithPagination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveListOfSnapshotsWithPagination`: Snapshots
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.RetrieveListOfSnapshotsWithPagination`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters



### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGetSnapshotsRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveListOfSnapshotsWithPaginationRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page Number | 
 **pageSize** | **string** | Data Per Page | 
 **search** | **string** | Search By Snapshot ID or Name | 

### Return type

[**Snapshots**](Snapshots.md)

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

## PostSnapshotRestore

> Instance PostSnapshotRestore(ctx, id).Payload(payload).Execute()

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
	resp, r, err := apiClient.SnapshotsAPI.PostSnapshotRestore(context.Background(), id).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.PostSnapshotRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSnapshotRestore`: Instance
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.PostSnapshotRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSnapshotRestoreRequest struct via the builder pattern


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

=======
>>>>>>> main
