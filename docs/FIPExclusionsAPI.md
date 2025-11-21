# \FIPExclusionsAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckIfOrgIsExcludedFromFloatingIPDetachment**](FIPExclusionsAPI.md#CheckIfOrgIsExcludedFromFloatingIPDetachment) | **Get** /core/fip-detachment-exclusions/org/{org_id} | 



## CheckIfOrgIsExcludedFromFloatingIPDetachment

> ResponseModel CheckIfOrgIsExcludedFromFloatingIPDetachment(ctx, orgId).Execute()





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
	orgId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FIPExclusionsAPI.CheckIfOrgIsExcludedFromFloatingIPDetachment(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FIPExclusionsAPI.CheckIfOrgIsExcludedFromFloatingIPDetachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckIfOrgIsExcludedFromFloatingIPDetachment`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FIPExclusionsAPI.CheckIfOrgIsExcludedFromFloatingIPDetachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckIfOrgIsExcludedFromFloatingIPDetachmentRequest struct via the builder pattern


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

