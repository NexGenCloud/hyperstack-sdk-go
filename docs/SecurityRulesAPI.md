# \SecurityRulesAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFirewallRuleProtocols**](SecurityRulesAPI.md#ListFirewallRuleProtocols) | **Get** /core/sg-rules-protocols | List firewall rule protocols



## ListFirewallRuleProtocols

> SecurityRulesProtocolFields ListFirewallRuleProtocols(ctx).Execute()

List firewall rule protocols



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
	resp, r, err := apiClient.SecurityRulesAPI.ListFirewallRuleProtocols(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityRulesAPI.ListFirewallRuleProtocols``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFirewallRuleProtocols`: SecurityRulesProtocolFields
	fmt.Fprintf(os.Stdout, "Response from `SecurityRulesAPI.ListFirewallRuleProtocols`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFirewallRuleProtocolsRequest struct via the builder pattern


### Return type

[**SecurityRulesProtocolFields**](SecurityRulesProtocolFields.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

