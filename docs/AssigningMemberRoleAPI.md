# \AssigningMemberRoleAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**AssignRBACRoleToUser**](AssigningMemberRoleAPI.md#AssignRBACRoleToUser) | **Put** /auth/users/{user_id}/assign-roles | Assign RBAC Role
[**RemoveRBACRoleFromUser**](AssigningMemberRoleAPI.md#RemoveRBACRoleFromUser) | **Delete** /auth/users/{user_id}/roles | Remove RBAC Role From User



## AssignRBACRoleToUser

> RbacRoleDetailResponseModel AssignRBACRoleToUser(ctx, userId).Payload(payload).Execute()
=======
[**AssignRbacRole**](AssigningMemberRoleAPI.md#AssignRbacRole) | **Put** /auth/users/{user_id}/assign-roles | Assign RBAC Role
[**RemoveRbacRoleFromUser**](AssigningMemberRoleAPI.md#RemoveRbacRoleFromUser) | **Delete** /auth/users/{user_id}/roles | Remove RBAC Role From User



## AssignRbacRole

> RbacRoleDetailResponseModel AssignRbacRole(ctx, userId).Payload(payload).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.AssigningMemberRoleAPI.AssignRBACRoleToUser(context.Background(), userId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssigningMemberRoleAPI.AssignRBACRoleToUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignRBACRoleToUser`: RbacRoleDetailResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AssigningMemberRoleAPI.AssignRBACRoleToUser`: %v\n", resp)
=======
	resp, r, err := apiClient.AssigningMemberRoleAPI.AssignRbacRole(context.Background(), userId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssigningMemberRoleAPI.AssignRbacRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignRbacRole`: RbacRoleDetailResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AssigningMemberRoleAPI.AssignRbacRole`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiAssignRBACRoleToUserRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiAssignRbacRoleRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AssignRbacRolePayload**](AssignRbacRolePayload.md) |  | 

### Return type

[**RbacRoleDetailResponseModel**](RbacRoleDetailResponseModel.md)

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


<<<<<<< HEAD
## RemoveRBACRoleFromUser

> CommonResponseModel RemoveRBACRoleFromUser(ctx, userId).Execute()
=======
## RemoveRbacRoleFromUser

> CommonResponseModel RemoveRbacRoleFromUser(ctx, userId).Execute()
>>>>>>> main

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
<<<<<<< HEAD
	resp, r, err := apiClient.AssigningMemberRoleAPI.RemoveRBACRoleFromUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssigningMemberRoleAPI.RemoveRBACRoleFromUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveRBACRoleFromUser`: CommonResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AssigningMemberRoleAPI.RemoveRBACRoleFromUser`: %v\n", resp)
=======
	resp, r, err := apiClient.AssigningMemberRoleAPI.RemoveRbacRoleFromUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssigningMemberRoleAPI.RemoveRbacRoleFromUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveRbacRoleFromUser`: CommonResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AssigningMemberRoleAPI.RemoveRbacRoleFromUser`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiRemoveRBACRoleFromUserRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRemoveRbacRoleFromUserRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommonResponseModel**](CommonResponseModel.md)

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

