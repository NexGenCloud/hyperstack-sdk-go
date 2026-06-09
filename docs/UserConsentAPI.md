# \UserConsentAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddANewConsentForAUser**](UserConsentAPI.md#AddANewConsentForAUser) | **Post** /auth/user-consent | Add a new User consent
[**GetAllConsentTemplates**](UserConsentAPI.md#GetAllConsentTemplates) | **Get** /auth/user-consent/templates | Get all consent templates
[**GetAllConsentsForAUser**](UserConsentAPI.md#GetAllConsentsForAUser) | **Get** /auth/user-consent | Get Consents for a User
[**GetConsentTemplateByType**](UserConsentAPI.md#GetConsentTemplateByType) | **Get** /auth/user-consent/templates/{consent_type} | Get consent template for a specific type
[**UpdateAConsentActionByType**](UserConsentAPI.md#UpdateAConsentActionByType) | **Patch** /auth/user-consent/{consent_type} | Grant or revoke an existing consent



## AddANewConsentForAUser

> ConsentActionResponse AddANewConsentForAUser(ctx).Payload(payload).Execute()

Add a new User consent



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
	payload := *openapiclient.NewRecordConsentRequest("granted", "auto_top_up") // RecordConsentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserConsentAPI.AddANewConsentForAUser(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserConsentAPI.AddANewConsentForAUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddANewConsentForAUser`: ConsentActionResponse
	fmt.Fprintf(os.Stdout, "Response from `UserConsentAPI.AddANewConsentForAUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddANewConsentForAUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**RecordConsentRequest**](RecordConsentRequest.md) |  | 

### Return type

[**ConsentActionResponse**](ConsentActionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllConsentTemplates

> ConsentTemplatesResponse GetAllConsentTemplates(ctx).Execute()

Get all consent templates



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
	resp, r, err := apiClient.UserConsentAPI.GetAllConsentTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserConsentAPI.GetAllConsentTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllConsentTemplates`: ConsentTemplatesResponse
	fmt.Fprintf(os.Stdout, "Response from `UserConsentAPI.GetAllConsentTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConsentTemplatesRequest struct via the builder pattern


### Return type

[**ConsentTemplatesResponse**](ConsentTemplatesResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllConsentsForAUser

> UserConsentsResponse GetAllConsentsForAUser(ctx).ConsentType(consentType).Execute()

Get Consents for a User



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
	consentType := "consentType_example" // string | Filter by consent type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserConsentAPI.GetAllConsentsForAUser(context.Background()).ConsentType(consentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserConsentAPI.GetAllConsentsForAUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllConsentsForAUser`: UserConsentsResponse
	fmt.Fprintf(os.Stdout, "Response from `UserConsentAPI.GetAllConsentsForAUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConsentsForAUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consentType** | **string** | Filter by consent type | 

### Return type

[**UserConsentsResponse**](UserConsentsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsentTemplateByType

> ConsentTemplate GetConsentTemplateByType(ctx, consentType).Execute()

Get consent template for a specific type



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
	resp, r, err := apiClient.UserConsentAPI.GetConsentTemplateByType(context.Background(), consentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserConsentAPI.GetConsentTemplateByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConsentTemplateByType`: ConsentTemplate
	fmt.Fprintf(os.Stdout, "Response from `UserConsentAPI.GetConsentTemplateByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentTemplateByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsentTemplate**](ConsentTemplate.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAConsentActionByType

> ConsentActionResponse UpdateAConsentActionByType(ctx, consentType).Payload(payload).Execute()

Grant or revoke an existing consent



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
	payload := *openapiclient.NewUpdateConsentRequest("granted") // UpdateConsentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserConsentAPI.UpdateAConsentActionByType(context.Background(), consentType).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserConsentAPI.UpdateAConsentActionByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAConsentActionByType`: ConsentActionResponse
	fmt.Fprintf(os.Stdout, "Response from `UserConsentAPI.UpdateAConsentActionByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAConsentActionByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**UpdateConsentRequest**](UpdateConsentRequest.md) |  | 

### Return type

[**ConsentActionResponse**](ConsentActionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

