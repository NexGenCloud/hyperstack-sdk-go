# \ComplianceAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompliance**](ComplianceAPI.md#CreateCompliance) | **Post** /core/compliance | Create compliance
[**DeleteACompliance**](ComplianceAPI.md#DeleteACompliance) | **Delete** /core/compliance/{gpu_model} | Delete a compliance
[**RetrieveCompliance**](ComplianceAPI.md#RetrieveCompliance) | **Get** /core/compliance | Retrieve GPU compliance
[**UpdateACompliance**](ComplianceAPI.md#UpdateACompliance) | **Put** /core/compliance | Update a compliance



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


## DeleteACompliance

> ResponseModel DeleteACompliance(ctx, gpuModel).Execute()

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
	resp, r, err := apiClient.ComplianceAPI.DeleteACompliance(context.Background(), gpuModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.DeleteACompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteACompliance`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.DeleteACompliance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gpuModel** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAComplianceRequest struct via the builder pattern


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


## RetrieveCompliance

> ComplianceResponse RetrieveCompliance(ctx).Gpu(gpu).Execute()

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
	resp, r, err := apiClient.ComplianceAPI.RetrieveCompliance(context.Background()).Gpu(gpu).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.RetrieveCompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveCompliance`: ComplianceResponse
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.RetrieveCompliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveComplianceRequest struct via the builder pattern


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


## UpdateACompliance

> CreateUpdateComplianceResponse UpdateACompliance(ctx).Payload(payload).Execute()

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
	resp, r, err := apiClient.ComplianceAPI.UpdateACompliance(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.UpdateACompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateACompliance`: CreateUpdateComplianceResponse
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.UpdateACompliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAComplianceRequest struct via the builder pattern


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

