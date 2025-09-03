# \PartnerConfigAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPartnerConfig**](PartnerConfigAPI.md#GetPartnerConfig) | **Get** /auth/partner-config | Get partner config
[**GetPartnerConfigByDomain**](PartnerConfigAPI.md#GetPartnerConfigByDomain) | **Get** /auth/partner-config/docs | 



## GetPartnerConfig

> PartnerConfig GetPartnerConfig(ctx).Execute()

Get partner config



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
	resp, r, err := apiClient.PartnerConfigAPI.GetPartnerConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnerConfigAPI.GetPartnerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPartnerConfig`: PartnerConfig
	fmt.Fprintf(os.Stdout, "Response from `PartnerConfigAPI.GetPartnerConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnerConfigRequest struct via the builder pattern


### Return type

[**PartnerConfig**](PartnerConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartnerConfigByDomain

> PartnerConfig GetPartnerConfigByDomain(ctx).Domain(domain).Execute()





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
	domain := "domain_example" // string | The domain to look up the partner config for. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PartnerConfigAPI.GetPartnerConfigByDomain(context.Background()).Domain(domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnerConfigAPI.GetPartnerConfigByDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPartnerConfigByDomain`: PartnerConfig
	fmt.Fprintf(os.Stdout, "Response from `PartnerConfigAPI.GetPartnerConfigByDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnerConfigByDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | The domain to look up the partner config for. | 

### Return type

[**PartnerConfig**](PartnerConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

