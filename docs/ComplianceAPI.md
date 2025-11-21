# \ComplianceAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompliance**](ComplianceAPI.md#CreateCompliance) | **Post** /core/compliance | Create compliance
[**DeleteCompliance**](ComplianceAPI.md#DeleteCompliance) | **Delete** /core/compliance/{gpu_model} | Delete a compliance
[**GetCompliance**](ComplianceAPI.md#GetCompliance) | **Get** /core/compliance | Retrieve GPU compliance
[**UpdateCompliance**](ComplianceAPI.md#UpdateCompliance) | **Put** /core/compliance | Update a compliance



## CreateCompliance

> CreateUpdateComplianceResponse CreateCompliance(ctx).Payload(payload).Execute()

Create compliance

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
	payload := *openapiclient.NewCompliancePayload(int32(123), "GpuModel_example", "ResourceType_example", int32(123), int32(123), int32(123)) // CompliancePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.CreateCompliance(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CreateCompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCompliance`: CreateUpdateComplianceResponse
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CreateCompliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateComplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CompliancePayload**](CompliancePayload.md) |  | 

### Return type

[**CreateUpdateComplianceResponse**](CreateUpdateComplianceResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCompliance

> ResponseModel DeleteCompliance(ctx, gpuModel).Execute()

Delete a compliance

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
	gpuModel := "gpuModel_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.DeleteCompliance(context.Background(), gpuModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.DeleteCompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCompliance`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.DeleteCompliance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gpuModel** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComplianceRequest struct via the builder pattern


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


## GetCompliance

> ComplianceResponse GetCompliance(ctx).Gpu(gpu).Execute()

Retrieve GPU compliance



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
	gpu := "gpu_example" // string | This is for gpu model (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetCompliance(context.Background()).Gpu(gpu).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetCompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompliance`: ComplianceResponse
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetCompliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gpu** | **string** | This is for gpu model | 

### Return type

[**ComplianceResponse**](ComplianceResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompliance

> CreateUpdateComplianceResponse UpdateCompliance(ctx).Payload(payload).Execute()

Update a compliance

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
	payload := *openapiclient.NewCompliancePayload(int32(123), "GpuModel_example", "ResourceType_example", int32(123), int32(123), int32(123)) // CompliancePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.UpdateCompliance(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.UpdateCompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCompliance`: CreateUpdateComplianceResponse
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.UpdateCompliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateComplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CompliancePayload**](CompliancePayload.md) |  | 

### Return type

[**CreateUpdateComplianceResponse**](CreateUpdateComplianceResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

