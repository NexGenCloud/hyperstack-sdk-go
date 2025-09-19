# \AssigningMemberRoleAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignRBACRoleToUser**](AssigningMemberRoleAPI.md#AssignRBACRoleToUser) | **Put** /auth/users/{user_id}/assign-roles | Assign RBAC Role
[**RemoveRBACRoleFromUser**](AssigningMemberRoleAPI.md#RemoveRBACRoleFromUser) | **Delete** /auth/users/{user_id}/roles | Remove RBAC Role From User



## AssignRBACRoleToUser

> RbacRoleDetailResponseModel AssignRBACRoleToUser(ctx, userId).Payload(payload).Execute()

Assign RBAC Role



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
	userId := int32(56) // int32 | 
	payload := *openapiclient.NewAssignRbacRolePayload(int32(123)) // AssignRbacRolePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssigningMemberRoleAPI.AssignRBACRoleToUser(context.Background(), userId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssigningMemberRoleAPI.AssignRBACRoleToUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignRBACRoleToUser`: RbacRoleDetailResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AssigningMemberRoleAPI.AssignRBACRoleToUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignRBACRoleToUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AssignRbacRolePayload**](AssignRbacRolePayload.md) |  | 

### Return type

[**RbacRoleDetailResponseModel**](RbacRoleDetailResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRBACRoleFromUser

> CommonResponseModel RemoveRBACRoleFromUser(ctx, userId).Execute()

Remove RBAC Role From User



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
	userId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssigningMemberRoleAPI.RemoveRBACRoleFromUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssigningMemberRoleAPI.RemoveRBACRoleFromUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveRBACRoleFromUser`: CommonResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AssigningMemberRoleAPI.RemoveRBACRoleFromUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRBACRoleFromUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommonResponseModel**](CommonResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

