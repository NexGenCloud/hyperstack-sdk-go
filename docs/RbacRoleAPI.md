# \RbacRoleAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**CreateRBACRole**](RbacRoleAPI.md#CreateRBACRole) | **Post** /auth/roles | Create RBAC Role
[**DeleteRBACRole**](RbacRoleAPI.md#DeleteRBACRole) | **Delete** /auth/roles/{id} | Delete RBAC Role
[**ListRBACRoles**](RbacRoleAPI.md#ListRBACRoles) | **Get** /auth/roles | List RBAC Roles
[**RetrieveRBACRoleDetails**](RbacRoleAPI.md#RetrieveRBACRoleDetails) | **Get** /auth/roles/{id} | Retrieve RBAC Role Details
[**UpdateRBACRole**](RbacRoleAPI.md#UpdateRBACRole) | **Put** /auth/roles/{id} | Update RBAC Role



## CreateRBACRole

> RbacRoleDetailResponseModel CreateRBACRole(ctx).Payload(payload).Execute()
=======
[**CreateRbacRole**](RbacRoleAPI.md#CreateRbacRole) | **Post** /auth/roles | Create RBAC Role
[**DeleteRbacRole**](RbacRoleAPI.md#DeleteRbacRole) | **Delete** /auth/roles/{id} | Delete RBAC Role
[**ListRbacRoles**](RbacRoleAPI.md#ListRbacRoles) | **Get** /auth/roles | List RBAC Roles
[**RetrieveRbacRoleDetails**](RbacRoleAPI.md#RetrieveRbacRoleDetails) | **Get** /auth/roles/{id} | Retrieve RBAC Role Details
[**UpdateRbacRole**](RbacRoleAPI.md#UpdateRbacRole) | **Put** /auth/roles/{id} | Update RBAC Role



## CreateRbacRole

> RbacRoleDetailResponseModel CreateRbacRole(ctx).Payload(payload).Execute()
>>>>>>> main

Create RBAC Role



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
	payload := *openapiclient.NewCreateUpdateRbacRolePayload("Description_example", "Name_example") // CreateUpdateRbacRolePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.RbacRoleAPI.CreateRBACRole(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacRoleAPI.CreateRBACRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRBACRole`: RbacRoleDetailResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RbacRoleAPI.CreateRBACRole`: %v\n", resp)
=======
	resp, r, err := apiClient.RbacRoleAPI.CreateRbacRole(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacRoleAPI.CreateRbacRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRbacRole`: RbacRoleDetailResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RbacRoleAPI.CreateRbacRole`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters



### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiCreateRBACRoleRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiCreateRbacRoleRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateUpdateRbacRolePayload**](CreateUpdateRbacRolePayload.md) |  | 

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
## DeleteRBACRole

> CommonResponseModel DeleteRBACRole(ctx, id).Execute()
=======
## DeleteRbacRole

> CommonResponseModel DeleteRbacRole(ctx, id).Execute()
>>>>>>> main

Delete RBAC Role



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
	resp, r, err := apiClient.RbacRoleAPI.DeleteRBACRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacRoleAPI.DeleteRBACRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRBACRole`: CommonResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RbacRoleAPI.DeleteRBACRole`: %v\n", resp)
=======
	resp, r, err := apiClient.RbacRoleAPI.DeleteRbacRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacRoleAPI.DeleteRbacRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRbacRole`: CommonResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RbacRoleAPI.DeleteRbacRole`: %v\n", resp)
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
Other parameters are passed through a pointer to a apiDeleteRBACRoleRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiDeleteRbacRoleRequest struct via the builder pattern
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


<<<<<<< HEAD
## ListRBACRoles

> GetRbacRolesResponseModel ListRBACRoles(ctx).Execute()
=======
## ListRbacRoles

> GetRbacRolesResponseModel ListRbacRoles(ctx).Execute()
>>>>>>> main

List RBAC Roles



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
	resp, r, err := apiClient.RbacRoleAPI.ListRBACRoles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacRoleAPI.ListRBACRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRBACRoles`: GetRbacRolesResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RbacRoleAPI.ListRBACRoles`: %v\n", resp)
=======
	resp, r, err := apiClient.RbacRoleAPI.ListRbacRoles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacRoleAPI.ListRbacRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRbacRoles`: GetRbacRolesResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RbacRoleAPI.ListRbacRoles`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiListRBACRolesRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiListRbacRolesRequest struct via the builder pattern
>>>>>>> main


### Return type

[**GetRbacRolesResponseModel**](GetRbacRolesResponseModel.md)

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
## RetrieveRBACRoleDetails

> RbacRoleDetailResponseModelFixed RetrieveRBACRoleDetails(ctx, id).Execute()
=======
## RetrieveRbacRoleDetails

> RbacRoleDetailResponseModelFixed RetrieveRbacRoleDetails(ctx, id).Execute()
>>>>>>> main

Retrieve RBAC Role Details



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
	resp, r, err := apiClient.RbacRoleAPI.RetrieveRBACRoleDetails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacRoleAPI.RetrieveRBACRoleDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveRBACRoleDetails`: RbacRoleDetailResponseModelFixed
	fmt.Fprintf(os.Stdout, "Response from `RbacRoleAPI.RetrieveRBACRoleDetails`: %v\n", resp)
=======
	resp, r, err := apiClient.RbacRoleAPI.RetrieveRbacRoleDetails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacRoleAPI.RetrieveRbacRoleDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveRbacRoleDetails`: RbacRoleDetailResponseModelFixed
	fmt.Fprintf(os.Stdout, "Response from `RbacRoleAPI.RetrieveRbacRoleDetails`: %v\n", resp)
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
Other parameters are passed through a pointer to a apiRetrieveRBACRoleDetailsRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveRbacRoleDetailsRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RbacRoleDetailResponseModelFixed**](RbacRoleDetailResponseModelFixed.md)

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
## UpdateRBACRole

> RbacRoleDetailResponseModel UpdateRBACRole(ctx, id).Payload(payload).Execute()
=======
## UpdateRbacRole

> RbacRoleDetailResponseModel UpdateRbacRole(ctx, id).Payload(payload).Execute()
>>>>>>> main

Update RBAC Role



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
	payload := *openapiclient.NewCreateUpdateRbacRolePayload("Description_example", "Name_example") // CreateUpdateRbacRolePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.RbacRoleAPI.UpdateRBACRole(context.Background(), id).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacRoleAPI.UpdateRBACRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRBACRole`: RbacRoleDetailResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RbacRoleAPI.UpdateRBACRole`: %v\n", resp)
=======
	resp, r, err := apiClient.RbacRoleAPI.UpdateRbacRole(context.Background(), id).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacRoleAPI.UpdateRbacRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRbacRole`: RbacRoleDetailResponseModel
	fmt.Fprintf(os.Stdout, "Response from `RbacRoleAPI.UpdateRbacRole`: %v\n", resp)
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
Other parameters are passed through a pointer to a apiUpdateRBACRoleRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiUpdateRbacRoleRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**CreateUpdateRbacRolePayload**](CreateUpdateRbacRolePayload.md) |  | 

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

