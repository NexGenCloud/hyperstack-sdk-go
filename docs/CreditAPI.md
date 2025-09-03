# \CreditAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCredit2**](CreditAPI.md#GetCredit2) | **Get** /billing/user-credit/credit | GET: View credit and threshold



## GetCredit2

> GetCreditAndThresholdInfoInResponse GetCredit2(ctx).Execute()

GET: View credit and threshold



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
	resp, r, err := apiClient.CreditAPI.GetCredit2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditAPI.GetCredit2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredit2`: GetCreditAndThresholdInfoInResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditAPI.GetCredit2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredit2Request struct via the builder pattern


### Return type

[**GetCreditAndThresholdInfoInResponse**](GetCreditAndThresholdInfoInResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

