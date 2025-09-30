# \FirewallsAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFirewallRuleToAnExistingFirewall**](FirewallsAPI.md#AddFirewallRuleToAnExistingFirewall) | **Post** /core/firewalls/{firewall_id}/firewall-rules | Add firewall rule to firewall
[**CreateANewFirewall**](FirewallsAPI.md#CreateANewFirewall) | **Post** /core/firewalls | Create firewall
[**DeleteExistingFirewall**](FirewallsAPI.md#DeleteExistingFirewall) | **Delete** /core/firewalls/{id} | Delete firewall
[**DeleteFirewallRulesFromFirewall**](FirewallsAPI.md#DeleteFirewallRulesFromFirewall) | **Delete** /core/firewalls/{firewall_id}/firewall-rules/{firewall_rule_id} | Delete firewall rules from firewall
[**ListExistingFirewalls**](FirewallsAPI.md#ListExistingFirewalls) | **Get** /core/firewalls | List firewalls
[**RetrieveTheDetailsOfAnExistingFirewall**](FirewallsAPI.md#RetrieveTheDetailsOfAnExistingFirewall) | **Get** /core/firewalls/{id} | Retrieve firewall details



## AddFirewallRuleToAnExistingFirewall

> FirewallRule AddFirewallRuleToAnExistingFirewall(ctx, firewallId).Payload(payload).Execute()

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
	resp, r, err := apiClient.FirewallsAPI.AddFirewallRuleToAnExistingFirewall(context.Background(), firewallId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.AddFirewallRuleToAnExistingFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFirewallRuleToAnExistingFirewall`: FirewallRule
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.AddFirewallRuleToAnExistingFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFirewallRuleToAnExistingFirewallRequest struct via the builder pattern


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


## CreateANewFirewall

> FirewallResponse CreateANewFirewall(ctx).Payload(payload).Execute()

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
	resp, r, err := apiClient.FirewallsAPI.CreateANewFirewall(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.CreateANewFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateANewFirewall`: FirewallResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.CreateANewFirewall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateANewFirewallRequest struct via the builder pattern


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


## DeleteExistingFirewall

> ResponseModel DeleteExistingFirewall(ctx, id).Execute()

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
	resp, r, err := apiClient.FirewallsAPI.DeleteExistingFirewall(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.DeleteExistingFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteExistingFirewall`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.DeleteExistingFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExistingFirewallRequest struct via the builder pattern


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


## DeleteFirewallRulesFromFirewall

> ResponseModel DeleteFirewallRulesFromFirewall(ctx, firewallId, firewallRuleId).Execute()

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
	resp, r, err := apiClient.FirewallsAPI.DeleteFirewallRulesFromFirewall(context.Background(), firewallId, firewallRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.DeleteFirewallRulesFromFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFirewallRulesFromFirewall`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.DeleteFirewallRulesFromFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** |  | 
**firewallRuleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRulesFromFirewallRequest struct via the builder pattern


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


## ListExistingFirewalls

> FirewallsListResponse ListExistingFirewalls(ctx).Page(page).PageSize(pageSize).Search(search).Environment(environment).Execute()

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
	resp, r, err := apiClient.FirewallsAPI.ListExistingFirewalls(context.Background()).Page(page).PageSize(pageSize).Search(search).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.ListExistingFirewalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExistingFirewalls`: FirewallsListResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.ListExistingFirewalls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExistingFirewallsRequest struct via the builder pattern


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


## RetrieveTheDetailsOfAnExistingFirewall

> FirewallDetailResponse RetrieveTheDetailsOfAnExistingFirewall(ctx, id).Execute()

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
	resp, r, err := apiClient.FirewallsAPI.RetrieveTheDetailsOfAnExistingFirewall(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsAPI.RetrieveTheDetailsOfAnExistingFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveTheDetailsOfAnExistingFirewall`: FirewallDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallsAPI.RetrieveTheDetailsOfAnExistingFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveTheDetailsOfAnExistingFirewallRequest struct via the builder pattern


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

