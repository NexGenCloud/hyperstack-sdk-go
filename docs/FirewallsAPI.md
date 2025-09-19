# \FirewallsAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
<<<<<<< HEAD
[**DeleteSecurityGroupDetails**](FirewallsAPI.md#DeleteSecurityGroupDetails) | **Delete** /core/firewalls/{id} | Delete firewall
[**DeleteSecurityGroupRuleDelete**](FirewallsAPI.md#DeleteSecurityGroupRuleDelete) | **Delete** /core/firewalls/{firewall_id}/firewall-rules/{firewall_rule_id} | Delete firewall rules from firewall
[**GetSecurityGroup**](FirewallsAPI.md#GetSecurityGroup) | **Get** /core/firewalls | List firewalls
[**GetSecurityGroupDetails**](FirewallsAPI.md#GetSecurityGroupDetails) | **Get** /core/firewalls/{id} | Retrieve firewall details
[**PostSecurityGroup**](FirewallsAPI.md#PostSecurityGroup) | **Post** /core/firewalls | Create firewall
[**PostSecurityGroupRules**](FirewallsAPI.md#PostSecurityGroupRules) | **Post** /core/firewalls/{firewall_id}/firewall-rules | Add firewall rule to firewall



## DeleteSecurityGroupDetails

> ResponseModel DeleteSecurityGroupDetails(ctx, id).Execute()
=======
[**AddFirewallRuleToFirewall**](FirewallsAPI.md#AddFirewallRuleToFirewall) | **Post** /core/firewalls/{firewall_id}/firewall-rules | Add firewall rule to firewall
[**CreateFirewall**](FirewallsAPI.md#CreateFirewall) | **Post** /core/firewalls | Create firewall
[**DeleteFirewall**](FirewallsAPI.md#DeleteFirewall) | **Delete** /core/firewalls/{id} | Delete firewall
[**DeleteFirewallRulesFromFirewall**](FirewallsAPI.md#DeleteFirewallRulesFromFirewall) | **Delete** /core/firewalls/{firewall_id}/firewall-rules/{firewall_rule_id} | Delete firewall rules from firewall
[**ListFirewalls**](FirewallsAPI.md#ListFirewalls) | **Get** /core/firewalls | List firewalls
[**RetrieveFirewallDetails**](FirewallsAPI.md#RetrieveFirewallDetails) | **Get** /core/firewalls/{id} | Retrieve firewall details



## AddFirewallRuleToFirewall

> FirewallRule AddFirewallRuleToFirewall(ctx, firewallId).Payload(payload).Execute()

Add firewall rule to firewall



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
	firewallId := int32(56) // int32 | 
	payload := *openapiclient.NewCreateFirewallRulePayload("Direction_example", "Ethertype_example", "any", "RemoteIpPrefix_example") // CreateFirewallRulePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsAPI.AddFirewallRuleToFirewall(context.Background(), firewallId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.AddFirewallRuleToFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFirewallRuleToFirewall`: FirewallRule
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.AddFirewallRuleToFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFirewallRuleToFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**CreateFirewallRulePayload**](CreateFirewallRulePayload.md) |  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFirewall

> FirewallResponse CreateFirewall(ctx).Payload(payload).Execute()

Create firewall



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
	payload := *openapiclient.NewCreateFirewallPayload(int32(123), "Name_example") // CreateFirewallPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsAPI.CreateFirewall(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.CreateFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFirewall`: FirewallResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.CreateFirewall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateFirewallPayload**](CreateFirewallPayload.md) |  | 

### Return type

[**FirewallResponse**](FirewallResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewall

> ResponseModel DeleteFirewall(ctx, id).Execute()
>>>>>>> main

Delete firewall



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
	resp, r, err := apiClient.FirewallsAPI.DeleteSecurityGroupDetails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.DeleteSecurityGroupDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSecurityGroupDetails`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.DeleteSecurityGroupDetails`: %v\n", resp)
=======
	resp, r, err := apiClient.FirewallsAPI.DeleteFirewall(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.DeleteFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFirewall`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.DeleteFirewall`: %v\n", resp)
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
Other parameters are passed through a pointer to a apiDeleteSecurityGroupDetailsRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiDeleteFirewallRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseModel**](ResponseModel.md)

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
## DeleteSecurityGroupRuleDelete

> ResponseModel DeleteSecurityGroupRuleDelete(ctx, firewallId, firewallRuleId).Execute()
=======
## DeleteFirewallRulesFromFirewall

> ResponseModel DeleteFirewallRulesFromFirewall(ctx, firewallId, firewallRuleId).Execute()
>>>>>>> main

Delete firewall rules from firewall



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
	firewallId := int32(56) // int32 | 
	firewallRuleId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.FirewallsAPI.DeleteSecurityGroupRuleDelete(context.Background(), firewallId, firewallRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.DeleteSecurityGroupRuleDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSecurityGroupRuleDelete`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.DeleteSecurityGroupRuleDelete`: %v\n", resp)
=======
	resp, r, err := apiClient.FirewallsAPI.DeleteFirewallRulesFromFirewall(context.Background(), firewallId, firewallRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.DeleteFirewallRulesFromFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFirewallRulesFromFirewall`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.DeleteFirewallRulesFromFirewall`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** |  | 
**firewallRuleId** | **int32** |  | 

### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiDeleteSecurityGroupRuleDeleteRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiDeleteFirewallRulesFromFirewallRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseModel**](ResponseModel.md)

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
## GetSecurityGroup

> FirewallsListResponse GetSecurityGroup(ctx).Page(page).PageSize(pageSize).Search(search).Environment(environment).Execute()
=======
## ListFirewalls

> FirewallsListResponse ListFirewalls(ctx).Page(page).PageSize(pageSize).Search(search).Environment(environment).Execute()
>>>>>>> main

List firewalls



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
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	environment := "environment_example" // string | Filter Environment ID or Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
<<<<<<< HEAD
	resp, r, err := apiClient.FirewallsAPI.GetSecurityGroup(context.Background()).Page(page).PageSize(pageSize).Search(search).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.GetSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityGroup`: FirewallsListResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.GetSecurityGroup`: %v\n", resp)
=======
	resp, r, err := apiClient.FirewallsAPI.ListFirewalls(context.Background()).Page(page).PageSize(pageSize).Search(search).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.ListFirewalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFirewalls`: FirewallsListResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.ListFirewalls`: %v\n", resp)
>>>>>>> main
}
```

### Path Parameters



### Other Parameters

<<<<<<< HEAD
Other parameters are passed through a pointer to a apiGetSecurityGroupRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiListFirewallsRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **search** | **string** |  | 
 **environment** | **string** | Filter Environment ID or Name | 

### Return type

[**FirewallsListResponse**](FirewallsListResponse.md)

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
## GetSecurityGroupDetails

> FirewallDetailResponse GetSecurityGroupDetails(ctx, id).Execute()
=======
## RetrieveFirewallDetails

> FirewallDetailResponse RetrieveFirewallDetails(ctx, id).Execute()
>>>>>>> main

Retrieve firewall details



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
	resp, r, err := apiClient.FirewallsAPI.GetSecurityGroupDetails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.GetSecurityGroupDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityGroupDetails`: FirewallDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.GetSecurityGroupDetails`: %v\n", resp)
=======
	resp, r, err := apiClient.FirewallsAPI.RetrieveFirewallDetails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.RetrieveFirewallDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveFirewallDetails`: FirewallDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.RetrieveFirewallDetails`: %v\n", resp)
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
Other parameters are passed through a pointer to a apiGetSecurityGroupDetailsRequest struct via the builder pattern
=======
Other parameters are passed through a pointer to a apiRetrieveFirewallDetailsRequest struct via the builder pattern
>>>>>>> main


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FirewallDetailResponse**](FirewallDetailResponse.md)

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

## PostSecurityGroup

> FirewallResponse PostSecurityGroup(ctx).Payload(payload).Execute()

Create firewall



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
	payload := *openapiclient.NewCreateFirewallPayload(int32(123), "Name_example") // CreateFirewallPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsAPI.PostSecurityGroup(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.PostSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSecurityGroup`: FirewallResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.PostSecurityGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateFirewallPayload**](CreateFirewallPayload.md) |  | 

### Return type

[**FirewallResponse**](FirewallResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSecurityGroupRules

> FirewallRule PostSecurityGroupRules(ctx, firewallId).Payload(payload).Execute()

Add firewall rule to firewall



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
	firewallId := int32(56) // int32 | 
	payload := *openapiclient.NewCreateFirewallRulePayload("Direction_example", "Ethertype_example", "any", "RemoteIpPrefix_example") // CreateFirewallRulePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsAPI.PostSecurityGroupRules(context.Background(), firewallId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.PostSecurityGroupRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSecurityGroupRules`: FirewallRule
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.PostSecurityGroupRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSecurityGroupRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**CreateFirewallRulePayload**](CreateFirewallRulePayload.md) |  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

=======
>>>>>>> main
