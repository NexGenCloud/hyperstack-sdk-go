# \ZendeskAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetZendeskMessagingToken**](ZendeskAPI.md#GetZendeskMessagingToken) | **Get** /auth/zendesk/messaging-token | Retrieve a Zendesk messaging JWT for the current user



## GetZendeskMessagingToken

> ZendeskMessagingToken GetZendeskMessagingToken(ctx).Execute()

Retrieve a Zendesk messaging JWT for the current user



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
	resp, r, err := apiClient.ZendeskAPI.GetZendeskMessagingToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZendeskAPI.GetZendeskMessagingToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZendeskMessagingToken`: ZendeskMessagingToken
	fmt.Fprintf(os.Stdout, "Response from `ZendeskAPI.GetZendeskMessagingToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetZendeskMessagingTokenRequest struct via the builder pattern


### Return type

[**ZendeskMessagingToken**](ZendeskMessagingToken.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

