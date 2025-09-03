# \FirewallsAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSecurityGroupDetails**](FirewallsAPI.md#DeleteSecurityGroupDetails) | **Delete** /core/firewalls/{id} | Delete firewall
[**DeleteSecurityGroupRuleDelete**](FirewallsAPI.md#DeleteSecurityGroupRuleDelete) | **Delete** /core/firewalls/{firewall_id}/firewall-rules/{firewall_rule_id} | Delete firewall rules from firewall
[**GetSecurityGroup**](FirewallsAPI.md#GetSecurityGroup) | **Get** /core/firewalls | List firewalls
[**GetSecurityGroupDetails**](FirewallsAPI.md#GetSecurityGroupDetails) | **Get** /core/firewalls/{id} | Retrieve firewall details
[**PostSecurityGroup**](FirewallsAPI.md#PostSecurityGroup) | **Post** /core/firewalls | Create firewall
[**PostSecurityGroupRules**](FirewallsAPI.md#PostSecurityGroupRules) | **Post** /core/firewalls/{firewall_id}/firewall-rules | Add firewall rule to firewall



## DeleteSecurityGroupDetails

> ResponseModel DeleteSecurityGroupDetails(ctx, id).Execute()

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
	resp, r, err := apiClient.FirewallsAPI.DeleteSecurityGroupDetails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.DeleteSecurityGroupDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSecurityGroupDetails`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.DeleteSecurityGroupDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityGroupDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityGroupRuleDelete

> ResponseModel DeleteSecurityGroupRuleDelete(ctx, firewallId, firewallRuleId).Execute()

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
	resp, r, err := apiClient.FirewallsAPI.DeleteSecurityGroupRuleDelete(context.Background(), firewallId, firewallRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.DeleteSecurityGroupRuleDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSecurityGroupRuleDelete`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.DeleteSecurityGroupRuleDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** |  | 
**firewallRuleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityGroupRuleDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityGroup

> FirewallsListResponse GetSecurityGroup(ctx).Page(page).PageSize(pageSize).Search(search).Environment(environment).Execute()

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
	resp, r, err := apiClient.FirewallsAPI.GetSecurityGroup(context.Background()).Page(page).PageSize(pageSize).Search(search).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.GetSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityGroup`: FirewallsListResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.GetSecurityGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **search** | **string** |  | 
 **environment** | **string** | Filter Environment ID or Name | 

### Return type

[**FirewallsListResponse**](FirewallsListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityGroupDetails

> FirewallDetailResponse GetSecurityGroupDetails(ctx, id).Execute()

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
	resp, r, err := apiClient.FirewallsAPI.GetSecurityGroupDetails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.GetSecurityGroupDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityGroupDetails`: FirewallDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.GetSecurityGroupDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityGroupDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FirewallDetailResponse**](FirewallDetailResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

