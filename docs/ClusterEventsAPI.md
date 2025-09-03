# \ClusterEventsAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAllOfAClusterEvents**](ClusterEventsAPI.md#FetchAllOfAClusterEvents) | **Get** /core/clusters/{cluster_id}/events | Fetch all of a cluster events



## FetchAllOfAClusterEvents

> ClusterEvents FetchAllOfAClusterEvents(ctx, clusterId).Execute()

Fetch all of a cluster events

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
	clusterId := "clusterId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterEventsAPI.FetchAllOfAClusterEvents(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterEventsAPI.FetchAllOfAClusterEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchAllOfAClusterEvents`: ClusterEvents
	fmt.Fprintf(os.Stdout, "Response from `ClusterEventsAPI.FetchAllOfAClusterEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllOfAClusterEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterEvents**](ClusterEvents.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

