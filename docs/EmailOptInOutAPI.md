# \EmailOptInOutAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmailPreferencesForAUser**](EmailOptInOutAPI.md#GetEmailPreferencesForAUser) | **Get** /auth/email/opt-out | Get all email preferences for the authenticated user
[**ToggleAllOptionalEmailPreferencesForTheUser**](EmailOptInOutAPI.md#ToggleAllOptionalEmailPreferencesForTheUser) | **Put** /auth/email/opt-out | Toggle all optional email preferences for the authenticated user
[**UpdateEmailPreferenceForACategoryBySlug**](EmailOptInOutAPI.md#UpdateEmailPreferenceForACategoryBySlug) | **Put** /auth/email/opt-out/{slug} | Update email preference opted_in status for a category slug



## GetEmailPreferencesForAUser

> EmailPreferencesResponse GetEmailPreferencesForAUser(ctx).Execute()

Get all email preferences for the authenticated user



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
	resp, r, err := apiClient.EmailOptInOutAPI.GetEmailPreferencesForAUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailOptInOutAPI.GetEmailPreferencesForAUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailPreferencesForAUser`: EmailPreferencesResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailOptInOutAPI.GetEmailPreferencesForAUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailPreferencesForAUserRequest struct via the builder pattern


### Return type

[**EmailPreferencesResponse**](EmailPreferencesResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleAllOptionalEmailPreferencesForTheUser

> UpdateEmailPreferenceResponse ToggleAllOptionalEmailPreferencesForTheUser(ctx).Payload(payload).Execute()

Toggle all optional email preferences for the authenticated user



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
	payload := *openapiclient.NewUpdateEmailPreferenceInput(false) // UpdateEmailPreferenceInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailOptInOutAPI.ToggleAllOptionalEmailPreferencesForTheUser(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailOptInOutAPI.ToggleAllOptionalEmailPreferencesForTheUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToggleAllOptionalEmailPreferencesForTheUser`: UpdateEmailPreferenceResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailOptInOutAPI.ToggleAllOptionalEmailPreferencesForTheUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToggleAllOptionalEmailPreferencesForTheUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**UpdateEmailPreferenceInput**](UpdateEmailPreferenceInput.md) |  | 

### Return type

[**UpdateEmailPreferenceResponse**](UpdateEmailPreferenceResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailPreferenceForACategoryBySlug

> UpdateEmailPreferenceResponse UpdateEmailPreferenceForACategoryBySlug(ctx, slug).Payload(payload).Execute()

Update email preference opted_in status for a category slug



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
	slug := "slug_example" // string | 
	payload := *openapiclient.NewUpdateEmailPreferenceInput(false) // UpdateEmailPreferenceInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailOptInOutAPI.UpdateEmailPreferenceForACategoryBySlug(context.Background(), slug).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailOptInOutAPI.UpdateEmailPreferenceForACategoryBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEmailPreferenceForACategoryBySlug`: UpdateEmailPreferenceResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailOptInOutAPI.UpdateEmailPreferenceForACategoryBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailPreferenceForACategoryBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**UpdateEmailPreferenceInput**](UpdateEmailPreferenceInput.md) |  | 

### Return type

[**UpdateEmailPreferenceResponse**](UpdateEmailPreferenceResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

