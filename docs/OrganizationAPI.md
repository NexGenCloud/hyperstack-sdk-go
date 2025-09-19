# \OrganizationAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemoveOrganizationMember**](OrganizationAPI.md#RemoveOrganizationMember) | **Post** /auth/organizations/remove-member | Remove Organization Member
[**RetrieveOrganizationInformation**](OrganizationAPI.md#RetrieveOrganizationInformation) | **Get** /auth/organizations | Retrieve Organization Information
[**UpdateOrganizationInformation**](OrganizationAPI.md#UpdateOrganizationInformation) | **Put** /auth/organizations/update | Update Organization Information



## RemoveOrganizationMember

> RemoveMemberFromOrganizationResponseModel RemoveOrganizationMember(ctx).Payload(payload).Execute()

Remove Organization Member



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
	payload := *openapiclient.NewRemoveMemberPayload("Email_example") // RemoveMemberPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.RemoveOrganizationMember(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RemoveOrganizationMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveOrganizationMember`: RemoveMemberFromOrganizationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.RemoveOrganizationMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOrganizationMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**RemoveMemberPayload**](RemoveMemberPayload.md) |  | 

### Return type

[**RemoveMemberFromOrganizationResponseModel**](RemoveMemberFromOrganizationResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveOrganizationInformation

> GetOrganizationResponseModel RetrieveOrganizationInformation(ctx).Execute()

Retrieve Organization Information



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
	resp, r, err := apiClient.OrganizationAPI.RetrieveOrganizationInformation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RetrieveOrganizationInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveOrganizationInformation`: GetOrganizationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.RetrieveOrganizationInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveOrganizationInformationRequest struct via the builder pattern


### Return type

[**GetOrganizationResponseModel**](GetOrganizationResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationInformation

> UpdateOrganizationResponseModel UpdateOrganizationInformation(ctx).Payload(payload).Execute()

Update Organization Information



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
	payload := *openapiclient.NewUpdateOrganizationPayload("Name_example") // UpdateOrganizationPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.UpdateOrganizationInformation(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateOrganizationInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationInformation`: UpdateOrganizationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateOrganizationInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**UpdateOrganizationPayload**](UpdateOrganizationPayload.md) |  | 

### Return type

[**UpdateOrganizationResponseModel**](UpdateOrganizationResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

