# \AuthAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveAuthenticatedUserDetails**](AuthAPI.md#RetrieveAuthenticatedUserDetails) | **Get** /auth/me | Retrieve Authenticated User Details



## RetrieveAuthenticatedUserDetails

> AuthUserInfoResponseModel RetrieveAuthenticatedUserDetails(ctx).Execute()

Retrieve Authenticated User Details



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
	resp, r, err := apiClient.AuthAPI.RetrieveAuthenticatedUserDetails(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.RetrieveAuthenticatedUserDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveAuthenticatedUserDetails`: AuthUserInfoResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.RetrieveAuthenticatedUserDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAuthenticatedUserDetailsRequest struct via the builder pattern


### Return type

[**AuthUserInfoResponseModel**](AuthUserInfoResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

