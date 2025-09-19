# \GpuAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**ListGPUs**](GpuAPI.md#ListGPUs) | **Get** /core/gpus | List GPUs



## ListGPUs

> GPUList ListGPUs(ctx).Execute()
=======
[**ListGpus**](GpuAPI.md#ListGpus) | **Get** /core/gpus | List GPUs



## ListGpus

> GPUList ListGpus(ctx).Execute()
>>>>>>> main

List GPUs



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
	resp, r, err := apiClient.GpuAPI.ListGPUs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpuAPI.ListGPUs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGPUs`: GPUList
	fmt.Fprintf(os.Stdout, "Response from `GpuAPI.ListGPUs`: %v\n", resp)
=======
	resp, r, err := apiClient.GpuAPI.ListGpus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpuAPI.ListGpus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGpus`: GPUList
	fmt.Fprintf(os.Stdout, "Response from `GpuAPI.ListGpus`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiListGPUsRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiListGpusRequest struct via the builder pattern
>>>>>>> main


### Return type

[**GPUList**](GPUList.md)

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

