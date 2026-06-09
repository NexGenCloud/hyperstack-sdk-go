# \UserConsentEventsAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConsentAuditEvents**](UserConsentEventsAPI.md#GetConsentAuditEvents) | **Get** /auth/user-consent-events/{consent_type}/events | Get audit trail for a consent



## GetConsentAuditEvents

> ConsentEventsResponse GetConsentAuditEvents(ctx, consentType).Execute()

Get audit trail for a consent



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
	consentType := "consentType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserConsentEventsAPI.GetConsentAuditEvents(context.Background(), consentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserConsentEventsAPI.GetConsentAuditEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConsentAuditEvents`: ConsentEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `UserConsentEventsAPI.GetConsentAuditEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentAuditEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsentEventsResponse**](ConsentEventsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

