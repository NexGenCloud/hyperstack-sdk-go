# \RegionsAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRegionsEndpoint**](RegionsAPI.md#ListRegionsEndpoint) | **Get** /object-storage/regions | Get a list of supported regions



## ListRegionsEndpoint

> ObjectStorageRegionListResponse ListRegionsEndpoint(ctx).Execute()

Get a list of supported regions

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
	resp, r, err := apiClient.RegionsAPI.ListRegionsEndpoint(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionsAPI.ListRegionsEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRegionsEndpoint`: ObjectStorageRegionListResponse
	fmt.Fprintf(os.Stdout, "Response from `RegionsAPI.ListRegionsEndpoint`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRegionsEndpointRequest struct via the builder pattern


### Return type

[**ObjectStorageRegionListResponse**](ObjectStorageRegionListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

