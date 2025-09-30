# \AdminAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendPasswordChangeNotificationEmail**](AdminAPI.md#SendPasswordChangeNotificationEmail) | **Post** /auth/admin/password-change-mail | Send Password Change Notification Email



## SendPasswordChangeNotificationEmail

> CommonResponseModel SendPasswordChangeNotificationEmail(ctx).Execute()

Send Password Change Notification Email



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
	resp, r, err := apiClient.AdminAPI.SendPasswordChangeNotificationEmail(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.SendPasswordChangeNotificationEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendPasswordChangeNotificationEmail`: CommonResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.SendPasswordChangeNotificationEmail`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSendPasswordChangeNotificationEmailRequest struct via the builder pattern


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

