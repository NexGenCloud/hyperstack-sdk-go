# \UserAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**GetUser**](UserAPI.md#GetUser) | **Get** /billing/user/info | GET: Retrieve billing info
[**PostUser**](UserAPI.md#PostUser) | **Post** /billing/user/info | POST: Insert billing info
[**PutUser**](UserAPI.md#PutUser) | **Put** /billing/user/info | PUT: Update billing info



## GetUser

> UsersInfoListResponse GetUser(ctx).Execute()

GET: Retrieve billing info


=======
[**GetFetchUserInfo**](UserAPI.md#GetFetchUserInfo) | **Get** /billing/user/info | GET: Fetch User Info
[**PostInsertUserInfo**](UserAPI.md#PostInsertUserInfo) | **Post** /billing/user/info | POST: Insert user info
[**PutUpdateUserInfo**](UserAPI.md#PutUpdateUserInfo) | **Put** /billing/user/info | PUT: Update user info



## GetFetchUserInfo

> UsersInfoListResponse GetFetchUserInfo(ctx).Execute()

GET: Fetch User Info
>>>>>>> main

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
	resp, r, err := apiClient.UserAPI.GetUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: UsersInfoListResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUser`: %v\n", resp)
=======
	resp, r, err := apiClient.UserAPI.GetFetchUserInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetFetchUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFetchUserInfo`: UsersInfoListResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetFetchUserInfo`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiGetFetchUserInfoRequest struct via the builder pattern
>>>>>>> main


### Return type

[**UsersInfoListResponse**](UsersInfoListResponse.md)

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
## PostUser

> AddUserInfoSuccessResponseModel PostUser(ctx).Payload(payload).Execute()

POST: Insert billing info


=======
## PostInsertUserInfo

> AddUserInfoSuccessResponseModel PostInsertUserInfo(ctx).Payload(payload).Execute()

POST: Insert user info
>>>>>>> main

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
<<<<<<< HEAD
	payload := *openapiclient.NewUserInfoPostPayload(false, "Country_example", "ZipCode_example") // UserInfoPostPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.PostUser(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.PostUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUser`: AddUserInfoSuccessResponseModel
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.PostUser`: %v\n", resp)
=======
	payload := *openapiclient.NewUserinfopostpayload(false, "Country_example", "ZipCode_example") // Userinfopostpayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.PostInsertUserInfo(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.PostInsertUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInsertUserInfo`: AddUserInfoSuccessResponseModel
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.PostInsertUserInfo`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters



### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiPostUserRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiPostInsertUserInfoRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
<<<<<<< HEAD
 **payload** | [**UserInfoPostPayload**](UserInfoPostPayload.md) |  | 
=======
 **payload** | [**Userinfopostpayload**](Userinfopostpayload.md) |  | 
>>>>>>> main

### Return type

[**AddUserInfoSuccessResponseModel**](AddUserInfoSuccessResponseModel.md)

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
## PutUser

> AddUserInfoSuccessResponseModel PutUser(ctx).Payload(payload).Execute()

PUT: Update billing info


=======
## PutUpdateUserInfo

> AddUserInfoSuccessResponseModel PutUpdateUserInfo(ctx).Payload(payload).Execute()

PUT: Update user info
>>>>>>> main

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
<<<<<<< HEAD
	payload := *openapiclient.NewUserInfoPostPayload(false, "Country_example", "ZipCode_example") // UserInfoPostPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.PutUser(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.PutUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutUser`: AddUserInfoSuccessResponseModel
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.PutUser`: %v\n", resp)
=======
	payload := *openapiclient.NewUserinfopostpayload(false, "Country_example", "ZipCode_example") // Userinfopostpayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.PutUpdateUserInfo(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.PutUpdateUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutUpdateUserInfo`: AddUserInfoSuccessResponseModel
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.PutUpdateUserInfo`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters



### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiPutUserRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiPutUpdateUserInfoRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
<<<<<<< HEAD
 **payload** | [**UserInfoPostPayload**](UserInfoPostPayload.md) |  | 
=======
 **payload** | [**Userinfopostpayload**](Userinfopostpayload.md) |  | 
>>>>>>> main

### Return type

[**AddUserInfoSuccessResponseModel**](AddUserInfoSuccessResponseModel.md)

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

